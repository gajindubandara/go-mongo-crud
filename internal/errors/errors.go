package errors

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func InternalServerError(err error) *AppError {
	return NewAppError(http.StatusInternalServerError, "Internal Server Error", err)
}

func BadRequestError(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message, nil)
}

func NotFoundError(message string) *AppError {
	return NewAppError(http.StatusNotFound, message, nil)
}
