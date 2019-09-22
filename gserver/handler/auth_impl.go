package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ahInstance AuthHandler

type AuthHandlerImpl struct {
	RequestHandler RequestHandler
	AuthService    service.AuthService
}

func GetAuthHandlerInstance() AuthHandler {
	if ahInstance == nil {
		ahInstance = NewAuthHandler()
	}
	return ahInstance
}

func NewAuthHandler() AuthHandler {
	log.Logger.Info("New `AuthHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `AuthService` to `AuthHandler`")
	return AuthHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		AuthService:    service.GetAuthServiceInstance(),
	}
}

func (ah AuthHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		ah.Get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (ah AuthHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	_, err := ah.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
