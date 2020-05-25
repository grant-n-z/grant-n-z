package service

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var (
	permissionService PermissionService
)

// Set up
func init() {
	log.InitLogger("info")

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	permissionService = PermissionServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
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
	_, err := permissionService.GetPermissionByUuid(uuid.New().String())
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
	_, err := permissionService.GetPermissionsByGroupUuid(uuid.New().String())
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
	_, err := permissionService.InsertWithRelationalData(uuid.New().String(), entity.Permission{})
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

func (pri StubPermissionRepositoryImpl) FindAll() ([]*entity.Permission, error) {
	permissions := []*entity.Permission{{Id: 1, Name: "test"}}
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindByUuid(uuid string) (*entity.Permission, error) {
	var permission entity.Permission
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByName(name string) (*entity.Permission, error) {
	var permission entity.Permission
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByNames(names []string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	permissions = append(permissions, entity.Permission{Name: "test"})
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindByGroupUuid(groupUuid string) ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindNameByUuid(uuid string) *string {
	return nil
}

func (pri StubPermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, error) {
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) SaveWithRelationalData(groupUuid string, permission entity.Permission) (*entity.Permission, error) {
	return &permission, nil
}
