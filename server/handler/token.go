package handler

import "net/http"

type TokenHandler interface {
	Post(w http.ResponseWriter, r *http.Request)
}
