package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ughInstance UserGroup

type UserGroup interface {
	Api(w http.ResponseWriter, r *http.Request)

	post(w http.ResponseWriter, r *http.Request)
}

type UserGroupImpl struct {
	Request          api.Request
	UserGroupService service.UserGroupService
}

func GetUserGroupInstance() UserGroup {
	if ughInstance == nil {
		ughInstance = NewUserGroup()
	}
	return ughInstance
}

func NewUserGroup() UserGroup {
	log.Logger.Info("New `Group` instance")
	log.Logger.Info("Inject `Request`, `UserGroupService` to `UserGroup`")
	return UserGroupImpl{
		Request:          api.GetRequestInstance(),
		UserGroupService: service.GetUserGroupServiceInstance(),
	}
}

func (ugh UserGroupImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		ugh.post(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (ugh UserGroupImpl) post(w http.ResponseWriter, r *http.Request) {
	var userGroupEntity *entity.UserGroup

	body, err := ugh.Request.Intercept(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &userGroupEntity)
	if err := ugh.Request.ValidateBody(w, userGroupEntity); err != nil {
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
