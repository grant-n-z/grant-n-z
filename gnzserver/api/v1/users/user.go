package users

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var uhInstance User

type User interface {
	// Http POST method.
	// Endpoint is `/api/v1/users`
	Post(w http.ResponseWriter, r *http.Request)

	// Http PUT method.
	// Endpoint is `/api/v1/users`
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

	serviceEntity, err := uh.Service.GetServiceBySecret(r.Context().Value(middleware.ScopeSecret).(string))
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	userServiceEntity := &entity.UserService{
		UserUuid:    userEntity.Uuid,
		ServiceUuid: serviceEntity.Uuid,
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

	jwt := r.Context().Value(middleware.ScopeJwt).(model.JwtPayload)
	userUuid := uuid.MustParse(jwt.UserUuid)
	userEntity.Uuid = userUuid
	if _, err := uh.UserService.UpdateUser(*userEntity); err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User update succeeded."})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
