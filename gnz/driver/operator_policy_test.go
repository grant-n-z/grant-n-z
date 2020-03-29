package driver

import (
	"testing"

	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var operatorPolicyRepository OperatorPolicyRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnectoin, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnectoin
	operatorPolicyRepository = GetOperatorPolicyRepositoryInstance()
}

// FindAll InternalServerError test
func TestOperatorPolicyFindAll_InternalServerError(t *testing.T) {
	_, err := operatorPolicyRepository.FindAll()
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestOperatorPolicyFindAll_InternalServerError test")
	}
}

// FindByUserId InternalServerError test
func TestOperatorPolicyFindByUserId_InternalServerError(t *testing.T) {
	_, err := operatorPolicyRepository.FindByUserId(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestOperatorPolicyFindByUserId_InternalServerError test")
	}
}

// FindByUserIdAndRoleId InternalServerError test
func TestOperatorPolicyFindByUserIdAndRoleId_InternalServerError(t *testing.T) {
	_, err := operatorPolicyRepository.FindByUserIdAndRoleId(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestOperatorPolicyFindByUserIdAndRoleId_InternalServerError test")
	}
}

// FindRoleNameByUserId InternalServerError test
func TestOperatorPolicyFindRoleNameByUserId_InternalServerError(t *testing.T) {
	_, err := operatorPolicyRepository.FindRoleNameByUserId(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestOperatorPolicyFindRoleNameByUserId_InternalServerError test")
	}
}

// Save InternalServerError test
func TestOperatorPolicySave_InternalServerError(t *testing.T) {
	_, err := operatorPolicyRepository.Save(entity.OperatorPolicy{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestOperatorPolicySave_InternalServerError test")
	}
}
