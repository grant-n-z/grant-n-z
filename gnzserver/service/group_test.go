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
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
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
		etcdClient: stubEtcdClient,
		groupRepository: driver.GroupRepositoryImpl{Connection: stubConnection},
		roleRepository: StubRoleRepositoryImpl{Connection: stubConnection},
		permissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}
}

// Test get groups
func TestGetGroupServiceInstance(t *testing.T) {
	GetGroupServiceInstance()
}

// Test get groups
func TestGetGroups_GetError(t *testing.T) {
	_, err := groupService.GetGroups()
	if err == nil {
		t.Errorf("Incorrect TestGetGroups_GetError test")
		t.FailNow()
	}
}

// Test get group by id
func TestGetGroupById_GetError(t *testing.T) {
	_, err := groupService.GetGroupById(1)
	if err == nil {
		t.Errorf("Incorrect TestGetGroupById_GetError test")
		t.FailNow()
	}
}

// Test get group of login user
func TestGetGroupOfUser_GetError(t *testing.T) {
	_, err := groupService.GetGroupOfUser()
	if err == nil {
		t.Errorf("Incorrect TestGetGroupOfUser_GetError test")
		t.FailNow()
	}
}

// Test insert group with relational data
func TestInsertGroupWithRelationalData_GetError(t *testing.T) {
	_, err := groupService.InsertGroupWithRelationalData(entity.Group{Id:1, Name:"test", Uuid: uuid.New()})
	if err == nil {
		t.Errorf("Incorrect TestInsertGroupWithRelationalData_GetError test")
		t.FailNow()
	}
}
