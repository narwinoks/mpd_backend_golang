package middleware

import (
	"backend-app/internal/core/cache"
	"backend-app/internal/core/exception"
	"backend-app/internal/modules/auth/repository/module"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ModuleMiddleware struct {
	repo module.ModuleRepository
}

func NewModuleMiddleware(repo module.ModuleRepository) *ModuleMiddleware {
	return &ModuleMiddleware{repo: repo}
}

func (m *ModuleMiddleware) Handle(requiredModuleCode string) gin.HandlerFunc {
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

		cacheKey := fmt.Sprintf("user_modules_%d", userID)
		cacher := cache.GetCache()

		var moduleCodes []string
		if found := cacher.Get(cacheKey, &moduleCodes); !found {
			modules, err := m.repo.GetUserModules(userID, roleID)
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}

			moduleCodes = make([]string, len(modules))
			for i, mod := range modules {
				moduleCodes[i] = mod.Code
			}

			cacher.Set(cacheKey, moduleCodes, 5*time.Minute)
		}

		hasModule := false
		for _, code := range moduleCodes {
			if code == requiredModuleCode {
				hasModule = true
				break
			}
		}

		if !hasModule {
			c.Error(exception.NewForbiddenError(fmt.Sprintf("Module access '%s' is required", requiredModuleCode)))
			c.Abort()
			return
		}

		c.Next()
	}
}
