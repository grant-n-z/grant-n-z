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

	post(w http.ResponseWriter, r *http.Request, body []byte)
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
	body, _, err := ugh.Request.Intercept(w, r, "")
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodPost:
		ugh.post(w, r, body)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (ugh UserGroupImpl) post(w http.ResponseWriter, r *http.Request, body []byte) {
	var userGroupEntity *entity.UserGroup

	json.Unmarshal(body, &userGroupEntity)
	if err := ugh.Request.ValidateBody(w, userGroupEntity); err != nil {
		return
	}

	userGroup, err := ugh.UserGroupService.InsertUserGroup(userGroupEntity)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(userGroup)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
