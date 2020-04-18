package users

import (
	"bytes"
	"testing"

	"io/ioutil"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	group      Group
	statusCode int
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	group = GroupImpl{
		groupService: StubGroupService{},
	}
}

// Test constructor
func TestGetGroupInstance(t *testing.T) {
	GetGroupInstance()
}

// Test method not allowed
func TestGroup_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodPut}
	group.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestGroup_MethodNotAllowed test.")
		t.FailNow()
	}
}

// Test get
func TestGroup_Get(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	group.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestGroup_Get test.")
		t.FailNow()
	}
}

// Test post bad request
func TestGroup_Post_BadRequest(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"invalid\":\"test\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: invalid}
	group.Api(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestGroup_Post_BadRequest test.")
		t.FailNow()
	}
}

// Test post
func TestGroup_Post(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"name\":\"test\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPost, Body: invalid}
	group.Api(response, &request)

	if statusCode != http.StatusCreated {
		t.Errorf("Incorrect TestGroup_Post test.")
		t.FailNow()
	}
}

// Less than stub struct
// GroupService
type StubGroupService struct {
}

func (gs StubGroupService) GetGroups() ([]*entity.Group, *model.ErrorResBody) {
	return []*entity.Group{}, nil
}

func (gs StubGroupService) GetGroupById(id int) (*entity.Group, *model.ErrorResBody) {
	return &entity.Group{}, nil
}

func (gs StubGroupService) GetGroupOfUser() ([]*entity.Group, *model.ErrorResBody) {
	return []*entity.Group{}, nil
}

func (gs StubGroupService) InsertGroupWithRelationalData(group entity.Group) (*entity.Group, *model.ErrorResBody) {
	return &group, nil
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
