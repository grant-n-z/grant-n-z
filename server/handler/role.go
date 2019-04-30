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
}

func NewRoleHandler() RoleHandler {
	return RoleHandler{RoleService: service.NewRoleService()}
}

func (rh RoleHandler) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST roles")
	var roleEntity *entity.Role

	body, err := Interceptor(w, r, http.MethodPost)
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
