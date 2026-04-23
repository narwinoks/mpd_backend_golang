package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		c.Next()

		// Get any errors from the context
		if len(c.Errors) > 0 {
			// Extract custom fields from context if available
			username, _ := c.Get("username")
			if username == nil {
				username = "anonymous"
			}
			module, _ := c.Get("module")
			if module == nil {
				module = "unknown"
			}
			action, _ := c.Get("action")
			if action == nil {
				action = c.Request.Method + " " + c.Request.URL.Path
			}

			entry := logrus.WithFields(logrus.Fields{
				"username":   username,
				"ip_address": c.ClientIP(),
				"module":     module,
				"action":     action,
			})

			for _, e := range c.Errors {
				entry.Error(e.Err.Error())
			}
		}
	}
}
