package driver

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var operatorPolicyRepository OperatorPolicyRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	operatorPolicyRepository = GetOperatorPolicyRepositoryInstance()
}

// FindAll InternalServerError test
func TestOperatorPolicyFindAll_Error(t *testing.T) {
	_, err := operatorPolicyRepository.FindAll()
	if err == nil {
		t.Errorf("Incorrect TestOperatorPolicyFindAll_Error test")
		t.FailNow()
	}
}

// FindByUserUuid InternalServerError test
func TestOperatorPolicyFindByUserId_Error(t *testing.T) {
	_, err := operatorPolicyRepository.FindByUserUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestOperatorPolicyFindByUserId_Error test")
		t.FailNow()
	}
}

// FindByUserUuidAndRoleUuid InternalServerError test
func TestOperatorPolicyFindByUserIdAndRoleId_Error(t *testing.T) {
	_, err := operatorPolicyRepository.FindByUserUuidAndRoleUuid("uuid", "uuid")
	if err == nil {
		t.Errorf("Incorrect TestOperatorPolicyFindByUserIdAndRoleId_Error test")
		t.FailNow()
	}
}

// FindRoleNameByUserUuid InternalServerError test
func TestOperatorPolicyFindRoleNameByUserId_Error(t *testing.T) {
	_, err := operatorPolicyRepository.FindRoleNameByUserUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestOperatorPolicyFindRoleNameByUserId_Error test")
		t.FailNow()
	}
}

// Save InternalServerError test
func TestOperatorPolicySave_Error(t *testing.T) {
	_, err := operatorPolicyRepository.Save(entity.OperatorPolicy{})
	if err == nil {
		t.Errorf("Incorrect TestOperatorPolicySave_Error test")
		t.FailNow()
	}
}
