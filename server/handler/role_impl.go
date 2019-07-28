package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type RoleHandlerImpl struct {
	RoleService service.RoleService
}

func NewRoleHandler() RoleHandler {
	log.Logger.Info("Inject `RoleService` to `RoleHandler`")
	return RoleHandlerImpl{RoleService: service.NewRoleService()}
}

func (rh RoleHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: rh.Get(w, r)
	case http.MethodPost: rh.Post(w, r)
	case http.MethodPut: rh.Put(w, r)
	case http.MethodDelete: rh.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (rh RoleHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET roles list")

	roleEntities, err := rh.RoleService.GetRoles()
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (rh RoleHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST roles")
	var roleEntity *entity.Role

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &roleEntity)
	if err := ValidateHttpRequest(w, roleEntity); err != nil {
		return
	}

	roleEntity, err = rh.RoleService.InsertRole(roleEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleEntity)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (rh RoleHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (rh RoleHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}