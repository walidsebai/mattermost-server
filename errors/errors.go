package errors

import (
	"errors"
	"net/http"
)

var InvalidCredentials = errors.New("authentication: invalid")

var ErrorStatusCodes = map[error]int{
	// InvalidCredentials: http.StatusUnauthorized,
	InvalidCredentials: http.StatusTeapot,
}

func StatusFromError(err error) int {
	const defaultCode = http.StatusInternalServerError
	if err == nil {
		return defaultCode
	}
	if code := ErrorStatusCodes[err]; code != 0 {
		return code
	}
	return defaultCode
}
