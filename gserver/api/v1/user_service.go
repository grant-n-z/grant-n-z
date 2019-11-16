package v1

import (
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ushInstance UserService

type UserService interface {
	// Implement user service api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)
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
	log.Logger.Info("Inject `request`, `userService` to `UserService`")
	return UserServiceImpl{
		Request:     api.GetRequestInstance(),
		UserService: service.GetUserServiceServiceInstance(),
	}
}

func (ush UserServiceImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := ush.Request.Intercept(w, r, property.AuthUser)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		ush.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (ush UserServiceImpl) get(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get(entity.UserServiceId.String())
	//res, _ := json.Marshal(userServiceEntities)
	//w.WriteHeader(http.StatusOK)
	//w.Write(res)
}
