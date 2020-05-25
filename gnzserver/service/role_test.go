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
	roleService RoleService
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

	roleService = RoleServiceImpl{
		EtcdClient: cache.EtcdClientImpl{Connection: stubEtcdConnection},
		RoleRepository: StubRoleRepositoryImpl{Connection: stubConnection},
	}
}

// Test constructor
func TestGetRoleServiceInstance(t *testing.T) {
	GetRoleServiceInstance()
}

// Test get roles
func TestGetRoles_Success(t *testing.T) {
	_, err := roleService.GetRoles()
	if err != nil {
		t.Errorf("Incorrect TestGetRoles_Success test")
		t.FailNow()
	}
}

// Test get role by uuid
func TestGetRoleByUuid_Success(t *testing.T) {
	_, err := roleService.GetRoleByUuid(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetRoleById_Success test")
		t.FailNow()
	}
}

// Test get role by name
func TestGetRoleByName_Success(t *testing.T) {
	_, err := roleService.GetRoleByName("name")
	if err != nil {
		t.Errorf("Incorrect TestGetRoleByName_Success test")
		t.FailNow()
	}
}

// Test get role by names
func TestGetRoleByNames_Success(t *testing.T) {
	_, err := roleService.GetRoleByNames([]string{"name"})
	if err != nil {
		t.Errorf("Incorrect TestGetRoleByNames_Success test")
		t.FailNow()
	}
}

// Test get role by group uuid
func TestGetRolesByGroupUuid_Success(t *testing.T) {
	_, err := roleService.GetRolesByGroupUuid(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetRolesByGroupId_Success test")
		t.FailNow()
	}
}

// Test insert
func TestInsertRoleInsertRole_Success(t *testing.T) {
	_, err := roleService.InsertRole(&entity.Role{})
	if err != nil {
		t.Errorf("Incorrect v test")
		t.FailNow()
	}
}

// Test insert with relational data
func TestRoleInsertWithRelationalData_Success(t *testing.T) {
	_, err := roleService.InsertWithRelationalData(uuid.New().String(), entity.Role{})
	if err != nil {
		t.Errorf("Incorrect TestRoleInsertWithRelationalData_Success test")
		t.FailNow()
	}
}

// Less than stub struct
// Role repository
type StubRoleRepositoryImpl struct {
	Connection *gorm.DB
}

func (rri StubRoleRepositoryImpl) FindAll() ([]*entity.Role, error) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Role, error) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindByUuid(uuid string) (*entity.Role, error) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByName(name string) (*entity.Role, error) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByNames(names []string) ([]entity.Role, error) {
	var roles []entity.Role
	roles = append(roles, entity.Role{Name: "test"})
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindByGroupUuid(groupUuid string) ([]*entity.Role, error) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindNameByUuid(uuid string) *string {
	role, err := rri.FindByUuid(uuid)
	if err != nil {
		return nil
	}
	return &role.Name
}

func (rri StubRoleRepositoryImpl) Save(role entity.Role) (*entity.Role, error) {
	return &role, nil
}

func (rri StubRoleRepositoryImpl) SaveWithRelationalData(groupUuid string, role entity.Role) (*entity.Role, error) {
	return &role, nil
}
