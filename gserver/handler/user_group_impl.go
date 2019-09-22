package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ughInstance UserGroupHandler

type UserGroupHandlerImpl struct {
	RequestHandler   RequestHandler
	UserGroupService service.UserGroupService
	AuthService      service.AuthService
}

func GetUserGroupHandlerInstance() UserGroupHandler {
	if ughInstance == nil {
		ughInstance = NewUserGroupHandler()
	}
	return ughInstance
}

func NewUserGroupHandler() UserGroupHandler {
	log.Logger.Info("New `GroupHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `AuthService`, `UserGroupService` to `UserGroupHandler`")
	return UserGroupHandlerImpl{
		RequestHandler:   GetRequestHandlerInstance(),
		UserGroupService: service.GetUserGroupServiceInstance(),
		AuthService:      service.GetAuthServiceInstance(),
	}
}

func (ugh UserGroupHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		ugh.Post(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (ugh UserGroupHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	var userGroupEntity *entity.UserGroup

	body, err := ugh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &userGroupEntity)
	if err := ugh.RequestHandler.ValidateHttpRequest(w, userGroupEntity); err != nil {
		return
	}

	userGroupEntity, err = ugh.UserGroupService.InsertUserGroup(userGroupEntity)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(userGroupEntity)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
