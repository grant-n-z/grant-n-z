package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var thInstance Token

type Token interface {
	// Implement token api
	Api(w http.ResponseWriter, r *http.Request)

	// Http POST method
	post(w http.ResponseWriter, r *http.Request, body []byte)
}

// Token api struct
type TokenImpl struct {
	Request      api.Request
	TokenService service.TokenService
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetTokenInstance() Token {
	if thInstance == nil {
		thInstance = NewToken()
	}
	return thInstance
}

// Constructor
func NewToken() Token {
	log.Logger.Info("New `Token` instance")
	log.Logger.Info("Inject `request`, `TokenService` to `Token`")
	return TokenImpl{
		Request:      api.GetRequestInstance(),
		TokenService: service.GetTokenServiceInstance(),
	}
}

func (th TokenImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := th.Request.Intercept(w, r, "")
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodPost:
		th.post(w, r, body)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (th TokenImpl) post(w http.ResponseWriter, r *http.Request, body []byte) {
	var userEntity *entity.User

	json.Unmarshal(body, &userEntity)
	userEntity.Username = userEntity.Email
	if err := th.Request.ValidateBody(w, userEntity); err != nil {
		return
	}

	token, err := th.TokenService.Generate(r.URL.Query().Get("type"), *userEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"token": *token})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
