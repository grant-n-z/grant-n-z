package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var uInstance User

type User interface {
	// Implement permission api
	// Endpoint is `/api/v1/groups/{id}/user`
	Api(w http.ResponseWriter, r *http.Request)

	// Http PUT method
	// Add user to group
	put(w http.ResponseWriter, r *http.Request)

	// Http GET method
	// get user of group
	get(w http.ResponseWriter, r *http.Request)
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
	return UserImpl{
		GroupService: service.GetGroupServiceInstance(),
		UserService:  service.GetUserServiceInstance(),
	}
}

func (u UserImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		u.put(w, r)
	case http.MethodGet:
		u.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (u UserImpl) put(w http.ResponseWriter, r *http.Request) {
	var addUserEntity *model.AddUser
	if err := middleware.BindBody(w, r, &addUserEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, addUserEntity); err != nil {
		return
	}

	group, errGroup := u.GroupService.GetGroupByUuid(middleware.ParamGroupUuid(r))
	if errGroup != nil {
		model.WriteError(w, errGroup.ToJson(), errGroup.Code)
		return
	}

	user, errUser := u.UserService.GetUserWithUserServiceWithServiceByEmail(addUserEntity.UserEmail)
	if errUser != nil {
		model.WriteError(w, errUser.ToJson(), errUser.Code)
		return
	}

	userGroupEntity := entity.UserGroup{
		UserUuid:  user.User.Uuid,
		GroupUuid: group.Uuid,
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

func (u UserImpl) get(w http.ResponseWriter, r *http.Request) {
	userResponse, errUser := u.UserService.GetUserByGroupUuid(middleware.ParamGroupUuid(r))
	if errUser != nil {
		model.WriteError(w, errUser.ToJson(), errUser.Code)
		return
	}

	res, _ := json.Marshal(userResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
