package handler

import "net/http"

type TokenHandler interface {
	Api(w http.ResponseWriter, r *http.Request)

	Post(w http.ResponseWriter, r *http.Request)
}
