package v1

import (
	"encoding/json"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var sInstance Service

type Service interface {
	// Http GET method
	// All service data
	// Endpoint is `/api/v1/services`
	Get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	// Add user to not main service
	// Endpoint is `/api/v1/services/add_user`
	Post(w http.ResponseWriter, r *http.Request)
}

type ServiceImpl struct {
	ServiceService service.Service
	UserService    service.UserService
	TokenProcessor middleware.TokenProcessor
}

func GetServiceInstance() Service {
	if sInstance == nil {
		sInstance = NewService()
	}
	return sInstance
}

func NewService() Service {
	log.Logger.Info("New `v1.Service` instance")
	return ServiceImpl{
		ServiceService: service.GetServiceInstance(),
		UserService:    service.GetUserServiceInstance(),
		TokenProcessor: middleware.GetTokenProcessorInstance(),
	}
}

func (s ServiceImpl) Get(w http.ResponseWriter, r *http.Request) {
	services, err := s.ServiceService.GetServices()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
	}

	res, _ := json.Marshal(services)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (s ServiceImpl) Post(w http.ResponseWriter, r *http.Request) {
	var tokenRequest *model.TokenRequest
	if err := middleware.BindBody(w, r, &tokenRequest); err != nil {
		return
	}
	if err := middleware.ValidateTokenRequest(w, tokenRequest); err != nil {
		return
	}

	// Authentication user
	token, err := s.TokenProcessor.Generate(common.AuthUser, *tokenRequest)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	jwt, err := s.TokenProcessor.GetJwtPayload("Bearer " + token.Token, false)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	// Check exist service
	serviceEntity, err := s.ServiceService.GetServiceBySecret(r.Context().Value(middleware.ScopeSecret).(string))
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	// Insert user_services
	userServiceEntity := &entity.UserService{
		UserUuid:    jwt.UserUuid,
		ServiceUuid: serviceEntity.Uuid,
	}
	userService, err := s.UserService.InsertUserService(*userServiceEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(userService)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
