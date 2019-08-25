package handler

import (
	"github.com/tomoyane/grant-n-z/gserver/model"
	"net/http"
)

type UserServiceHandler interface {
	Api(w http.ResponseWriter, r *http.Request)

	Get(w http.ResponseWriter, r *http.Request, authUser *model.AuthUser)

	Post(w http.ResponseWriter, r *http.Request, authUser *model.AuthUser)

	Put(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)
}
