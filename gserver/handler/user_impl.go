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

var uhInstance UserHandler

type UserHandlerImpl struct {
	RequestHandler RequestHandler
	UserService    service.UserService
	Service        service.Service
}

func GetUserHandlerInstance() UserHandler {
	if uhInstance == nil {
		uhInstance = NewUserHandler()
	}
	return uhInstance
}

func NewUserHandler() UserHandler {
	log.Logger.Info("New `UserHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `UserGroup`, `Service` to `UserHandler`")
	return UserHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		UserService:    service.GetUserServiceInstance(),
		Service:        service.GetServiceInstance(),
	}
}

func (uh UserHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		uh.Get(w, r)
	case http.MethodPost:
		uh.Post(w, r)
	case http.MethodPut:
		uh.Put(w, r)
	case http.MethodDelete:
		uh.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (uh UserHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
}

func (uh UserHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	var userEntity *entity.User

	body, err := uh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &userEntity)
	if err := uh.RequestHandler.ValidateHttpRequest(w, userEntity); err != nil {
		return
	}

	serviceEntity, err := uh.Service.GetServiceByApiKey(r.Header.Get("Api-Key"))
	if err != nil {
		return
	}

	var errorResponse *model.ErrorResBody
	if serviceEntity == nil {
		_, errorResponse = uh.UserService.InsertUser(userEntity)
	} else {
		userServiceEntity := &entity.UserService{
			UserId: userEntity.Id,
			ServiceId: serviceEntity.Id,
		}
		_, errorResponse = uh.UserService.InsertUserWithService(userEntity, userServiceEntity)
	}

	if errorResponse != nil {
		model.Error(w, errorResponse.ToJson(), errorResponse.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User creation succeeded."})
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (uh UserHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
	var userEntity *entity.User

	authUser, err := uh.RequestHandler.VerifyToken(w, r, property.AuthUser)
	if err != nil {
		return
	}

	body, err := uh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &userEntity)
	if err := uh.RequestHandler.ValidateHttpRequest(w, userEntity); err != nil {
		return
	}

	userEntity.Id = authUser.UserId
	userEntity.Uuid = authUser.UserUuid
	if _, err := uh.UserService.UpdateUser(userEntity); err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User update succeeded."})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (uh UserHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
