package groups

import (
	"testing"

	"net/http"
	"net/url"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	group Group
)

func init() {
	log.InitLogger("info")

	group = GroupImpl{GroupService: StubGroupService{}}
}

// Test constructor
func TestGetGroupInstance(t *testing.T) {
	GetGroupInstance()
}

// Test get
func TestGroup_Get_Success(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, URL: &url.URL{}, Method: http.MethodGet}
	request.URL.Host = "localhost:8080"
	request.URL.Path = "/api/v1/groups/8532541a-9aa5-4a44-8f87-b630d2a3d01f"
	group.Get(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestGroup_Get_InternalServer test. %d", statusCode)
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

func (gs StubGroupService) GetGroupByUuid(uuid string) (*entity.Group, *model.ErrorResBody) {
	return &entity.Group{}, nil
}

func (gs StubGroupService) GetGroupByUser(userUuid string) ([]*entity.Group, *model.ErrorResBody) {
	return []*entity.Group{}, nil
}

func (gs StubGroupService) InsertGroupWithRelationalData(group entity.Group, userUuid string, serviceUuid string) (*entity.Group, *model.ErrorResBody) {
	return &entity.Group{}, nil
}
