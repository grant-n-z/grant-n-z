package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnzserver/api"
	"github.com/tomoyane/grant-n-z/gnzserver/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var phInstance Permission

type Permission interface {
	// Http GET method
	Get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	Post(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
	Delete(w http.ResponseWriter, r *http.Request)
}

type PermissionImpl struct {
	Request           api.Request
	PermissionService service.PermissionService
}

func GetPermissionInstance() Permission {
	if phInstance == nil {
		phInstance = NewPermission()
	}
	return phInstance
}

func NewPermission() Permission {
	log.Logger.Info("New `Permission` instance")
	return PermissionImpl{
		Request:           api.GetRequestInstance(),
		PermissionService: service.GetPermissionServiceInstance(),
	}
}

func (ph PermissionImpl) Get(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.ParamGroupId(r)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	permissions, err := ph.PermissionService.GetPermissionsByGroupId(id)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(permissions)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ph PermissionImpl) Post(w http.ResponseWriter, r *http.Request) {
	var permissionEntity *entity.Permission

	id, err := middleware.ParamGroupId(r)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	if err := middleware.BindBody(w, r, &permissionEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, permissionEntity); err != nil {
		return
	}

	permission, err := ph.PermissionService.InsertWithRelationalData(id, *permissionEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(permission)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (ph PermissionImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
