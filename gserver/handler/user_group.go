package handler

import "net/http"

type UserGroupHandler interface {
	Api(w http.ResponseWriter, r *http.Request)

	Get(w http.ResponseWriter, r *http.Request)
}

