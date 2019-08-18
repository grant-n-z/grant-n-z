package handler

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/model"
)

type RequestHandler interface {

	// Http interceptor
	// Verify http header
	// If it has request body, verify bind request body
	InterceptHttp(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResponse)

	// If need to verify authentication, verify authentication and authorization
	VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResponse)

	// Validate http request
	ValidateHttpRequest(w http.ResponseWriter, i interface{}) *model.ErrorResponse

	// Validate http header
	validateHttpHeader(r *http.Request) *model.ErrorResponse

	// Bind http request body
	bindRequestBody(r *http.Request) ([]byte, *model.ErrorResponse)
}
