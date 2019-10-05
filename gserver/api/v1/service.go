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

var shInstance Service

type Service interface {
	Api(w http.ResponseWriter, r *http.Request)

	get(w http.ResponseWriter, r *http.Request)

	post(w http.ResponseWriter, r *http.Request)

	put(w http.ResponseWriter, r *http.Request)

	delete(w http.ResponseWriter, r *http.Request)
}

type ServiceImpl struct {
	Request api.Request
	Service service.Service
}

func GetServiceInstance() Service {
	if shInstance == nil {
		shInstance = NewService()
	}
	return shInstance
}

func NewService() Service {
	log.Logger.Info("New `Service` instance")
	log.Logger.Info("Inject `Request`, `Service` to `Service`")
	return ServiceImpl{
		Request: api.GetRequestInstance(),
		Service: service.GetServiceInstance(),
	}
}

func (sh ServiceImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := sh.Request.VerifyToken(w, r, property.AuthOperator)
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		sh.get(w, r)
	case http.MethodPost:
		sh.post(w, r)
	case http.MethodPut:
		sh.put(w, r)
	case http.MethodDelete:
		sh.delete(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (sh ServiceImpl) get(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(entity.ServiceName.String())

	result, err := sh.Service.Get(name)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (sh ServiceImpl) post(w http.ResponseWriter, r *http.Request) {
	var serviceEntity *entity.Service

	body, err := sh.Request.Intercept(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &serviceEntity)
	if err := sh.Request.ValidateBody(w, serviceEntity); err != nil {
		return
	}

	serviceEntity, err = sh.Service.InsertService(serviceEntity)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(serviceEntity)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (sh ServiceImpl) put(w http.ResponseWriter, r *http.Request) {
}

func (sh ServiceImpl) delete(w http.ResponseWriter, r *http.Request) {
}
