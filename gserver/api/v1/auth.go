package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var ahInstance Auth

type Auth interface {
	Api(w http.ResponseWriter, r *http.Request)

	get(w http.ResponseWriter, r *http.Request)
}

type AuthImpl struct {
	Request api.Request
}

func GetAuthInstance() Auth {
	if ahInstance == nil {
		ahInstance = NewAuth()
	}
	return ahInstance
}

func NewAuth() Auth {
	log.Logger.Info("New `Auth` instance")
	log.Logger.Info("Inject `Request`to `Auth`")
	return AuthImpl{
		Request: api.GetRequestInstance(),
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
	_, err := ah.Request.Intercept(w, r)
	if err != nil {
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
