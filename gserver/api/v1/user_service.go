package v1

import (
	"encoding/json"
	"github.com/tomoyane/grant-n-z/gserver/api"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ushInstance UserService

type UserService interface {
	Api(w http.ResponseWriter, r *http.Request)

	get(w http.ResponseWriter, r *http.Request, authUser *model.AuthUser)

	post(w http.ResponseWriter, r *http.Request, authUser *model.AuthUser)

	put(w http.ResponseWriter, r *http.Request)

	delete(w http.ResponseWriter, r *http.Request)
}

type UserServiceImpl struct {
	Request     api.Request
	UserService service.UserServiceService
}

func GetUserServiceInstance() UserService {
	if ushInstance == nil {
		ushInstance = NewUserService()
	}
	return ushInstance
}

func NewUserService() UserService {
	log.Logger.Info("New `UserService` instance")
	log.Logger.Info("Inject `Request`, `userService` to `UserService`")
	return UserServiceImpl{
		Request:     api.GetRequestInstance(),
		UserService: service.GetUserServiceServiceInstance(),
	}
}

func (ush UserServiceImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authUser, err := ush.Request.VerifyToken(w, r, property.AuthUser)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		ush.get(w, r, authUser)
	case http.MethodPost:
		ush.post(w, r, authUser)
	case http.MethodPut:
		ush.put(w, r)
	case http.MethodDelete:
		ush.delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (ush UserServiceImpl) get(w http.ResponseWriter, r *http.Request, authUser *model.AuthUser) {
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

func (ush UserServiceImpl) post(w http.ResponseWriter, r *http.Request, authUser *model.AuthUser) {
	var userServiceEntity *entity.UserService

	body, err := ush.Request.Intercept(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &userServiceEntity)
	if err := ush.Request.ValidateBody(w, userServiceEntity); err != nil {
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

func (ush UserServiceImpl) put(w http.ResponseWriter, r *http.Request) {
}

func (ush UserServiceImpl) delete(w http.ResponseWriter, r *http.Request) {
}
