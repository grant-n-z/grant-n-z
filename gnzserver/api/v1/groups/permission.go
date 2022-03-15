package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var phInstance Permission

type Permission interface {
	// Http GET method
	// Endpoint is `/api/v1/groups/{group_uuid}/permission`
	Get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	// Endpoint is `/api/v1/groups/{group_uuid}/permission`
	Post(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
	// Endpoint is `/api/v1/groups/{group_uuid}/permission`
	Delete(w http.ResponseWriter, r *http.Request)
}

type PermissionImpl struct {
	PermissionService service.PermissionService
}

func GetPermissionInstance() Permission {
	if phInstance == nil {
		phInstance = NewPermission()
	}
	return phInstance
}

func NewPermission() Permission {
	return PermissionImpl{
		PermissionService: service.GetPermissionServiceInstance(),
	}
}

func (ph PermissionImpl) Get(w http.ResponseWriter, r *http.Request) {
	groupUuid := middleware.ParamGroupUuid(r)
	permissions, err := ph.PermissionService.GetPermissionsByGroupUuid(groupUuid)
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

	if err := middleware.BindBody(w, r, &permissionEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, permissionEntity); err != nil {
		return
	}

	permission, err := ph.PermissionService.InsertWithRelationalData(middleware.ParamGroupUuid(r), *permissionEntity)
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
