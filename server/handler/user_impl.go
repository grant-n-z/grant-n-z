package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler() UserHandler {
	log.Logger.Info("Inject `UserService` to `UserHandler`")
	return UserHandlerImpl{UserService: service.NewUserService()}
}

func (uh UserHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: uh.Get(w, r)
	case http.MethodPost: uh.Post(w, r)
	case http.MethodPut: uh.Put(w, r)
	case http.MethodDelete: uh.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (uh UserHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
}

func (uh UserHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST users")
	var userEntity *entity.User

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &userEntity)
	if err := ValidateHttpRequest(w, userEntity); err != nil {
		return
	}

	if _, err := uh.UserService.InsertUser(userEntity); err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string {"message": "user creation succeeded."})
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (uh UserHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (uh UserHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
