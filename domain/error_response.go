package domain

import "net/http"

type ErrorResponse struct {
	Code    int
	Message string
	Detail  string
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