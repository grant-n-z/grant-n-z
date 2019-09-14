package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var sghInstance ServiceGroupHandler

type ServiceGroupHandlerImpl struct {
	RequestHandler RequestHandler
	AuthService    service.AuthService
}

func GetServiceGroupHandlerInstance() ServiceGroupHandler {
	if sghInstance == nil {
		sghInstance = NewServiceGroupHandler()
	}
	return sghInstance
}

func NewServiceGroupHandler() ServiceGroupHandler {
	log.Logger.Info("New `ServiceGroupHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `AuthService` to `ServiceGroupHandler`")
	return ServiceGroupHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		AuthService:    service.GetAuthServiceInstance(),
	}
}

func (sgh ServiceGroupHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		sgh.Get(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (sgh ServiceGroupHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	_, err := sgh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
