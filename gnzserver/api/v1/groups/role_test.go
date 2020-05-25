package groups

import (
	"bytes"
	"io/ioutil"
	"testing"

	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	role Role
)

func init() {
	log.InitLogger("info")

	role = RoleImpl{RoleService: StubRoleService{}}
}

// Test constructor
func TestGetRoleInstance(t *testing.T) {
	GetRoleInstance()
}

// Test get
func TestRole_Get_Success(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	role.Get(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestGet_BadRequest test. %d", statusCode)
		t.FailNow()
	}
}

// Test post bad request
func TestRole_Post_BadRequest_Body(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"invalid\":\"test\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: invalid}
	role.Post(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestRole_Post_BadRequest_Body test.")
		t.FailNow()
	}
}

// Test post bad request
func TestRole_Post_BadRequest_Success(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"name\":\"test\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: invalid}
	role.Post(response, &request)

	if statusCode != http.StatusCreated {
		t.Errorf("Incorrect TestRole_Post_BadRequest_QueryParam test. %d", statusCode)
		t.FailNow()
	}
}

// Less than stub struct
// RoleService
type StubRoleService struct {
}

func (rs StubRoleService) GetRoles() ([]*entity.Role, *model.ErrorResBody) {
	return []*entity.Role{}, nil
}

func (rs StubRoleService) GetRoleByUuid(uuid string) (*entity.Role, *model.ErrorResBody) {
	return &entity.Role{}, nil
}

func (rs StubRoleService) GetRoleByName(name string) (*entity.Role, *model.ErrorResBody) {
	return &entity.Role{}, nil
}

func (rs StubRoleService) GetRoleByNames(names []string) ([]entity.Role, *model.ErrorResBody) {
	return []entity.Role{}, nil
}

func (rs StubRoleService) GetRolesByGroupUuid(groupUuid string) ([]*entity.Role, *model.ErrorResBody) {
	return []*entity.Role{}, nil
}

func (rs StubRoleService) InsertRole(role *entity.Role) (*entity.Role, *model.ErrorResBody) {
	return &entity.Role{}, nil
}

func (rs StubRoleService) InsertWithRelationalData(groupUuid string, role entity.Role) (*entity.Role, *model.ErrorResBody) {
	return &entity.Role{}, nil
}
