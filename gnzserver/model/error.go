package model

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/log"
)

// GrantNZ error data
type ErrorResBody struct {
	Code      int    `json:"code"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

// To json
func (er ErrorResBody) ToJson() string {
	jsonBytes, _ := json.Marshal(er)
	return string(jsonBytes)
}

// WriteError response
func WriteError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(code)
	w.Write([]byte(error))
}

// Option
func Options() *ErrorResBody {
	return &ErrorResBody{
		Code: http.StatusNoContent,
	}
}

// BadRequest
func BadRequest(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}
	return &ErrorResBody{
		Code:    http.StatusBadRequest,
		Title:   "Bad request.",
		Message: detail,
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
		Title:   "Unauthorized.",
		Message: detail,
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
		Title:   "Forbidden.",
		Message: detail,
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
		Title:   "Not found.",
		Message: detail,
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
		Title:   "Conflict.",
		Message: detail,
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
		Title:   "Method Not Allowed.",
		Message: detail,
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
		Title:   "UnProcessable Entity.",
		Message: detail,
	}
}

// InternalServerError
func InternalServerError(err ...string) *ErrorResBody {
	var detail string
	if err != nil {
		detail = err[0]
	}

	body := ErrorResBody{
		Code:    http.StatusInternalServerError,
		Title:   "Internal server error.",
		Message: detail,
	}
	log.Logger.Error(body.ToJson())
	return &body
}
