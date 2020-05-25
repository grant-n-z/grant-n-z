package operator

import (
	"bytes"
	"net/http"
	"testing"

	"io/ioutil"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	operatorService OperatorService
)

func init() {
	log.InitLogger("info")

	operatorService = OperatorServiceImpl{Service: StubService{}}
}

// Test constructor
func TestGetOperatorServiceInstance(t *testing.T) {
	GetOperatorServiceInstance()
}

// Test method not allowed
func TestOperatorService_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodPatch}
	operatorService.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestOperatorService_MethodNotAllowed test.")
		t.FailNow()
	}
}

// Test get
func TestOperatorService_Get(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	operatorService.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestOperatorService_Get test.")
		t.FailNow()
	}
}

// Test post bad request
func TestOperatorService_Post_BadRequest(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"invalid\":\"test\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: invalid}
	operatorService.Api(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestOperatorService_Post_BadRequest test.")
		t.FailNow()
	}
}

// Test post
func TestOperatorService_Post(t *testing.T) {
	response := StubResponseWriter{}
	body := ioutil.NopCloser(bytes.NewReader([]byte("{\"name\":\"test\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: body}
	operatorService.Api(response, &request)

	if statusCode != http.StatusCreated {
		t.Errorf("Incorrect TestOperatorService_Post test.")
		t.FailNow()
	}
}

// Test put
func TestOperatorService_Put(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"name\":\"test\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPut, Body: invalid}
	operatorService.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestOperatorService_Put test.")
		t.FailNow()
	}
}

// Test delete
func TestOperatorService_Delete(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodDelete}
	operatorService.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestOperatorService_Delete test.")
		t.FailNow()
	}
}

// Less than stub struct
// Service
type StubService struct {
}

func (ss StubService) GetServices() ([]*entity.Service, *model.ErrorResBody) {
	return []*entity.Service{}, nil
}

func (ss StubService) GetServiceByUuid(uuid string) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceByName(name string) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceBySecret(secret string) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceByUser(userUuid string) ([]*entity.Service, *model.ErrorResBody) {
	return []*entity.Service{}, nil
}

func (ss StubService) InsertService(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) InsertServiceWithRelationalData(service *entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GenerateSecret() string {
	return ""
}
