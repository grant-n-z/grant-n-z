package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/service"
	"github.com/tomoyane/grant-n-z/server/log"
)

type RoleHandler struct {
	RoleService service.RoleService
	Service     service.Service
}

func NewRoleHandler() RoleHandler {
	log.Logger.Debug("inject `Service`, `RoleService` to `RoleHandler`")
	return RoleHandler{
		RoleService: service.NewRoleService(),
		Service:     service.NewServiceService(),
	}
}

func (rh RoleHandler) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: rh.Get(w, r)
	case http.MethodPost: rh.Post(w, r)
	case http.MethodPut: rh.Put(w, r)
	case http.MethodDelete: rh.Delete(w, r)
	default:
		err := entity.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (rh RoleHandler) Get(w http.ResponseWriter, r *http.Request) {
}

func (rh RoleHandler) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST roles")
	var roleEntity *entity.Role

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &roleEntity)
	if err := BodyValidator(w, roleEntity); err != nil {
		return
	}

	if serviceEntity, _ := rh.Service.GetService(roleEntity.ServiceId); serviceEntity == nil {
		err = entity.BadRequest()
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	roleEntity, err = rh.RoleService.InsertRole(roleEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleEntity)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}

func (rh RoleHandler) Put(w http.ResponseWriter, r *http.Request) {
}

func (rh RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
