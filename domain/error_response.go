package domain

import (
	"net/http"
	"fmt"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e ErrorResponse) Error(code int, internalCode string) ErrorResponse {
	e.Code = code
	e.Detail = internalCode

	switch e.Code {
	case http.StatusBadRequest:
		e.Message = "Bad request."
	case http.StatusNotFound:
		e.Message = "Not found."
	case http.StatusUnprocessableEntity:
		e.Message = "Unprocessable Entity."
	case http.StatusInternalServerError:
		e.Message = "Internal server error."
	}

	return e
}

func (e ErrorResponse) Print(code int, msg string, detail string) {
	fmt.Print(code)
	fmt.Print(msg)
	fmt.Print(detail)
}