package entity

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BadRequest() *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusBadRequest,
		Message: "Bad request.",
	}
}

func Unauthorized() ErrorResponse {
	return ErrorResponse{
		Code: http.StatusUnauthorized,
		Message: "Unauthorized.",
	}
}

func Forbidden() ErrorResponse {
	return ErrorResponse{
		Code: http.StatusForbidden,
		Message: "Forbidden.",
	}
}

func NotFound() ErrorResponse {
	return ErrorResponse{
		Code: http.StatusNotFound,
		Message: "Not found.",
	}
}

func Conflict() *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusConflict,
		Message: "Conflict.",
	}
}

func MethodNotAllowed() *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusMethodNotAllowed,
		Message: "Method Not Allowed.",
	}
}

func UnProcessableEntity() ErrorResponse {
	return ErrorResponse{
		Code: http.StatusUnprocessableEntity,
		Message: "UnProcessable Entity.",
	}
}

func InternalServerError() *ErrorResponse {
	return &ErrorResponse{
		Code: http.StatusInternalServerError,
		Message: "Internal server error.",
	}
}

func (er ErrorResponse) ToJson() string {
	jsonBytes, _ := json.Marshal(er)
	return string(jsonBytes)
}
