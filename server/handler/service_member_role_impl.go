package handler

import (
	"encoding/json"
	"github.com/tomoyane/grant-n-z/server/usecase/service"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"
)

type ServiceMemberRoleHandlerImpl struct {
	ServiceMemberRoleService service.ServiceMemberRoleService
}

func NewServiceMemberRoleHandler() ServiceMemberRoleHandler {
	log.Logger.Info("Inject `ServiceMemberRoleService` to `ServiceMemberRoleHandler`")
	return ServiceMemberRoleHandlerImpl{
		ServiceMemberRoleService: service.NewServiceMemberRoleService(),
	}
}

func (smrhi ServiceMemberRoleHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: smrhi.Get(w, r)
	case http.MethodPost: smrhi.Post(w, r)
	case http.MethodPut: smrhi.Put(w, r)
	case http.MethodDelete: smrhi.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (smrhi ServiceMemberRoleHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET service_member_roles")

	serviceMemberRoleEntities, err := smrhi.ServiceMemberRoleService.GetAll()
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(serviceMemberRoleEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (smrhi ServiceMemberRoleHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST service_member_roles")
	var serviceMemberRoleEntity *entity.ServiceMemberRole

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &serviceMemberRoleEntity)
	if err := ValidateHttpRequest(w, serviceMemberRoleEntity); err != nil {
		return
	}

	serviceMemberRoleEntity, err = smrhi.ServiceMemberRoleService.Insert(serviceMemberRoleEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(serviceMemberRoleEntity)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (smrhi ServiceMemberRoleHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (smrhi ServiceMemberRoleHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
