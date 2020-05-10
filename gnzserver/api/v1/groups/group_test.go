package groups

import (
	"testing"

	"net/http"
	"net/url"

	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	group Group
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	group = GroupImpl{GroupService: StubGroupService{}}
}

// Test constructor
func TestGetGroupInstance(t *testing.T) {
	GetGroupInstance()
}

// Test get bad request
func TestGroup_Get_BadRequest(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, URL: &url.URL{}, Method: http.MethodGet}
	group.Get(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestGroup_Get_BadRequest test.")
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
	return &entity.Group{}, nil
}
