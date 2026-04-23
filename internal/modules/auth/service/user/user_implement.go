package user

import (
	"backend-app/config"
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/repository/user"
	req "backend-app/internal/modules/auth/request/user"
	res "backend-app/internal/modules/auth/response/user"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	repo   user.UserRepository
	config *config.Config
}

func NewUserService(repo user.UserRepository, config *config.Config) UserService {
	return &userServiceImpl{
		repo:   repo,
		config: config,
	}
}

func (s *userServiceImpl) Login(request *req.LoginRequest) (*res.LoginResponse, error) {
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
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      accessTokenExpiration.Unix(),
		"type":     "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		logrus.Errorf("Failed to sign access token: %v", err)
		return nil, fmt.Errorf("failed to generate access token")
	}

	refreshTokenExpiration := time.Now().Add(time.Duration(s.config.JWT.RefreshTokenExpiration) * 24 * time.Hour)
	refreshTokenClaims := &jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      refreshTokenExpiration.Unix(),
		"type":     "refresh",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		logrus.Errorf("Failed to sign refresh token: %v", err)
		return nil, fmt.Errorf("failed to generate refresh token")
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
