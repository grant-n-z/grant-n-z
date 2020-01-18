package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/middleware"
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

// Auth api struct
type AuthImpl struct {
	tokenService service.TokenService
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
	return AuthImpl{tokenService: service.GetTokenServiceInstance()}
}

func (ah AuthImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		ah.get(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (ah AuthImpl) get(w http.ResponseWriter, r *http.Request) {
	var result bool
	token := r.Header.Get(middleware.Authorization)
	roleName := r.URL.Query().Get("role")
	permissionName := r.URL.Query().Get("permission")

	_, err := ah.tokenService.VerifyUserToken(token, &roleName, &permissionName)
	if err != nil {
		result = false
	} else {
		result = true
	}

	res, _ := json.Marshal(map[string]bool{"grant": result})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
