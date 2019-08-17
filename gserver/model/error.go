package model

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (er ErrorResponse) ToJson() string {
	jsonBytes, _ := json.Marshal(er)
	return string(jsonBytes)
}

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

func InternalServerError(err ...string) *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error.",
		Detail:  "Error internal processing.",
	}
}
