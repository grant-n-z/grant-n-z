package operator

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

var operatorShInstance OperatorService

type OperatorService interface {
	// Implement service admin service api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	post(w http.ResponseWriter, r *http.Request, body []byte)

	// Http PUT method
	put(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
	delete(w http.ResponseWriter, r *http.Request)
}

type OperatorServiceImpl struct {
	Request api.Request
	Service service.Service
}

func GetOperatorServiceInstance() OperatorService {
	if operatorShInstance == nil {
		operatorShInstance = NewOperatorService()
	}
	return operatorShInstance
}

func NewOperatorService() OperatorService {
	log.Logger.Info("New `OperatorService` instance")
	log.Logger.Info("Inject `request`, `Service` to `OperatorService`")
	return OperatorServiceImpl{
		Request: api.GetRequestInstance(),
		Service: service.GetServiceInstance(),
	}
}

func (sh OperatorServiceImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := sh.Request.Intercept(w, r, property.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		sh.get(w, r)
	case http.MethodPost:
		sh.post(w, r, body)
	case http.MethodPut:
		sh.put(w, r)
	case http.MethodDelete:
		sh.delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (sh OperatorServiceImpl) get(w http.ResponseWriter, r *http.Request) {
	result, err := sh.Service.GetServices()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (sh OperatorServiceImpl) post(w http.ResponseWriter, r *http.Request, body []byte) {
	var serviceEntity *entity.Service

	json.Unmarshal(body, &serviceEntity)
	if err := sh.Request.ValidateBody(w, serviceEntity); err != nil {
		return
	}

	serviceData, err := sh.Service.InsertService(serviceEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(serviceData)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (sh OperatorServiceImpl) put(w http.ResponseWriter, r *http.Request) {
}

func (sh OperatorServiceImpl) delete(w http.ResponseWriter, r *http.Request) {
}
