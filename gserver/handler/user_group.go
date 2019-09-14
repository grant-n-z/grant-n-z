package handler

import "net/http"

type UserGroupHandler interface {
	Api(w http.ResponseWriter, r *http.Request)

	Post(w http.ResponseWriter, r *http.Request)
}

