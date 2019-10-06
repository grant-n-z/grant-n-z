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

var sghInstance ServiceGroup

type ServiceGroup interface {
	Api(w http.ResponseWriter, r *http.Request)

	post(w http.ResponseWriter, r *http.Request, body []byte)
}

type ServiceGroupImpl struct {
	Request             api.Request
	ServiceGroupService service.ServiceGroupService
}

func GetServiceGroupInstance() ServiceGroup {
	if sghInstance == nil {
		sghInstance = NewServiceGroup()
	}
	return sghInstance
}

func NewServiceGroup() ServiceGroup {
	log.Logger.Info("New `ServiceGroup` instance")
	log.Logger.Info("Inject `Request`, `ServiceGroupService` to `ServiceGroup`")
	return ServiceGroupImpl{
		Request:             api.GetRequestInstance(),
		ServiceGroupService: service.GetServiceGroupServiceInstance(),
	}
}

func (sgh ServiceGroupImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _, err := sgh.Request.Intercept(w, r, property.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodPost:
		sgh.post(w, r, body)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (sgh ServiceGroupImpl) post(w http.ResponseWriter, r *http.Request, body []byte) {
	var serviceGroupEntity *entity.ServiceGroup

	json.Unmarshal(body, &serviceGroupEntity)
	if err := sgh.Request.ValidateBody(w, serviceGroupEntity); err != nil {
		return
	}

	serviceGroup, err := sgh.ServiceGroupService.InsertServiceGroup(serviceGroupEntity)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(serviceGroup)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
