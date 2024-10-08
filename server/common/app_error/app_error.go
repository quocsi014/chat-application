package app_error

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

func (e AppError)Error() string{
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
		Message: "Something went wrong in server, pls come back later",
		Key: "INTERNAL_ERROR",
	}
}

func ErrDatabase(rootError error) *AppError{
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		RootError: rootError,
		Message: "Something went wrong with database",
		Key: "DB_ERROR",
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
func ErrInvalidData(rootError error, key, message string) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootError: rootError,
		Message: message,
		Key: key,
	}
}
func ErrConflictData(rootError error, key string, message string) *AppError{
	return &AppError{
		StatusCode: http.StatusConflict,
		RootError: rootError,
		Message: message,
		Key: key,
	}
}

func ErrPermissionDenied() *AppError{
	return &AppError{
		StatusCode: http.StatusForbidden,
		RootError: errors.New("permission denied"),
		Message: "Yor are not allowed",
		Key: "NOT_ALLOWED",
	}
}

var ErrRecordNotFound error = errors.New("record rot round")

type ErrorResponse struct{
	Code int
	Err error
}
func NewErrorResponseWithAppError(err error) *ErrorResponse{
	appError := err.(*AppError)
	return &ErrorResponse{
		Code: appError.StatusCode,
		Err: appError,
	}
}

func ErrNotFound(rootError error, key string, message string) *AppError {
    return &AppError{
        StatusCode: http.StatusNotFound,
        RootError:  rootError,
        Message:    message,
        Key:        key,
    }
}
