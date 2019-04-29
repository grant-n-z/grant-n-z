package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/service"
	"github.com/tomoyane/grant-n-z/server/log"
)

var logger log.Log

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler() UserHandler {
	logger = log.NewLogger()
	return UserHandler{UserService: service.NewUserService()}
}

func (uh UserHandler) Post(w http.ResponseWriter, r *http.Request) {
	logger.Info("POST users")

	bodyBytes, err := Interceptor(r)
	if err != nil {
		logger.Error("error interceptor")
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	var user entity.User
	json.Unmarshal(bodyBytes, &user)
	if err := BodyValidator(user); err != nil {
		logger.Info(err.ToJson())
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	if _, err := uh.UserService.InsertUser(user); err != nil {
		logger.Info(err.ToJson())
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	ok, _ := json.Marshal(map[string]string {"message": "user creation succeeded."})
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(ok)
}
