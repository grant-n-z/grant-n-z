package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type UserServiceHandlerImpl struct {
	UserService service.UserServiceService
}

func NewUserServiceHandler() UserServiceHandler {
	log.Logger.Info("Inject `UserServiceService` to `UserServiceHandler`")
	return UserServiceHandlerImpl{UserService: service.NewUserServiceService()}
}

func (ush UserServiceHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: ush.Get(w, r)
	case http.MethodPost: ush.Post(w, r)
	case http.MethodPut: ush.Put(w, r)
	case http.MethodDelete: ush.Delete(w, r)
	default:
		err := entity.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (ush UserServiceHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET user_services")
	id := r.URL.Query().Get(entity.USER_SERVICE_ID.String())

	userServiceEntities, err := ush.UserService.Get(id)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(userServiceEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (ush UserServiceHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST user_services")
	var userServiceEntity *entity.UserService

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &userServiceEntity)
	if err := BodyValidator(w, userServiceEntity); err != nil {
		return
	}

	userServiceEntity, err = ush.UserService.InsertUserService(userServiceEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(userServiceEntity)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (ush UserServiceHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (ush UserServiceHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
