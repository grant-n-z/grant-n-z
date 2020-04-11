package service

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	permissionService PermissionService
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

	permissionService = PermissionServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}
}

// Test constructor
func TestGetPermissionServiceInstance(t *testing.T) {
	GetPermissionServiceInstance()
}

// Test get permissions
func TestGetPermissions_Success(t *testing.T) {
	_, err := permissionService.GetPermissions()
	if err != nil {
		t.Errorf("Incorrect TestGetPermissions_Success test")
		t.FailNow()
	}
}

// Test get by id
func TestGetPermissionById_Success(t *testing.T) {
	_, err := permissionService.GetPermissionById(1)
	if err != nil {
		t.Errorf("Incorrect TestGetPermissionById_Success test")
		t.FailNow()
	}
}

// Test get by name
func TestGetPermissionByName_Success(t *testing.T) {
	_, err := permissionService.GetPermissionByName("name")
	if err != nil {
		t.Errorf("Incorrect TestGetPermissionByName_Success test")
		t.FailNow()
	}
}

// Test get by group id
func TestGetPermissionsByGroupId_Success(t *testing.T) {
	_, err := permissionService.GetPermissionsByGroupId(1)
	if err != nil {
		t.Errorf("Incorrect TestGetPermissionsByGroupId_Success test")
		t.FailNow()
	}
}

// Test insert
func TestInsertPermission_Success(t *testing.T) {
	_, err := permissionService.InsertPermission(&entity.Permission{})
	if err != nil {
		t.Errorf("Incorrect TestInsertPermission_Success test")
		t.FailNow()
	}
}

// Test insert with relational data
func TestInsertWithRelationalData_Success(t *testing.T) {
	_, err := permissionService.InsertWithRelationalData(1, entity.Permission{})
	if err != nil {
		t.Errorf("Incorrect TestInsertWithRelationalData_Success test")
		t.FailNow()
	}
}

// Less than stub struct
// Permission repository
type StubPermissionRepositoryImpl struct {
	Connection *gorm.DB
}

func (pri StubPermissionRepositoryImpl) FindAll() ([]*entity.Permission, *model.ErrorResBody) {
	permissions := []*entity.Permission{{Id: 1, Name: "test"}}
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindById(id int) (*entity.Permission, *model.ErrorResBody) {
	var permission entity.Permission
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByName(name string) (*entity.Permission, *model.ErrorResBody) {
	var permission entity.Permission
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByNames(names []string) ([]entity.Permission, *model.ErrorResBody) {
	var permissions []entity.Permission
	permissions = append(permissions, entity.Permission{Name: "test"})
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindNameById(id int) *string {
	permission, _ := pri.FindById(id)
	return &permission.Name
}

func (pri StubPermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) SaveWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	return &permission, nil
}
