package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var ahInstance Auth

type Auth interface {
	// Api
	// Implement auth api
	// Endpoint is `/api/v1/auth`
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)
}

// AuthImpl
// Auth api struct
type AuthImpl struct {
	tokenProcessor middleware.TokenProcessor
}

// GetAuthInstance
// Get Policy instance.
// If use singleton pattern, call this instance method
func GetAuthInstance() Auth {
	if ahInstance == nil {
		ahInstance = NewAuth()
	}
	return ahInstance
}

// NewAuth
// Constructor
func NewAuth() Auth {
	log.Logger.Info("New `v1.Auth` instance")
	return AuthImpl{tokenProcessor: middleware.GetTokenProcessorInstance()}
}

func (ah AuthImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ah.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (ah AuthImpl) get(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get(middleware.Authorization)
	groupUuids := r.URL.Query().Get("group_uuid")
	roleNames := r.URL.Query().Get("role")
	permissionNames := r.URL.Query().Get("permission")

	_, err := ah.tokenProcessor.VerifyUserToken(token, roleNames, permissionNames, groupUuids)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
