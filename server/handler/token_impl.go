package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type TokenHandlerImpl struct {
	UserService service.UserService
}

func NewTokenHandler() TokenHandler {
	log.Logger.Info("Inject `UserService` to `TokenHandler`")
	return TokenHandlerImpl{UserService: service.NewUserService()}
}

func (th TokenHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST oauth")
	var userEntity *entity.User

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &userEntity)
	userEntity.Username = userEntity.Email
	if err := BodyValidator(w, userEntity); err != nil {
		return
	}

	user, err := th.UserService.GetUserByEmail(userEntity.Email)
	if err != nil || user == nil {
		errResponse := model.BadRequest("Failed to email or password")
		http.Error(w, errResponse.ToJson(), errResponse.Code)
		return
	}

	if !th.UserService.ComparePw(user.Password, userEntity.Password) {
		errResponse := model.BadRequest("Failed to email or password")
		http.Error(w, errResponse.ToJson(), errResponse.Code)
		return
	}

	res, _ := json.Marshal(map[string]string {"token": *th.UserService.GenerateJwt(user, "test")})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
