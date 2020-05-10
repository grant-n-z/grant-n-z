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
	// Implement auth api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)
}

// Auth api struct
type AuthImpl struct {
	tokenProcessor middleware.TokenProcessor
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetAuthInstance() Auth {
	if ahInstance == nil {
		ahInstance = NewAuth()
	}
	return ahInstance
}

// Constructor
func NewAuth() Auth {
	log.Logger.Info("New `Auth` instance")
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
	roleName := r.URL.Query().Get("role")
	permissionName := r.URL.Query().Get("permission")

	_, err := ah.tokenProcessor.VerifyUserToken(token, []string{roleName}, permissionName, 0)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]bool{"grant": true})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
