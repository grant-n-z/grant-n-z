package driver

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var policyRepository PolicyRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	policyRepository = GetPolicyRepositoryInstance()
}

// FindAll InternalServerError test
func TestPolicyFindAll_Error(t *testing.T) {
	_, err := policyRepository.FindAll()
	if err == nil {
		t.Errorf("Incorrect TestPolicyFindAll_Error test")
		t.FailNow()
	}
}

// FindOffSetAndLimit InternalServerError test
func TestPolicyFindOffSetAndLimit_Error(t *testing.T) {
	_, err := policyRepository.FindOffSetAndLimit(1, 1)
	if err == nil {
		t.Errorf("Incorrect TestPolicyFindOffSetAndLimit_Error test")
		t.FailNow()
	}
}

// FindByRoleUuid InternalServerError test
func TestPolicyFindByRoleId_Error(t *testing.T) {
	_, err := policyRepository.FindByRoleUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestPolicyFindByRoleId_Error test")
		t.FailNow()
	}
}

// FindByUuid InternalServerError test
func TestPolicyFindById_Error(t *testing.T) {
	_, err := policyRepository.FindByUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestPolicyFindById_Error test")
		t.FailNow()
	}
}

// TestFindPolicyOfUserGroupByUserUuidAndGroupUuid InternalServerError test
func TestFindPolicyOfUserGroupByUserUuidAndGroupUuid_Error(t *testing.T) {
	_, err := policyRepository.FindPolicyOfUserGroupByUserUuidAndGroupUuid("uuid", "uuid")
	if err == nil {
		t.Errorf("Incorrect TestFindPolicyOfUserGroupByUserUuidAndGroupUuid_Error test")
		t.FailNow()
	}
}

// FindPolicyOfUserServiceByUserUuidAndServiceUuid InternalServerError test
func TestFindPolicyOfUserServiceByUserUuidAndGroupUuid_Error(t *testing.T) {
	_, err := policyRepository.FindPolicyOfUserServiceByUserUuidAndServiceUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect FindPolicyOfUserServiceByUserUuidAndServiceUuid test")
		t.FailNow()
	}
}

// Update InternalServerError test
func TestPolicyUpdate_Error(t *testing.T) {
	_, err := policyRepository.Update(entity.Policy{})
	if err == nil {
		t.Errorf("Incorrect TestPolicyUpdate_Error test")
		t.FailNow()
	}
}
