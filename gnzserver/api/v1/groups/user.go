package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var uInstance User

type User interface {
	// Implement permission api
	Api(w http.ResponseWriter, r *http.Request)

	// Http PUT method
	// Add user to group
	put(w http.ResponseWriter, r *http.Request)
}

type UserImpl struct {
	GroupService service.GroupService
	UserService  service.UserService
}

func GetUserInstance() User {
	if uInstance == nil {
		uInstance = NewUser()
	}
	return uInstance
}

func NewUser() User {
	log.Logger.Info("New `groups.User` instance")
	return UserImpl{
		GroupService: service.GetGroupServiceInstance(),
		UserService:  service.GetUserServiceInstance(),
	}
}

func (u UserImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		u.put(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (u UserImpl) put(w http.ResponseWriter, r *http.Request) {
	var addUserEntity *entity.AddUser
	if err := middleware.BindBody(w, r, &addUserEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, addUserEntity); err != nil {
		return
	}

	id, err := middleware.ParamGroupId(r)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	group, errGroup := u.GroupService.GetGroupById(id)
	if errGroup != nil {
		model.WriteError(w, errGroup.ToJson(), errGroup.Code)
		return
	}

	user, errUser := u.UserService.GetUserWithUserServiceWithServiceByEmail(addUserEntity.UserEmail)
	if errUser != nil || user == nil {
		model.WriteError(w, errUser.ToJson(), errUser.Code)
		return
	}

	userGroupEntity := entity.UserGroup{
		UserId:  user.User.Id,
		GroupId: group.Id,
	}
	userGroup, errUserGroup := u.UserService.InsertUserGroup(userGroupEntity)
	if errUserGroup != nil {
		model.WriteError(w, errUserGroup.ToJson(), errUserGroup.Code)
		return
	}

	res, _ := json.Marshal(userGroup)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
