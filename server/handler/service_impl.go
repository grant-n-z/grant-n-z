package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/entity"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/model"

	"github.com/tomoyane/grant-n-z/server/usecase/service"
)

type ServiceHandlerImpl struct {
	RequestHandler RequestHandler
	Service        service.Service
}

func NewServiceHandler() ServiceHandler {
	log.Logger.Info("Inject `RequestHandler`, `Service` to `ServiceHandler`")
	return ServiceHandlerImpl{
		RequestHandler: NewRequestHandler(),
		Service: service.NewServiceService(),
	}
}

func (sh ServiceHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		sh.Get(w, r)
	case http.MethodPost:
		sh.Post(w, r)
	case http.MethodPut:
		sh.Put(w, r)
	case http.MethodDelete:
		sh.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (sh ServiceHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET services")
	name := r.URL.Query().Get(entity.ServiceName.String())

	result, err := sh.Service.Get(name)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (sh ServiceHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST services")
	var serviceEntity *entity.Service

	body, err := sh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &serviceEntity)
	if err := sh.RequestHandler.ValidateHttpRequest(w, serviceEntity); err != nil {
		return
	}

	serviceEntity, err = sh.Service.InsertService(serviceEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(serviceEntity)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (sh ServiceHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("PUT services")
}

func (sh ServiceHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("DELETE services")
}
