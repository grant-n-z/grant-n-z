package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var smrhInstance ServiceMemberRoleHandler

type ServiceMemberRoleHandlerImpl struct {
	RequestHandler           RequestHandler
	ServiceMemberRoleService service.ServiceMemberRoleService
}

func GetServiceMemberRoleHandlerInstance() ServiceMemberRoleHandler {
	if smrhInstance == nil {
		smrhInstance = NewServiceMemberRoleHandler()
	}
	return smrhInstance
}

func NewServiceMemberRoleHandler() ServiceMemberRoleHandler {
	log.Logger.Info("New `ServiceMemberRoleHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `ServiceMemberRoleService` to `ServiceMemberRoleHandler`")
	return ServiceMemberRoleHandlerImpl{
		RequestHandler:           GetRequestHandlerInstance(),
		ServiceMemberRoleService: service.NewServiceMemberRoleService(),
	}
}

func (smrhi ServiceMemberRoleHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		smrhi.Get(w, r)
	case http.MethodPost:
		smrhi.Post(w, r)
	case http.MethodPut:
		smrhi.Put(w, r)
	case http.MethodDelete:
		smrhi.Delete(w, r)
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

	body, err := smrhi.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &serviceMemberRoleEntity)
	if err := smrhi.RequestHandler.ValidateHttpRequest(w, serviceMemberRoleEntity); err != nil {
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
