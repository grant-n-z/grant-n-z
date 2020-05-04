package users

import (
	"encoding/json"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
	"net/http"
)

var uhInstance User

type User interface {
	// Http POST method.
	Post(w http.ResponseWriter, r *http.Request)

	// Http PUT method.
	Put(w http.ResponseWriter, r *http.Request)
}

// User api struct.
type UserImpl struct {
	UserService service.UserService
	Service     service.Service
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetUserInstance() User {
	if uhInstance == nil {
		uhInstance = NewUser()
	}
	return uhInstance
}

// Constructor.
func NewUser() User {
	log.Logger.Info("New `User` instance")
	return UserImpl{
		UserService: service.GetUserServiceInstance(),
		Service:     service.GetServiceInstance(),
	}
}

func (uh UserImpl) Post(w http.ResponseWriter, r *http.Request) {
	var userEntity *entity.User
	if err := middleware.BindBody(w, r, &userEntity); err != nil {
		return
	}

	userEntity.Username = uh.UserService.GenInitialName()
	if err := middleware.ValidateBody(w, userEntity); err != nil {
		return
	}

	serviceEntity, err := uh.Service.GetServiceOfSecret()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	userServiceEntity := &entity.UserService{
		UserId:    userEntity.Id,
		ServiceId: serviceEntity.Id,
	}

	if _, err = uh.UserService.InsertUserWithUserService(*userEntity, *userServiceEntity); err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User creation succeeded."})
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (uh UserImpl) Put(w http.ResponseWriter, r *http.Request) {
	var userEntity *entity.User
	if err := middleware.BindBody(w, r, &userEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, userEntity); err != nil {
		return
	}

	if _, err := uh.UserService.UpdateUser(*userEntity); err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User update succeeded."})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
