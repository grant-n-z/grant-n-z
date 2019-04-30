package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/service"
	"github.com/tomoyane/grant-n-z/server/log"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler() UserHandler {
	return UserHandler{UserService: service.NewUserService()}
}

func (uh UserHandler) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST users")
	var userEntity *entity.User

	body, err := Interceptor(w, r, http.MethodPost)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &userEntity)
	if err := BodyValidator(w, userEntity); err != nil {
		return
	}

	if _, err := uh.UserService.InsertUser(userEntity); err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string {"message": "user creation succeeded."})
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}
