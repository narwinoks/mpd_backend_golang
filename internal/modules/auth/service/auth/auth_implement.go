package auth

import (
	"backend-app/config"
	baseModels "backend-app/internal/base/models"
	"backend-app/internal/core/exception"
	authModels "backend-app/internal/modules/auth/models"
	tokenRepo "backend-app/internal/modules/auth/repository/personal_access_token"
	"backend-app/internal/modules/auth/repository/user"
	req "backend-app/internal/modules/auth/request/user"
	res "backend-app/internal/modules/auth/response/user"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	repo      user.UserRepository
	tokenRepo tokenRepo.TokenRepository
	config    *config.Config
}

func NewAuthService(repo user.UserRepository, tokenRepo tokenRepo.TokenRepository, config *config.Config) AuthService {
	return &authServiceImpl{
		repo:      repo,
		tokenRepo: tokenRepo,
		config:    config,
	}
}

func (s *authServiceImpl) Login(request *req.LoginRequest) (*res.LoginResponse, error) {
	logrus.WithFields(logrus.Fields{
		"module":   "auth",
		"action":   "login",
		"identity": request.Identity,
	}).Info("Attempting login")

	user, err := s.repo.FindByIdentity(request.Identity)
	if err != nil {
		logrus.Errorf("Login failed: identity not found: %v", err)
		return nil, exception.NewUnauthorizedError("Invalid credentials")
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		logrus.Errorf("Login failed: password mismatch for %s", request.Identity)
		return nil, exception.NewUnauthorizedError("Invalid credentials")
	}

	// Generate Tokens
	accessTokenExpiration := time.Now().Add(time.Duration(s.config.JWT.AccessTokenExpiration) * time.Minute)
	accessTokenClaims := &jwt.MapClaims{
		"user_id":     user.ID,
		"username":    user.Username,
		"employee_id": user.EmployeeID,
		"profile_id":  user.ProfileID,
		"exp":         accessTokenExpiration.Unix(),
		"type":        "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		logrus.Errorf("Failed to sign access token: %v", err)
		return nil, fmt.Errorf("failed to generate access token")
	}

	refreshTokenExpiration := time.Now().Add(time.Duration(s.config.JWT.RefreshTokenExpiration) * 24 * time.Hour)
	refreshTokenClaims := &jwt.MapClaims{
		"user_id":     user.ID,
		"username":    user.Username,
		"employee_id": user.EmployeeID,
		"profile_id":  user.ProfileID,
		"exp":         refreshTokenExpiration.Unix(),
		"type":        "refresh",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		logrus.Errorf("Failed to sign refresh token: %v", err)
		return nil, fmt.Errorf("failed to generate refresh token")
	}

	// Store Access Token
	err = s.tokenRepo.Create(&authModels.PersonalAccessToken{
		BaseModel: baseModels.BaseModel{
			UUID:      uuid.New().String(),
			IsActive:  true,
			ProfileID: user.ProfileID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:    user.ID,
		Token:     accessTokenString,
		ExpiredAt: accessTokenExpiration,
	})
	if err != nil {
		logrus.Errorf("Failed to store access token: %v", err)
	}

	// Store Refresh Token
	err = s.tokenRepo.Create(&authModels.PersonalAccessToken{
		BaseModel: baseModels.BaseModel{
			UUID:      uuid.New().String(),
			IsActive:  true,
			ProfileID: user.ProfileID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:    user.ID,
		Token:     refreshTokenString,
		ExpiredAt: refreshTokenExpiration,
	})
	if err != nil {
		logrus.Errorf("Failed to store refresh token: %v", err)
	}

	logrus.Infof("Login successful for user: %s", user.Username)

	return &res.LoginResponse{
		AccessToken:      accessTokenString,
		RefreshToken:     refreshTokenString,
		TokenType:        "Bearer",
		ExpiresIn:        s.config.JWT.AccessTokenExpiration * 60,
		RefreshExpiresIn: s.config.JWT.RefreshTokenExpiration * 24 * 3600,
	}, nil
}

func (s *authServiceImpl) RefreshToken(request *req.RefreshTokenRequest) (*res.LoginResponse, error) {
	logrus.Info("Attempting to refresh token")

	// 1. Verify token signature and validity
	token, err := jwt.Parse(request.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		logrus.Errorf("Invalid refresh token: %v", err)
		return nil, exception.NewUnauthorizedError("Invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "refresh" {
		return nil, exception.NewUnauthorizedError("Invalid token type")
	}

	// 2. Check if token is blacklisted/revoked
	isRevoked, err := s.tokenRepo.IsRevoked(request.RefreshToken)
	if err != nil {
		logrus.Errorf("Failed to check token status: %v", err)
		return nil, fmt.Errorf("internal server error")
	}
	if isRevoked {
		return nil, exception.NewUnauthorizedError("Token has been revoked")
	}

	userID := uint32(claims["user_id"].(float64))
	username := claims["username"].(string)

	// Fetch user to get ProfileID and ensure user still exists
	user, err := s.repo.FindByID(userID)
	if err != nil {
		logrus.Errorf("User not found during refresh: %v", err)
		return nil, exception.NewUnauthorizedError("User not found")
	}

	// 3. Generate New Tokens
	accessTokenExpiration := time.Now().Add(time.Duration(s.config.JWT.AccessTokenExpiration) * time.Minute)
	accessTokenClaims := &jwt.MapClaims{
		"user_id":     userID,
		"username":    username,
		"employee_id": user.EmployeeID,
		"profile_id":  user.ProfileID,
		"exp":         accessTokenExpiration.Unix(),
		"type":        "access",
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	newAccessTokenString, err := newAccessToken.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token")
	}

	newRefreshTokenExpiration := time.Now().Add(time.Duration(s.config.JWT.RefreshTokenExpiration) * 24 * time.Hour)
	newRefreshTokenClaims := &jwt.MapClaims{
		"user_id":     userID,
		"username":    username,
		"employee_id": user.EmployeeID,
		"profile_id":  user.ProfileID,
		"exp":         newRefreshTokenExpiration.Unix(),
		"type":        "refresh",
	}

	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newRefreshTokenClaims)
	newRefreshTokenString, err := newRefreshToken.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token")
	}

	// 4. Blacklist old token and store new one
	err = s.tokenRepo.RevokeByToken(request.RefreshToken)
	if err != nil {
		logrus.Errorf("Failed to revoke old token: %v", err)
	}

	// Store new access token
	err = s.tokenRepo.Create(&authModels.PersonalAccessToken{
		BaseModel: baseModels.BaseModel{
			UUID:      uuid.New().String(),
			IsActive:  true,
			ProfileID: user.ProfileID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:    userID,
		Token:     newAccessTokenString,
		ExpiredAt: accessTokenExpiration,
	})
	if err != nil {
		logrus.Errorf("Failed to store new access token: %v", err)
	}

	err = s.tokenRepo.Create(&authModels.PersonalAccessToken{
		BaseModel: baseModels.BaseModel{
			UUID:      uuid.New().String(),
			IsActive:  true,
			ProfileID: user.ProfileID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserID:    userID,
		Token:     newRefreshTokenString,
		ExpiredAt: newRefreshTokenExpiration,
	})
	if err != nil {
		logrus.Errorf("Failed to store new refresh token: %v", err)
	}

	logrus.Infof("Token refresh successful for user_id: %d", userID)

	return &res.LoginResponse{
		AccessToken:      newAccessTokenString,
		RefreshToken:     newRefreshTokenString,
		TokenType:        "Bearer",
		ExpiresIn:        s.config.JWT.AccessTokenExpiration * 60,
		RefreshExpiresIn: s.config.JWT.RefreshTokenExpiration * 24 * 3600,
	}, nil
}

func (s *authServiceImpl) Logout(userID uint32) error {
	logrus.Infof("Attempting logout and revocation of all tokens for user_id: %d", userID)

	// Revoke all tokens for the user
	err := s.tokenRepo.RevokeByUserID(userID)
	if err != nil {
		logrus.Errorf("Failed to revoke tokens during logout for user %d: %v", userID, err)
		return fmt.Errorf("failed to logout")
	}

	return nil
}
