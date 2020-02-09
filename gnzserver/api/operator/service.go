package operator

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/config"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/api"
	"github.com/tomoyane/grant-n-z/gnzserver/entity"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var shInstance Service

type Service interface {
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

func GetOperatorServiceInstance() Service {
	if shInstance == nil {
		shInstance = NewOperatorService()
	}
	return shInstance
}

func NewOperatorService() Service {
	log.Logger.Info("New `OperatorService` instance")
	return OperatorServiceImpl{
		Request: api.GetRequestInstance(),
		Service: service.GetServiceInstance(),
	}
}

func (sh OperatorServiceImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := sh.Request.Intercept(w, r, config.AuthOperator)
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

	if err := middleware.BindBody(w, r, &serviceEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, serviceEntity); err != nil {
		return
	}

	serviceData, err := sh.Service.InsertServiceWithRelationalData(serviceEntity)
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
