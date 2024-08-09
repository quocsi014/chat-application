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

func ErrUnauthenticatedError(rootError error, message string) *AppError{
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootError: rootError,
		Message: message,
		Key: "AUTH_ERROR",
	}
}

func ErrInternal(rootError error) *AppError{
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		RootError: rootError,
		Message: "Something went wrong, pls come back later",
		Key: "INTERNAL_ERROR",
	}
}

func ErrInvalidRequest(rootError error) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootError: rootError,
		Message: "invalid request",
		Key: "INVALID_REQ_ERROR",
	}
}
var ErrRecordNotFound error = errors.New("Record Not Found")
