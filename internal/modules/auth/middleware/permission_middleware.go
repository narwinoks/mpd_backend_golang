package middleware

import (
	"backend-app/internal/core/cache"
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/repository/permission"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type PermissionMiddleware struct {
	repo permission.PermissionRepository
}

func NewPermissionMiddleware(repo permission.PermissionRepository) *PermissionMiddleware {
	return &PermissionMiddleware{repo: repo}
}

func (m *PermissionMiddleware) Handle(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDVal, exists := c.Get("user_id")
		if !exists {
			c.Error(exception.NewUnauthorizedError("User not found in context"))
			c.Abort()
			return
		}

		roleIDVal, exists := c.Get("role_id")
		if !exists {
			c.Error(exception.NewUnauthorizedError("Role not found in context"))
			c.Abort()
			return
		}

		userID := userIDVal.(uint32)
		roleID := roleIDVal.(uint32)

		cacheKey := fmt.Sprintf("user_permissions_%d", userID)
		cacher := cache.GetCache()

		var permissions []string
		if found := cacher.Get(cacheKey, &permissions); !found {
			var err error
			permissions, err = m.repo.GetUserPermissions(userID, roleID)
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}
			cacher.Set(cacheKey, permissions, 5*time.Minute)
		}

		hasPermission := false
		for _, p := range permissions {
			if p == requiredPermission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.Error(exception.NewForbiddenError(fmt.Sprintf("Permission '%s' is required", requiredPermission)))
			c.Abort()
			return
		}

		c.Next()
	}
}
