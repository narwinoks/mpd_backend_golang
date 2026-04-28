package exception

import "net/http"

type ForbiddenError struct {
	Message string
}

func (e ForbiddenError) Error() string {
	return e.Message
}

func NewForbiddenError(message string) ForbiddenError {
	return ForbiddenError{Message: message}
}

func IsForbiddenError(err error) bool {
	_, ok := err.(ForbiddenError)
	return ok
}

func (e ForbiddenError) HttpCode() int {
	return http.StatusForbidden
}
