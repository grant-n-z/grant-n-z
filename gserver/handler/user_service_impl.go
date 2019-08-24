package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ushInstance UserServiceHandler

type UserServiceHandlerImpl struct {
	RequestHandler RequestHandler
	UserService    service.UserServiceService
}

func GetUserServiceHandlerInstance() UserServiceHandler {
	if ushInstance == nil {
		ushInstance = NewUserServiceHandler()
	}
	return ushInstance
}

func NewUserServiceHandler() UserServiceHandler {
	log.Logger.Info("New `UserServiceHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `UserService` to `UserServiceHandler`")
	return UserServiceHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		UserService:    service.NewUserServiceService(),
	}
}

func (ush UserServiceHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		ush.Get(w, r)
	case http.MethodPost:
		ush.Post(w, r)
	case http.MethodPut:
		ush.Put(w, r)
	case http.MethodDelete:
		ush.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (ush UserServiceHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET user_services")
	id := r.URL.Query().Get(entity.UserServiceId.String())

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

	body, err := ush.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &userServiceEntity)
	if err := ush.RequestHandler.ValidateHttpRequest(w, userServiceEntity); err != nil {
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
