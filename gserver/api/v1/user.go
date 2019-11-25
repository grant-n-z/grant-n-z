package v1

import (
	"encoding/json"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/common/ctx"
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var uhInstance User

type User interface {
	// Implement user api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	post(w http.ResponseWriter, r *http.Request)

	// Http PUT method
	put(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
	delete(w http.ResponseWriter, r *http.Request)
}

// User api struct
type UserImpl struct {
	Request     api.Request
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

// Constructor
func NewUser() User {
	log.Logger.Info("New `User` instance")
	log.Logger.Info("Inject `request`, `UserGroup`, `Service` to `User`")
	return UserImpl{
		Request:     api.GetRequestInstance(),
		UserService: service.GetUserServiceInstance(),
		Service:     service.GetServiceInstance(),
	}
}

func (uh UserImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		uh.get(w, r)
	case http.MethodPost:
		uh.post(w, r)
	case http.MethodPut:
		uh.put(w, r)
	case http.MethodDelete:
		uh.delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (uh UserImpl) get(w http.ResponseWriter, r *http.Request) {
}

func (uh UserImpl) post(w http.ResponseWriter, r *http.Request) {
	var userEntity *entity.User

	body, err := uh.Request.Intercept(w, r, "")
	if err != nil {
		return
	}

	json.Unmarshal(body, &userEntity)
	if err := uh.Request.ValidateBody(w, userEntity); err != nil {
		return
	}

	serviceEntity, err := uh.Service.GetServiceByApiKey(ctx.GetApiKey().(string))
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
		_, errorResponse = uh.UserService.InsertUserWithUserService(userEntity, userServiceEntity)
	}

	if errorResponse != nil {
		model.WriteError(w, errorResponse.ToJson(), errorResponse.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User creation succeeded."})
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (uh UserImpl) put(w http.ResponseWriter, r *http.Request) {
	var userEntity *entity.User

	body, err := uh.Request.Intercept(w, r, property.AuthUser)
	if err != nil {
		return
	}

	json.Unmarshal(body, &userEntity)
	if err := uh.Request.ValidateBody(w, userEntity); err != nil {
		return
	}

	userEntity.Id = ctx.GetUserId().(int)
	userEntity.Uuid = ctx.GetUserUuid().(uuid.UUID)
	if _, err := uh.UserService.UpdateUser(userEntity); err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"message": "User update succeeded."})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (uh UserImpl) delete(w http.ResponseWriter, r *http.Request) {
}
