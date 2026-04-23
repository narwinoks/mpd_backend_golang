package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ResponseIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate UUID v7 (time-ordered)
		requestID, err := uuid.NewV7()
		if err != nil {
			// Fallback to v4 if v7 generation fails
			requestID = uuid.New()
		}

		// Set in context
		c.Set("request_id", requestID.String())

		// Set in response header
		c.Writer.Header().Set("X-Request-ID", requestID.String())

		c.Next()
	}
}
