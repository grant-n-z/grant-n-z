package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ughInstance UserGroupHandler

type UserGroupHandlerImpl struct {
	RequestHandler RequestHandler
	AuthService    service.AuthService
}

func GetUserGroupHandlerInstance() UserGroupHandler {
	if ughInstance == nil {
		ughInstance = NewUserGroupHandler()
	}
	return ghInstance
}

func NewUserGroupHandler() UserGroupHandler {
	log.Logger.Info("New `GroupHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `AuthService` to `UserGroupHandler`")
	return UserGroupHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		AuthService:    service.GetAuthServiceInstance(),
	}
}

func (ugh UserGroupHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		ugh.Get(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (ugh UserGroupHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	_, err := ugh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
