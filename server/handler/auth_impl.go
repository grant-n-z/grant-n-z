package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type AuthHandlerImpl struct {
	RequestHandler RequestHandler
	UserService    service.UserService
}

func NewAuthHandler() AuthHandler {
	log.Logger.Info("Inject `RequestHandler`, `UserService` to `TokenHandler`")
	return AuthHandlerImpl{
		RequestHandler: NewRequestHandler(),
		UserService: service.NewUserService(),
	}
}

func (ah AuthHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		ah.Post(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (ah AuthHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST oauth")
	var userEntity *entity.User

	body, err := ah.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &userEntity)
	userEntity.Username = userEntity.Email
	if err := ah.RequestHandler.ValidateHttpRequest(w, userEntity); err != nil {
		return
	}

	user, err := ah.UserService.GetUserByEmail(userEntity.Email)
	if err != nil || user == nil {
		errResponse := model.BadRequest("Failed to email or password")
		http.Error(w, errResponse.ToJson(), errResponse.Code)
		return
	}

	if !ah.UserService.ComparePw(user.Password, userEntity.Password) {
		errResponse := model.BadRequest("Failed to email or password")
		http.Error(w, errResponse.ToJson(), errResponse.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"token": *ah.UserService.GenerateJwt(user, "test")})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
