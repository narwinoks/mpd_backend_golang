package response

import (
	"backend-app/internal/core/exception"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Only execute if there are errors in the context
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// 1. Handle Custom NotFoundError
			var notFoundErr exception.NotFoundError
			if errors.As(err, &notFoundErr) {
				SendError(c, DataNotFound, notFoundErr.Error())
				c.Abort()
				return
			}

			// 2. Handle Custom BadRequestError
			var badRequestErr exception.BadRequestError
			if errors.As(err, &badRequestErr) {
				SendError(c, BadRequest, badRequestErr.Error())
				c.Abort()
				return
			}

			// 2.1 Handle Custom ConflictError
			var conflictErr exception.ConflictError
			if errors.As(err, &conflictErr) {
				SendError(c, Conflict, conflictErr.Error())
				c.Abort()
				return
			}

			// 2.2 Handle Custom UnauthorizedError
			var unauthorizedErr exception.UnauthorizedError
			if errors.As(err, &unauthorizedErr) {
				SendError(c, Unauthorized, unauthorizedErr.Error())
				c.Abort()
				return
			}

			// 2.3 Handle Custom ForbiddenError
			var forbiddenErr exception.ForbiddenError
			if errors.As(err, &forbiddenErr) {
				SendError(c, Forbidden, forbiddenErr.Error())
				c.Abort()
				return
			}

			// 3. Handle Binding/Validation Errors (400)

			// Case: Empty Body (EOF)
			if errors.Is(err, io.EOF) {
				SendError(c, BadRequest, "Request body cannot be empty")
				c.Abort()
				return
			}

			// Case: JSON Syntax/Type Mismatch
			if isJSONError(err) {
				SendError(c, BadRequest, err.Error())
				c.Abort()
				return
			}

			// Case: Struct Validation (validator/v10)
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				out := make(map[string]string)
				for _, fe := range ve {
					out[fe.Field()] = msgForTag(fe)
				}
				SendError(c, BadRequest, out)
				c.Abort()
				return
			}

			// 4. Default Fallback (500)
			SendError(c, InternalServerError, err.Error())
			c.Abort()
		}
	}
}

// Helper to check for JSON errors
func isJSONError(err error) bool {
	var syntaxErr *json.SyntaxError
	var unmarshalErr *json.UnmarshalTypeError
	return errors.As(err, &syntaxErr) || errors.As(err, &unmarshalErr)
}

// Helper to translate validator tags to messages
func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Minimum length is %s", fe.Param())
	case "max":
		return fmt.Sprintf("Maximum length is %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	case "is_npwp":
		return "Invalid NPWP format (must be exactly 15 digits)"
	case "unique":
		return "This value already exists and must be unique"
	}

	return fe.Error()
}
