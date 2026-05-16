package errors

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrBadRequest          = NewAppError(http.StatusBadRequest, "Dữ liệu đầu vào không hợp lệ")
	ErrUnauthorized        = NewAppError(http.StatusUnauthorized, "Không có quyền truy cập")
	ErrInternalServerError = NewAppError(http.StatusInternalServerError, "Lỗi hệ thống")
)
