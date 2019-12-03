package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var phInstance Permission

type Permission interface {
	// Implement permission api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	post(w http.ResponseWriter, r *http.Request, body []byte)

	// Http PUT method
	put(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
	delete(w http.ResponseWriter, r *http.Request)
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
	log.Logger.Info("Inject `request`, `PermissionService` to `Permission`")
	return PermissionImpl{
		Request:           api.GetRequestInstance(),
		PermissionService: service.GetPermissionServiceInstance(),
	}
}

func (ph PermissionImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ph.Request.Intercept(w, r, property.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		ph.get(w, r)
	case http.MethodPost:
		ph.post(w, r, body)
	case http.MethodPut:
		ph.put(w, r)
	case http.MethodDelete:
		ph.delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (ph PermissionImpl) get(w http.ResponseWriter, r *http.Request) {
	permissionEntities, err := ph.PermissionService.GetPermissions()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(permissionEntities)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (ph PermissionImpl) post(w http.ResponseWriter, r *http.Request, body []byte) {
	var permissionEntity *entity.Permission

	json.Unmarshal(body, &permissionEntity)
	if err := ph.Request.ValidateBody(w, permissionEntity); err != nil {
		return
	}

	permission, err := ph.PermissionService.InsertPermission(permissionEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(permission)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (ph PermissionImpl) put(w http.ResponseWriter, r *http.Request) {
}

func (ph PermissionImpl) delete(w http.ResponseWriter, r *http.Request) {
}
