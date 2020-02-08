package model

import (
	"encoding/json"
	"net/http"
)

// GrantNZ error data
type ErrorResBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// To json
func (er ErrorResBody) ToJson() string {
	jsonBytes, _ := json.Marshal(er)
	return string(jsonBytes)
}

// WriteError response
func WriteError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(error))
}

// BadRequest
func BadRequest(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResBody{
		Code:    http.StatusBadRequest,
		Message: "Bad request.",
		Detail:  detail,
	}
}

// Unauthorized
func Unauthorized(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResBody{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized.",
		Detail:  detail,
	}
}

// Forbidden
func Forbidden(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResBody{
		Code:    http.StatusForbidden,
		Message: "Forbidden.",
		Detail:  detail,
	}
}

// NotFound
func NotFound(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResBody{
		Code:    http.StatusNotFound,
		Message: "Not found.",
		Detail:  detail,
	}
}

// Conflict
func Conflict(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResBody{
		Code:    http.StatusConflict,
		Message: "Conflict.",
		Detail:  detail,
	}
}

// MethodNotAllowed
func MethodNotAllowed(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResBody{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method Not Allowed.",
		Detail:  detail,
	}
}

// UnProcessableEntity
func UnProcessableEntity(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResBody{
		Code:    http.StatusUnprocessableEntity,
		Message: "UnProcessable Entity.",
		Detail:  detail,
	}
}

// InternalServerError
func InternalServerError(err ...string) *ErrorResBody {
	return &ErrorResBody{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error.",
		Detail:  "Failed to internal processing.",
	}
}
