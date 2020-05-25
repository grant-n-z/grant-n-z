package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var rlhInstance Role

type Role interface {
	// Http GET method
	// Endpoint is `/api/v1/groups/{id}/role`
	Get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	// Endpoint is `/api/v1/groups/{id}/role`
	Post(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
	// Endpoint is `/api/v1/groups/{id}/role`
	Delete(w http.ResponseWriter, r *http.Request)
}

type RoleImpl struct {
	RoleService service.RoleService
}

func GetRoleInstance() Role {
	if rlhInstance == nil {
		rlhInstance = NewRole()
	}
	return rlhInstance
}

func NewRole() Role {
	log.Logger.Info("New `v1.groups.Role` instance")
	return RoleImpl{RoleService: service.GetRoleServiceInstance()}
}

func (rh RoleImpl) Get(w http.ResponseWriter, r *http.Request) {
	roles, err := rh.RoleService.GetRolesByGroupUuid(middleware.ParamGroupUuid(r))
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roles)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (rh RoleImpl) Post(w http.ResponseWriter, r *http.Request) {
	var roleEntity *entity.Role

	if err := middleware.BindBody(w, r, &roleEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, roleEntity); err != nil {
		return
	}

	role, err := rh.RoleService.InsertWithRelationalData(middleware.ParamGroupUuid(r), *roleEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(role)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (rh RoleImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
