package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
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
	log.Logger.Info("Inject `RequestHandler`, `userService` to `UserServiceHandler`")
	return UserServiceHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		UserService:    service.GetUserServiceServiceInstance(),
	}
}

func (ush UserServiceHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authUser, err := ush.RequestHandler.VerifyToken(w, r, property.AuthUser)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		ush.Get(w, r, authUser)
	case http.MethodPost:
		ush.Post(w, r, authUser)
	case http.MethodPut:
		ush.Put(w, r)
	case http.MethodDelete:
		ush.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (ush UserServiceHandlerImpl) Get(w http.ResponseWriter, r *http.Request, authUser *model.AuthUser) {
	id := r.URL.Query().Get(entity.UserServiceId.String())

	userServiceEntities, err := ush.UserService.Get(id)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(userServiceEntities)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ush UserServiceHandlerImpl) Post(w http.ResponseWriter, r *http.Request, authUser *model.AuthUser) {
	var userServiceEntity *entity.UserService

	body, err := ush.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &userServiceEntity)
	if err := ush.RequestHandler.ValidateHttpRequest(w, userServiceEntity); err != nil {
		return
	}

	userServiceEntity, err = ush.UserService.InsertUserService(userServiceEntity)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(userServiceEntity)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (ush UserServiceHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (ush UserServiceHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
