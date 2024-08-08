package common

import (
	"errors"
	"net/http"
)

type AppError struct{
	StatusCode int `json:"status_code"`
	RootError error `json:"-"`
	Message string `json:"message"`
	Key string `json:"key"`
}

func (e *AppError)Error() string{
	return e.RootError.Error()
}

func NewUnauthenticatedError(rootError error, message string) *AppError{
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootError: rootError,
		Message: message,
		Key: "AUTH_ERROR",
	}
}

var ErrRecordNotFound error = errors.New("Record Not Found")
