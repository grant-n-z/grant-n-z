package model

import (
	"encoding/json"
	"net/http"
)

// GrantNZ error data
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// To json
func (er ErrorResponse) ToJson() string {
	jsonBytes, _ := json.Marshal(er)
	return string(jsonBytes)
}

// BadRequest
func BadRequest(err ...string) *ErrorResponse {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad request.",
		Detail:  detail,
	}
}

// Unauthorized
func Unauthorized(err ...string) *ErrorResponse {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized.",
		Detail:  detail,
	}
}

// Forbidden
func Forbidden(err ...string) ErrorResponse {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return ErrorResponse{
		Code:    http.StatusForbidden,
		Message: "Forbidden.",
		Detail:  detail,
	}
}

// NotFound
func NotFound(err ...string) *ErrorResponse {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "Not found.",
		Detail:  detail,
	}
}

// Conflict
func Conflict(err ...string) *ErrorResponse {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResponse{
		Code:    http.StatusConflict,
		Message: "Conflict.",
		Detail:  detail,
	}
}

// MethodNotAllowed
func MethodNotAllowed(err ...string) *ErrorResponse {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResponse{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method Not Allowed.",
		Detail:  detail,
	}
}

// UnProcessableEntity
func UnProcessableEntity(err ...string) ErrorResponse {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return ErrorResponse{
		Code:    http.StatusUnprocessableEntity,
		Message: "UnProcessable Entity.",
		Detail:  detail,
	}
}

// InternalServerError
func InternalServerError(err ...string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error.",
		Detail:  "Error internal processing.",
	}
}
