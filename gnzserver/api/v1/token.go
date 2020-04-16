package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var thInstance Token

type Token interface {
	// Implement token api
	Api(w http.ResponseWriter, r *http.Request)

	// Http POST method
	post(w http.ResponseWriter, r *http.Request)
}

// Token api struct
type TokenImpl struct {
	TokenProcessor middleware.TokenProcessor
}

// Get Policy instance
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
	return TokenImpl{TokenProcessor: middleware.GetTokenProcessorInstance()}
}

func (th TokenImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		th.post(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (th TokenImpl) post(w http.ResponseWriter, r *http.Request) {
	var userEntity *entity.User
	if err := middleware.BindBody(w, r, &userEntity); err != nil {
		return
	}

	userEntity.Username = userEntity.Email
	if err := middleware.ValidateBody(w, userEntity); err != nil {
		return
	}

	userType := r.URL.Query().Get("type")
	groupId := r.URL.Query().Get("group_id")
	token, err := th.TokenProcessor.Generate(userType, groupId, *userEntity)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(map[string]string{"token": token})
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
