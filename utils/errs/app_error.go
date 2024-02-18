package errs

import "fmt"

type AppError interface {
	Error() string
	GetHttpStatus() int
	GetCode() string
}

type Error struct {
	HttpStatus int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message,omitempty"`
}

func New(httpStatus int, code, message string) AppError {
	return &Error{
		HttpStatus: httpStatus,
		Code:       code,
		Message:    message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
}

func (e *Error) GetCode() string {
	return e.Code
}

func (e *Error) GetHttpStatus() int {
	return e.HttpStatus
}
