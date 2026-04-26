package middleware

import (
	"backend-app/config"
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/repository/personal_access_token"
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(cfg *config.Config, tokenRepo personal_access_token.TokenRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(exception.NewUnauthorizedError("Authorization header is required"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Error(exception.NewUnauthorizedError("Invalid authorization header format"))
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil || !token.Valid {
			c.Error(exception.NewUnauthorizedError("Invalid or expired token"))
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["type"] != "access" {
			c.Error(exception.NewUnauthorizedError("Invalid token type"))
			c.Abort()
			return
		}

		// Check if token is revoked
		isRevoked, err := tokenRepo.IsRevoked(tokenString)
		if err != nil {
			c.Error(fmt.Errorf("failed to check token status: %v", err))
			c.Abort()
			return
		}
		if isRevoked {
			c.Error(exception.NewUnauthorizedError("Token has been revoked"))
			c.Abort()
			return
		}

		// Check if token exists in DB (to ensure it was issued by us if we want to be strict)
		// but IsRevoked already checks this and returns false if not found.
		// However, my current IsRevoked implementation returns false if NOT FOUND.
		// If we want to be strict, we should ensure it exists.
		// Let's refine IsRevoked or add a CheckExists.

		userID := uint32(claims["user_id"].(float64))
		username := claims["username"].(string)
		roleID := uint32(claims["role_id"].(float64))

		var employeeID *uint32
		if claims["employee_id"] != nil {
			val := uint32(claims["employee_id"].(float64))
			employeeID = &val
		}

		var profileID *uint32
		if claims["profile_id"] != nil {
			val := uint32(claims["profile_id"].(float64))
			profileID = &val
		}

		// Set to Gin context
		c.Set("user_id", userID)
		c.Set("username", username)
		c.Set("role_id", roleID)
		if employeeID != nil {
			c.Set("employee_id", *employeeID)
		}
		if profileID != nil {
			c.Set("profile_id", *profileID)
		}

		// Also set to standard context for GORM hooks
		ctx := c.Request.Context()
		if employeeID != nil {
			ctx = context.WithValue(ctx, "employee_id", *employeeID)
		}
		if profileID != nil {
			ctx = context.WithValue(ctx, "profile_id", *profileID)
		}
		ctx = context.WithValue(ctx, "external_code", cfg.App.ExternalCode)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
