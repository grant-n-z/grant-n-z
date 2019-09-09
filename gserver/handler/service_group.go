package handler

import "net/http"

type ServiceGroupHandler interface {
	Api(w http.ResponseWriter, r *http.Request)

	Get(w http.ResponseWriter, r *http.Request)
}

