package handler

import (
	"net/http"
)

type ServiceHandler interface {
	Api(w http.ResponseWriter, r *http.Request)

	Get(w http.ResponseWriter, r *http.Request)

	Post(w http.ResponseWriter, r *http.Request)

	Put(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)
}
