package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/usecase/service"
)

type AuthHandlerImpl struct {
	RequestHandler RequestHandler
	AuthService    service.AuthService
}

func NewAuthHandler() AuthHandler {
	log.Logger.Info("Inject `RequestHandler`, `UserService` to `TokenHandler`")
	return AuthHandlerImpl{
		RequestHandler: NewRequestHandler(),
		AuthService: service.NewAuthService(),
	}
}

func (ah AuthHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		ah.Get(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (ah AuthHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET auth")

	_, err := ah.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
