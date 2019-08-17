package handler

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/model"
)

type RequestHandler interface {
	InterceptHttp(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResponse)

	ValidateHttpRequest(w http.ResponseWriter, i interface{}) *model.ErrorResponse

	validateHttpHeader(r *http.Request) *model.ErrorResponse

	bindRequestBody(r *http.Request) ([]byte, *model.ErrorResponse)
}
