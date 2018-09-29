package handler

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func BadRequest(internalCode string) *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusBadRequest,
		Detail: internalCode,
		Message: "Bad request.",
	}
}

func Unauthorized(internalCode string) *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusUnauthorized,
		Detail: internalCode,
		Message: "Unauthorized.",
	}
}

func Forbidden(internalCode string) *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusForbidden,
		Detail: internalCode,
		Message: "Forbidden.",
	}
}

func NotFound(internalCode string) *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusNotFound,
		Detail: internalCode,
		Message: "Not found.",
	}
}

func Conflict(internalCode string) *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusConflict,
		Detail: internalCode,
		Message: "Conflict resource.",
	}
}

func UnProcessableEntity(internalCode string) *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusUnprocessableEntity,
		Detail: internalCode,
		Message: "UnProcessable Entity.",
	}
}

func InternalServerError(internalCode string) *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusInternalServerError,
		Detail: internalCode,
		Message: "Internal server error.",
	}
}

func (e ErrorResponse) Print(code int, msg string, detail string) {
	fmt.Print(code)
	fmt.Print(msg)
	fmt.Print(detail)
}