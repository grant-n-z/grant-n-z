package groups

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"testing"

	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	permission Permission
	statusCode int
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	permission = PermissionImpl{PermissionService: StubPermissionService{}}
}

// Test constructor
func TestGetPermissionInstance(t *testing.T) {
	GetPermissionInstance()
}

// Test get bad request
func TestPermission_Get_BadRequest(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, URL: &url.URL{}, Method: http.MethodGet}
	permission.Get(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestPermission_Get_BadRequest test.")
		t.FailNow()
	}
}

// Test post bad request
func TestPermission_Post_BadRequest(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"invalid\":\"test\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: invalid}
	permission.Post(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestPermission_Post_BadRequest test.")
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
// PermissionService
type StubPermissionService struct {
}

func (ps StubPermissionService) GetPermissions() ([]*entity.Permission, *model.ErrorResBody) {
	return []*entity.Permission{}, nil
}

func (ps StubPermissionService) GetPermissionById(id int) (*entity.Permission, *model.ErrorResBody) {
	return &entity.Permission{}, nil
}

func (ps StubPermissionService) GetPermissionByName(name string) (*entity.Permission, *model.ErrorResBody) {
	return &entity.Permission{}, nil
}

func (ps StubPermissionService) GetPermissionsByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody) {
	return []*entity.Permission{}, nil
}

func (ps StubPermissionService) InsertPermission(permission *entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	return &entity.Permission{}, nil
}

func (ps StubPermissionService) InsertWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	return &entity.Permission{}, nil
}
