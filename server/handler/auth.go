package handler

import "net/http"

type AuthHandler interface {
	Api(w http.ResponseWriter, r *http.Request)

	Post(w http.ResponseWriter, r *http.Request)
}

