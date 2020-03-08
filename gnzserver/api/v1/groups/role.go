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
	Get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	Post(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
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
	log.Logger.Info("New `Role` instance")
	return RoleImpl{RoleService: service.GetRoleServiceInstance()}
}
func (rh RoleImpl) Get(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.ParamGroupId(r)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	roles, err := rh.RoleService.GetRolesByGroupId(id)
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

	id, err := middleware.ParamGroupId(r)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	if err := middleware.BindBody(w, r, &roleEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, roleEntity); err != nil {
		return
	}

	role, err := rh.RoleService.InsertWithRelationalData(id, *roleEntity)
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
