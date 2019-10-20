package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ahInstance Auth

type Auth interface {
	// Implement auth api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)
}

type AuthImpl struct {
	request      api.Request
	tokenService service.TokenService
}

func GetAuthInstance() Auth {
	if ahInstance == nil {
		ahInstance = NewAuth()
	}
	return ahInstance
}

func NewAuth() Auth {
	log.Logger.Info("New `Auth` instance")
	log.Logger.Info("Inject `Request`, `TokenService` to `Auth`")
	return AuthImpl{
		request:      api.GetRequestInstance(),
		tokenService: service.GetTokenServiceInstance(),
	}
}

func (ah AuthImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		ah.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (ah AuthImpl) get(w http.ResponseWriter, r *http.Request) {
	var result bool
	_, _, err := ah.request.Intercept(w, r, property.AuthUser)
	if err != nil {
		result = false
	} else {
		result = true
	}

	res, _ := json.Marshal(map[string]bool{"grant": result})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
