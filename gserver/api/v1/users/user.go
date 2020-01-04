package users

import (
	"encoding/json"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/middleware"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
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
	log.Logger.Info("Inject `UserGroup`, `Service` to `User`")
	return UserImpl{
		UserService: service.GetUserServiceInstance(),
		Service:     service.GetServiceInstance(),
	}
}

func (uh UserImpl) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	var userEntity *entity.User
	if err := middleware.BindBody(w, r, &userEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, userEntity); err != nil {
		return
	}

	serviceEntity, err := uh.Service.GetServiceByApiKey(ctx.GetApiKey().(string))
	if err != nil {
		return
	}

	var errorResponse *model.ErrorResBody
	if serviceEntity == nil {
		_, errorResponse = uh.UserService.InsertUser(*userEntity)

	} else {
		userServiceEntity := &entity.UserService{
			UserId:    userEntity.Id,
			ServiceId: serviceEntity.Id,
		}
		_, errorResponse = uh.UserService.InsertUserWithUserService(*userEntity, *userServiceEntity)
	}

	if errorResponse != nil {
		model.WriteError(w, errorResponse.ToJson(), errorResponse.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User creation succeeded."})
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (uh UserImpl) Put(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	var userEntity *entity.User
	if err := middleware.BindBody(w, r, &userEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, userEntity); err != nil {
		return
	}

	userEntity.Id = ctx.GetUserId().(int)
	userEntity.Uuid = ctx.GetUserUuid().(uuid.UUID)
	if _, err := uh.UserService.UpdateUser(*userEntity); err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User update succeeded."})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
