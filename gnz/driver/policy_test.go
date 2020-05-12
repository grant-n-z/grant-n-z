package driver

import (
	"testing"

	"net/http"

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
func TestPolicyFindAll_InternalServerError(t *testing.T) {
	_, err := policyRepository.FindAll()
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPolicyFindAll_InternalServerError test")
		t.FailNow()
	}
}

// FindOffSetAndLimit InternalServerError test
func TestPolicyFindOffSetAndLimit_InternalServerError(t *testing.T) {
	_, err := policyRepository.FindOffSetAndLimit(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPolicyFindOffSetAndLimit_InternalServerError test")
		t.FailNow()
	}
}

// FindByRoleId InternalServerError test
func TestPolicyFindByRoleId_InternalServerError(t *testing.T) {
	_, err := policyRepository.FindByRoleId(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPolicyFindByRoleId_InternalServerError test")
		t.FailNow()
	}
}

// FindById InternalServerError test
func TestPolicyFindById_InternalServerError(t *testing.T) {
	_, err := policyRepository.FindById(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPolicyFindById_InternalServerError test")
		t.FailNow()
	}
}

// indPolicyResponseOfUserByUserIdAndGroupId InternalServerError test
func TestFindPolicyResponseOfUserByUserIdAndGroupId_InternalServerError(t *testing.T) {
	_, err := policyRepository.FindPolicyResponseOfUserByUserIdAndGroupId(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestFindPolicyResponseOfUserByUserIdAndGroupId_InternalServerError test")
		t.FailNow()
	}
}

// Update InternalServerError test
func TestPolicyUpdate_InternalServerError(t *testing.T) {
	_, err := policyRepository.Update(entity.Policy{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPolicyUpdate_InternalServerError test")
		t.FailNow()
	}
}
