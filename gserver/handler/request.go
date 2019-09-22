package handler

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/model"
)

type RequestHandler interface {

	// Http interceptor
	// Verify http header
	// If it has request body, verify bind request body
	InterceptHttp(w http.ResponseWriter, r *http.Request) ([]byte, *model.ErrorResBody)

	// If need to verify authentication, verify authentication and authorization
	VerifyToken(w http.ResponseWriter, r *http.Request, authType string) (*model.AuthUser, *model.ErrorResBody)

	// Validate http request
	ValidateHttpRequest(w http.ResponseWriter, i interface{}) *model.ErrorResBody

	// Validate http header
	validateHttpHeader(r *http.Request) *model.ErrorResBody

	// Bind http request body
	bindRequestBody(r *http.Request) ([]byte, *model.ErrorResBody)
}
