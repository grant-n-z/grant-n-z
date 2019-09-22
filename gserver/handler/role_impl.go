package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var rlhInstance RoleHandler

type RoleHandlerImpl struct {
	RequestHandler RequestHandler
	RoleService    service.RoleService
}

func GetRoleHandlerInstance() RoleHandler {
	if rlhInstance == nil {
		rlhInstance = NewRoleHandler()
	}
	return rlhInstance
}

func NewRoleHandler() RoleHandler {
	log.Logger.Info("New `RoleHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `RoleService` to `RoleHandler`")
	return RoleHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		RoleService:    service.GetRoleServiceInstance(),
	}
}

func (rh RoleHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := rh.RequestHandler.VerifyToken(w, r, property.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		rh.Get(w, r)
	case http.MethodPost:
		rh.Post(w, r)
	case http.MethodPut:
		rh.Put(w, r)
	case http.MethodDelete:
		rh.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (rh RoleHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	roleEntities, err := rh.RoleService.GetRoles()
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (rh RoleHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	var roleEntity *entity.Role

	body, err := rh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &roleEntity)
	if err := rh.RequestHandler.ValidateHttpRequest(w, roleEntity); err != nil {
		return
	}

	roleEntity, err = rh.RoleService.InsertRole(roleEntity)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleEntity)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (rh RoleHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (rh RoleHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
