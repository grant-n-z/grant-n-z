package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type PermissionHandlerImpl struct {
	PermissionService service.PermissionService
}

func NewPermissionHandler() PermissionHandler {
	log.Logger.Info("Inject `PermissionService` to `PermissionHandler`")
	return PermissionHandlerImpl{PermissionService: service.NewPermissionService()}
}

func (ph PermissionHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: ph.Get(w, r)
	case http.MethodPost: ph.Post(w, r)
	case http.MethodPut: ph.Put(w, r)
	case http.MethodDelete: ph.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (ph PermissionHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET permissions list")

	permissionEntities, err := ph.PermissionService.GetPermissions()
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(permissionEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (ph PermissionHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST permissions")
	var permissionEntity *entity.Permission

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &permissionEntity)
	if err := BodyValidator(w, permissionEntity); err != nil {
		return
	}

	permissionEntity, err = ph.PermissionService.InsertPermission(permissionEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(permissionEntity)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (ph PermissionHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (ph PermissionHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
