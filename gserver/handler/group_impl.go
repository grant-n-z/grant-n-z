package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ghInstance GroupHandler

type GroupHandlerImpl struct {
	RequestHandler RequestHandler
	AuthService    service.AuthService
}

func GetGroupHandlerInstance() GroupHandler {
	if ghInstance == nil {
		ghInstance = NewGroupHandler()
	}
	return ghInstance
}

func NewGroupHandler() GroupHandler {
	log.Logger.Info("New `GroupHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `AuthService` to `GroupHandler`")
	return GroupHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		AuthService:    service.GetAuthServiceInstance(),
	}
}

func (gh GroupHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		gh.Get(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (gh GroupHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	_, err := gh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (gh GroupHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	_, err := gh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (gh GroupHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
