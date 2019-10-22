package v1

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

var rlhInstance Role

type Role interface {
	// Implement role api
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

type RoleImpl struct {
	Request     api.Request
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
	log.Logger.Info("Inject `request`, `RoleService` to `Role`")
	return RoleImpl{
		Request:     api.GetRequestInstance(),
		RoleService: service.GetRoleServiceInstance(),
	}
}

func (rh RoleImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := rh.Request.Intercept(w, r, property.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		rh.get(w, r)
	case http.MethodPost:
		rh.post(w, r, body)
	case http.MethodPut:
		rh.put(w, r)
	case http.MethodDelete:
		rh.delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (rh RoleImpl) get(w http.ResponseWriter, r *http.Request) {
	roleEntities, err := rh.RoleService.GetRoles()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (rh RoleImpl) post(w http.ResponseWriter, r *http.Request, body []byte) {
	var roleEntity *entity.Role

	json.Unmarshal(body, &roleEntity)
	if err := rh.Request.ValidateBody(w, roleEntity); err != nil {
		return
	}

	role, err := rh.RoleService.InsertRole(roleEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(role)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (rh RoleImpl) put(w http.ResponseWriter, r *http.Request) {
}

func (rh RoleImpl) delete(w http.ResponseWriter, r *http.Request) {
}
