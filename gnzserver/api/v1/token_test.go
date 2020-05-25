package v1

import (
	"bytes"
	"testing"

	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	token Token
)

func init() {
	log.InitLogger("info")

	token = TokenImpl{
		TokenProcessor: StubTokenProcessor{},
	}
}

// Test constructor
func TestGetTokenInstance(t *testing.T) {
	GetTokenInstance()
}

// Test token method not allowed
func TestToken_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	token.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestToken_MethodNotAllowed test.")
		t.FailNow()
	}
}

// Test token bad request
func TestToken_Post_BadRequest(t *testing.T) {
	response := StubResponseWriter{}
	noneEmailBody := ioutil.NopCloser(bytes.NewReader([]byte("{\"username\":\"test\", \"password\":\"testtest\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: noneEmailBody}
	token.Api(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestToken_Post_BadRequest test.")
		t.FailNow()
	}
}

// Test token
func TestToken_Post(t *testing.T) {
	response := StubResponseWriter{}
	body := ioutil.NopCloser(bytes.NewReader([]byte("{\"email\":\"test@gmail.com\", \"password\":\"testtest\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: body, URL: &url.URL{}}
	token.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestToken_Post test.")
		t.FailNow()
	}
}

// Less than stub struct
// TokenProcessor
type StubTokenProcessor struct {
}

func (tp StubTokenProcessor) Generate(userType string, tokenRequest model.TokenRequest) (*model.TokenResponse, *model.ErrorResBody) {
	return &model.TokenResponse{}, nil
}

func (tp StubTokenProcessor) VerifyOperatorToken(token string) (*model.JwtPayload, *model.ErrorResBody) {
	return &model.JwtPayload{}, nil
}

func (tp StubTokenProcessor) VerifyUserToken(token string, roleNames []string, permissionNames []string, groupUuid string) (*model.JwtPayload, *model.ErrorResBody) {
	return &model.JwtPayload{}, nil
}

func (tp StubTokenProcessor) GetJwtPayload(token string, isRefresh bool) (*model.JwtPayload, *model.ErrorResBody) {
	return &model.JwtPayload{}, nil
}
