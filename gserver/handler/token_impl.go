package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var thInstance TokenHandler

type TokenHandlerImpl struct {
	RequestHandler RequestHandler
	TokenService   service.TokenService
}

func GetTokenHandlerInstance() TokenHandler {
	if thInstance == nil {
		thInstance = NewTokenHandler()
	}
	return thInstance
}

func NewTokenHandler() TokenHandler {
	log.Logger.Info("New `TokenHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `TokenService` to `TokenHandler`")
	return TokenHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		TokenService:   service.GetTokenServiceInstance(),
	}
}

func (th TokenHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		th.Post(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (th TokenHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST token")
	var userEntity *entity.User

	body, err := th.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &userEntity)
	userEntity.Username = userEntity.Email
	if err := th.RequestHandler.ValidateHttpRequest(w, userEntity); err != nil {
		return
	}

	token, err := th.TokenService.Generate(r.URL.Query().Get("type"), *userEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"token": *token})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
