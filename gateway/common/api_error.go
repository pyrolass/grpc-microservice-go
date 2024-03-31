package common

import "net/http"

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (e ApiError) Error() string {
	return e.Message
}

func NewError(code int, message string) *ApiError {
	return &ApiError{Code: code, Message: message}
}

func InvalidID() ApiError {
	return ApiError{Code: http.StatusBadRequest, Message: "Invalid ID"}
}

func Unauthorized() ApiError {
	return ApiError{Code: http.StatusUnauthorized, Message: "Unauthorized"}
}
