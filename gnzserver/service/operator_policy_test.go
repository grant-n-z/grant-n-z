package service

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	operatorPolicyService OperatorPolicyService
)

// Set up
func init() {
	log.InitLogger("info")
	ctx.InitContext()
	ctx.SetUserId(1)
	ctx.SetServiceId(1)

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")

	operatorPolicyService = OperatorPolicyServiceImpl{
		OperatorPolicyRepository: StubOperatorPolicyRepositoryImpl{Connection: stubConnection},
		UserRepository:           StubUserRepositoryImpl{Connection: stubConnection},
		RoleRepository:           StubRoleRepositoryImpl{Connection: stubConnection},
	}
}

// Test instance
func TestGetOperatorPolicyServiceInstance(t *testing.T) {
	GetOperatorPolicyServiceInstance()
}

// Test get error
func TestOperatorPolicy_Get_Error(t *testing.T) {
	operatorPolicyService = OperatorPolicyServiceImpl{
		OperatorPolicyRepository: driver.OperatorPolicyRepositoryImpl{Connection: stubConnection},
		UserRepository:           driver.UserRepositoryImpl{Connection: stubConnection},
		RoleRepository:           StubRoleRepositoryImpl{Connection: stubConnection},
	}

	_, err := operatorPolicyService.Get("test")
	if err == nil {
		t.Errorf("Incorrect TestOperatorPolicy_Get_Error test")
		t.FailNow()
	}
}

// Test get success
func TestOperatorPolicy_Get_Success(t *testing.T) {
	operatorPolicyService = OperatorPolicyServiceImpl{
		OperatorPolicyRepository: StubOperatorPolicyRepositoryImpl{Connection: stubConnection},
		UserRepository:           StubUserRepositoryImpl{Connection: stubConnection},
		RoleRepository:           StubRoleRepositoryImpl{Connection: stubConnection},
	}

	_, err := operatorPolicyService.Get("1")
	if err != nil {
		t.Errorf("Incorrect TestOperatorPolicy_Get_Success test")
		t.FailNow()
	}

	_, err = operatorPolicyService.Get("")
	if err != nil {
		t.Errorf("Incorrect TestOperatorPolicy_Get_Success test")
		t.FailNow()
	}
}

// Test get by user id and role id
func TestOperatorPolicy_GetByUserIdAndRoleId_Success(t *testing.T) {
	operatorPolicyService = OperatorPolicyServiceImpl{
		OperatorPolicyRepository: StubOperatorPolicyRepositoryImpl{Connection: stubConnection},
		UserRepository:           StubUserRepositoryImpl{Connection: stubConnection},
		RoleRepository:           StubRoleRepositoryImpl{Connection: stubConnection},
	}

	_, err := operatorPolicyService.GetByUserIdAndRoleId(1, 1)
	if err != nil {
		t.Errorf("Incorrect TestOperatorPolicy_GetByUserIdAndRoleId test")
		t.FailNow()
	}
}

// Test get by user id and role id
func TestOperatorPolicy_Insert_Success(t *testing.T) {
	operatorPolicyService = OperatorPolicyServiceImpl{
		OperatorPolicyRepository: StubOperatorPolicyRepositoryImpl{Connection: stubConnection},
		UserRepository:           StubUserRepositoryImpl{Connection: stubConnection},
		RoleRepository:           StubRoleRepositoryImpl{Connection: stubConnection},
	}

	_, err := operatorPolicyService.Insert(&entity.OperatorPolicy{Id:1, UserId:1, RoleId:1})
	if err != nil {
		t.Errorf("Incorrect TestOperatorPolicy_GetByUserIdAndRoleId test")
		t.FailNow()
	}
}

// Less than stub struct
// OperatorPolicy repository
type StubOperatorPolicyRepositoryImpl struct {
	Connection *gorm.DB
}

func (opr StubOperatorPolicyRepositoryImpl) FindAll() ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	var entities []*entity.OperatorPolicy
	return entities, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	var entities []*entity.OperatorPolicy
	return entities, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody) {
	var operatorMemberRole entity.OperatorPolicy
	return &operatorMemberRole, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindRoleNameByUserId(userId int) ([]string, *model.ErrorResBody) {
	var names []string
	return names, nil
}

func (opr StubOperatorPolicyRepositoryImpl) Save(entity entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody) {
	return &entity, nil
}
