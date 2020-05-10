package service

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	groupService GroupService
	stubConnection *gorm.DB
)

// Set up
func init() {
	log.InitLogger("info")
	ctx.InitContext()
	ctx.SetUserId(1)
	ctx.SetServiceId(1)

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")

	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})
	stubEtcdClient := cache.EtcdClientImpl{
		Connection: stubEtcdConnection,
		Ctx:        ctx.GetCtx(),
	}

	groupService = GroupServiceImpl{
		EtcdClient:           stubEtcdClient,
		GroupRepository:      StubGroupRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}
}

// Test get groups
func TestGetGroupServiceInstance(t *testing.T) {
	GetGroupServiceInstance()
}

// Test get groups
func TestGetGroups_Success(t *testing.T) {
	_, err := groupService.GetGroups()
	if err != nil {
		t.Errorf("Incorrect TestGetGroups_Success test")
		t.FailNow()
	}
}

// Test get group by id
func TestGetGroupById_Success(t *testing.T) {
	_, err := groupService.GetGroupById(1)
	if err != nil {
		t.Errorf("Incorrect TestGetGroupById_Success test")
		t.FailNow()
	}
}

// Test get group of login user
func TestGetGroupOfUser_Success(t *testing.T) {
	_, err := groupService.GetGroupOfUser()
	if err != nil {
		t.Errorf("Incorrect TestGetGroupOfUser_Success test")
		t.FailNow()
	}
}

// Test insert group with relational data
func TestInsertGroupWithRelationalData_Success(t *testing.T) {
	_, err := groupService.InsertGroupWithRelationalData(entity.Group{Id:1, Name:"test", Uuid: uuid.New()})
	if err != nil {
		t.Errorf("Incorrect TestInsertGroupWithRelationalData_Success test")
		t.FailNow()
	}
}

// Less than stub struct
// Group repository
type StubGroupRepositoryImpl struct {
	Connection *gorm.DB
}

func (gr StubGroupRepositoryImpl) FindAll() ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group
	return groups, nil
}

func (gr StubGroupRepositoryImpl) FindById(id int) (*entity.Group, *model.ErrorResBody) {
	var group entity.Group
	return &group, nil
}

func (gr StubGroupRepositoryImpl) FindByName(name string) (*entity.Group, *model.ErrorResBody) {
	var group *entity.Group
	return group, nil
}

func (gr StubGroupRepositoryImpl) FindByUserId(userId int) ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group
	return groups, nil
}

func (gr StubGroupRepositoryImpl) FindGroupWithUserWithPolicyGroupsByUserId(userId int) ([]*model.GroupWithUserGroupWithPolicy, *model.ErrorResBody) {
	var groupWithUserGroupWithPolicies []*model.GroupWithUserGroupWithPolicy
	groupWithUserGroupWithPolicies = append(groupWithUserGroupWithPolicies, &model.GroupWithUserGroupWithPolicy{entity.Group{}, entity.UserGroup{}, entity.Policy{ServiceId: 1}})
	return groupWithUserGroupWithPolicies, nil
}

func (gr StubGroupRepositoryImpl) FindGroupWithPolicyByUserIdAndGroupId(userId int, groupId int) (*model.GroupWithUserGroupWithPolicy, *model.ErrorResBody) {
	var groupWithUserGroupWithPolicy model.GroupWithUserGroupWithPolicy
	return &groupWithUserGroupWithPolicy, nil

}

func (gr StubGroupRepositoryImpl) SaveWithRelationalData(
	group entity.Group,
	serviceGroup entity.ServiceGroup,
	userGroup entity.UserGroup,
	groupPermission entity.GroupPermission,
	groupRole entity.GroupRole,
	policy entity.Policy) (*entity.Group, *model.ErrorResBody) {

	return &group, nil
}
