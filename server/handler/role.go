package handler

import (
	"encoding/json"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/service"
	"github.com/tomoyane/grant-n-z/server/log"
	"net/http"
)

type RoleHandler struct {
	RoleService service.RoleService
}

func NewRoleHandler() RoleHandler {
	log.Logger.Info("inject `Service`, `RoleService` to `RoleHandler`")
	return RoleHandler{RoleService: service.NewRoleService()}
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
	log.Logger.Info("GET roles list")
	var result interface{}

	roleEntities, err := rh.RoleService.GetRoles()
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	if roleEntities == nil {
		result = []string{}
	} else {
		result = roleEntities
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
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
