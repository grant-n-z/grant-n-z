package handler

import (
	"encoding/json"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/service"
	"github.com/tomoyane/grant-n-z/server/log"
	"net/http"
)

type ServiceHandler struct {
	Service service.Service
}

func NewServiceHandler() ServiceHandler {
	log.Logger.Info("inject `Service` to `ServiceHandler`")
	return ServiceHandler{Service: service.NewServiceService()}
}

func (sh ServiceHandler) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet: sh.Get(w, r)
	case http.MethodPost: sh.Post(w, r)
	case http.MethodPut: sh.Put(w, r)
	case http.MethodDelete: sh.Delete(w, r)
	default:
		err := entity.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (sh ServiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET services")
	name := r.URL.Query().Get(entity.SERVICE_NAME.String())

	result, err := sh.Service.Get(name)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (sh ServiceHandler) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST services")
	var serviceEntity *entity.Service

	body, err := Interceptor(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &serviceEntity)
	if err := BodyValidator(w, serviceEntity); err != nil {
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

func (sh ServiceHandler) Put(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("PUT services")
}

func (sh ServiceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("DELETE services")
}
