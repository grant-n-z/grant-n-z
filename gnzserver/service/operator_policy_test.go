package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var (
	operatorPolicyService OperatorPolicyService
)

// Set up
func init() {
	log.InitLogger("info")

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
		OperatorPolicyRepository: driver.RdbmsOperatorPolicyRepository{Connection: stubConnection},
		UserRepository:           driver.RdbmsUserRepository{Connection: stubConnection},
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

	_, err := operatorPolicyService.GetByUserUuidAndRoleUuid(uuid.New().String(), uuid.New().String())
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

	_, err := operatorPolicyService.Insert(&entity.OperatorPolicy{InternalId: "", UserUuid: uuid.New(), RoleUuid: uuid.New()})
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

func (opr StubOperatorPolicyRepositoryImpl) FindAll() ([]*entity.OperatorPolicy, error) {
	var entities []*entity.OperatorPolicy
	return entities, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindByUserUuid(userUuid string) ([]*entity.OperatorPolicy, error) {
	var entities []*entity.OperatorPolicy
	return entities, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindByUserUuidAndRoleUuid(userUuid string, roleUuid string) (*entity.OperatorPolicy, error) {
	var operatorMemberRole entity.OperatorPolicy
	return &operatorMemberRole, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindRoleNameByUserUuid(userUuid string) ([]string, error) {
	return []string{""}, nil
}

func (opr StubOperatorPolicyRepositoryImpl) Save(entity entity.OperatorPolicy) (*entity.OperatorPolicy, error) {
	return &entity, nil
}
