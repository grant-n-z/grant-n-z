package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/service"
	"github.com/tomoyane/grant-n-z/server/log"
)

type ServiceHandler struct {
	Service service.Service
}

func NewServiceHandler() ServiceHandler {
	return ServiceHandler{Service: service.NewServiceService()}
}

func (sh ServiceHandler) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST services")
	var serviceEntity *entity.Service

	body, err := Interceptor(w, r, http.MethodPost)
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
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}
