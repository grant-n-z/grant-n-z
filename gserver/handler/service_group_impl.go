package handler

import (
	"encoding/json"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var sghInstance ServiceGroupHandler

type ServiceGroupHandlerImpl struct {
	RequestHandler      RequestHandler
	ServiceGroupService service.ServiceGroupService
	AuthService         service.AuthService
}

func GetServiceGroupHandlerInstance() ServiceGroupHandler {
	if sghInstance == nil {
		sghInstance = NewServiceGroupHandler()
	}
	return sghInstance
}

func NewServiceGroupHandler() ServiceGroupHandler {
	log.Logger.Info("New `ServiceGroupHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `AuthService`, `ServiceGroupService` to `ServiceGroupHandler`")
	return ServiceGroupHandlerImpl{
		RequestHandler:      GetRequestHandlerInstance(),
		ServiceGroupService: service.GetServiceGroupServiceInstance(),
		AuthService:         service.GetAuthServiceInstance(),
	}
}

func (sgh ServiceGroupHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		sgh.Post(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (sgh ServiceGroupHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	var serviceGroupEntity *entity.ServiceGroup

	body, err := sgh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &serviceGroupEntity)
	if err := sgh.RequestHandler.ValidateHttpRequest(w, serviceGroupEntity); err != nil {
		return
	}

	serviceGroupEntity, err = sgh.ServiceGroupService.InsertServiceGroup(serviceGroupEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(serviceGroupEntity)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
