package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var uInstance User

type User interface {
	// Implement permission api
	Api(w http.ResponseWriter, r *http.Request)

	// Http POST method
	// Add user
	postUserAdding(w http.ResponseWriter, r *http.Request)

	// Http PUT method
	// Update user's policy
	potUserPolicy(w http.ResponseWriter, r *http.Request, body []byte)
}

type UserImpl struct {
	// TODO
}

func GetUserInstance() User {
	if uInstance == nil {
		uInstance = NewUser()
	}
	return uInstance
}

func NewUser() User {
	log.Logger.Info("New `groups.User` instance")
	return UserImpl{}
}

func (u UserImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		u.postUserAdding(w, r)
	case http.MethodPut:
		u.postUserAdding(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (u UserImpl) postUserAdding(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(permissionEntities)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (u UserImpl) potUserPolicy(w http.ResponseWriter, r *http.Request, body []byte) {
	res, _ := json.Marshal(permission)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
