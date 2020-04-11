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
	roleService RoleService
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

	roleService = RoleServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
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

// Test get role by id
func TestGetRoleById_Success(t *testing.T) {
	_, err := roleService.GetRoleById(1)
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

// Test get role by group id
func TestGetRolesByGroupId_Success(t *testing.T) {
	_, err := roleService.GetRolesByGroupId(1)
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
	_, err := roleService.InsertWithRelationalData(1, entity.Role{})
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

func (rri StubRoleRepositoryImpl) FindAll() ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindById(id int) (*entity.Role, *model.ErrorResBody) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByName(name string) (*entity.Role, *model.ErrorResBody) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByNames(names []string) ([]entity.Role, *model.ErrorResBody) {
	var roles []entity.Role
	roles = append(roles, entity.Role{Name: "test"})
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindNameById(id int) *string {
	role, _ := rri.FindById(id)
	return &role.Name
}

func (rri StubRoleRepositoryImpl) Save(role entity.Role) (*entity.Role, *model.ErrorResBody) {
	return &role, nil
}

func (rri StubRoleRepositoryImpl) SaveWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody) {
	return &role, nil
}
