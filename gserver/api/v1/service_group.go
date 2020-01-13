package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/common/constant"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var sghInstance ServiceGroup

type ServiceGroup interface {
	// Implement service group api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)
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
	log.Logger.Info("Inject `request`, `ServiceGroupService` to `ServiceGroup`")
	return ServiceGroupImpl{
		Request:             api.GetRequestInstance(),
		ServiceGroupService: service.GetServiceGroupServiceInstance(),
	}
}

func (sgh ServiceGroupImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := sgh.Request.Intercept(w, r, constant.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		sgh.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (sgh ServiceGroupImpl) get(w http.ResponseWriter, r *http.Request) {
	var serviceGroup *entity.ServiceGroup
	res, _ := json.Marshal(serviceGroup)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
