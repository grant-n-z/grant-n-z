package operator

import (
	"bytes"
	"testing"

	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	operatorPolicy OperatorPolicy
	statusCode     int
)

func init() {
	log.InitLogger("info")

	operatorPolicy = OperatorPolicyImpl{OperatorPolicyService: StubOperatorPolicyService{}}
}

// Test constructor
func TestGetOperatorPolicyInstance(t *testing.T) {
	GetOperatorPolicyInstance()
}

// Test method not allowed
func TestOperatorPolicy_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodPatch}
	operatorPolicy.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestOperatorPolicy_MethodNotAllowed test.")
		t.FailNow()
	}
}

// Test get
func TestOperatorPolicy_Get(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet, URL: &url.URL{}}
	operatorPolicy.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestOperatorPolicy_Get test.")
		t.FailNow()
	}
}

// Test post bad request
func TestOperatorPolicy_Post_BadRequest(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"role_uuid\":\"fdf515e7-b163-435e-bc69-0fc0f90bd90b\", \"user_uuid\":\"invalid\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: invalid}
	operatorPolicy.Api(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestOperatorPolicy_Post_BadRequest test.")
		t.FailNow()
	}
}

// Test post
func TestOperatorPolicy_Post(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"role_uuid\":\"fdf515e7-b163-435e-bc69-0fc0f90bd90b\", \"user_uuid\":\"123515e7-b163-435e-bc69-0fc0f90bd90b\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: invalid}
	operatorPolicy.Api(response, &request)

	if statusCode != http.StatusCreated {
		t.Errorf("Incorrect TestOperatorPolicy_Post test.")
		t.FailNow()
	}
}

// Test put
func TestOperatorPolicy_Put(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"role_uuid\":\"fdf515e7-b163-435e-bc69-0fc0f90bd90b\", \"user_uuid\":\"123515e7-b163-435e-bc69-0fc0f90bd90b\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPut, Body: invalid}
	operatorPolicy.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestOperatorPolicy_Put test.")
		t.FailNow()
	}
}

// Test delete
func TestOperatorPolicy_Delete(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodDelete}
	operatorPolicy.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestOperatorPolicy_Delete test.")
		t.FailNow()
	}
}

// Less than stub struct
// ResponseWriter
type StubResponseWriter struct {
}

func (w StubResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w StubResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w StubResponseWriter) WriteHeader(code int) {
	statusCode = code
}

// Less than stub struct
// OperatorPolicyService
type StubOperatorPolicyService struct {
}

func (ops StubOperatorPolicyService) Get(queryParam string) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	return []*entity.OperatorPolicy{}, nil
}

func (ops StubOperatorPolicyService) GetAll() ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	return []*entity.OperatorPolicy{}, nil
}

func (ops StubOperatorPolicyService) GetByUserUuid(userUuid string) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	return []*entity.OperatorPolicy{}, nil
}

func (ops StubOperatorPolicyService) GetByUserUuidAndRoleUuid(userUuid string, roleUuid string) (*entity.OperatorPolicy, *model.ErrorResBody) {
	return &entity.OperatorPolicy{}, nil
}

func (ops StubOperatorPolicyService) Insert(policy *entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody) {
	return &entity.OperatorPolicy{}, nil
}
