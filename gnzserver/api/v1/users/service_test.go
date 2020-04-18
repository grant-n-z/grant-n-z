package users

import (
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"net/http"
	"testing"
)

var (
	ser Service
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	ser = ServiceImpl{
		Service: StubService{},
	}
}

// Test constructor
func TestGetServiceInstance(t *testing.T) {
	GetServiceInstance()
}

// Test method not allowed
func TestService_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodPut}
	ser.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestService_MethodNotAllowed test.")
		t.FailNow()
	}
}

// Test get
func TestService_Get(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	ser.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestService_Get test.")
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

func (ss StubService) GetServiceById(id int) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceByName(name string) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceOfApiKey() (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceOfUser() ([]*entity.Service, *model.ErrorResBody) {
	return []*entity.Service{}, nil
}

func (ss StubService) InsertService(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &service, nil
}

func (ss StubService) InsertServiceWithRelationalData(service *entity.Service) (*entity.Service, *model.ErrorResBody) {
	return service, nil
}

func (ss StubService) GenerateApiKey() string {
	return ""
}
