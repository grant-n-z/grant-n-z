package users

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var shInstance Service

type Service interface {
	// Implement service api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)
}

// Service api struct
type ServiceImpl struct {
	Service service.Service
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetServiceInstance() Service {
	if shInstance == nil {
		shInstance = NewService()
	}
	return shInstance
}

// Constructor
func NewService() Service {
	log.Logger.Info("New `Service` instance")
	log.Logger.Info("Inject `Service` to `Service`")
	return ServiceImpl{Service: service.GetServiceInstance()}
}

func (sh ServiceImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		sh.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (sh ServiceImpl) get(w http.ResponseWriter, r *http.Request) {
	result, err := sh.Service.GetServiceOfUser()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
