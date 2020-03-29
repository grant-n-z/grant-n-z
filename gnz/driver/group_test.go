package driver

import (
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var groupRepository GroupRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	groupRepository = GetGroupRepositoryInstance()
}

// FindAll InternalServerError test
func TestGroupFindAll_InternalServerError(t *testing.T) {
	_, err := groupRepository.FindAll()
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestGroupFindAll_InternalServerError test")
	}
}

// FindById InternalServerError test
func TestGroupFindById_InternalServerError(t *testing.T) {
	_, err := groupRepository.FindById(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestGroupFindById_InternalServerError test")
	}
}

// FindByName InternalServerError test
func TestGroupFindByName_InternalServerError(t *testing.T) {
	_, err := groupRepository.FindByName("name")
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestGroupFindByName_InternalServerError test")
	}
}

// FindGroupsByUserId InternalServerError test
func TestGroupFindGroupsByUserId_InternalServerError(t *testing.T) {
	_, err := groupRepository.FindGroupsByUserId(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestGroupFindGroupsByUserId_InternalServerError test")
	}
}

// FindGroupWithUserWithPolicyGroupsByUserId InternalServerError test
func TestGroupFindGroupWithUserWithPolicyGroupsByUserId_InternalServerError(t *testing.T) {
	_, err := groupRepository.FindGroupWithUserWithPolicyGroupsByUserId(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestGroupFindGroupWithUserWithPolicyGroupsByUserId_InternalServerError test")
	}
}

// FindGroupWithUserWithPolicyGroupsByUserId InternalServerError test
func TestGroupFindGroupWithPolicyByUserIdAndGroupId_InternalServerError(t *testing.T) {
	_, err := groupRepository.FindGroupWithPolicyByUserIdAndGroupId(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestGroupFindGroupWithPolicyByUserIdAndGroupId_InternalServerError test")
	}
}

// SaveWithRelationalData InternalServerError test
func TestGroupSaveWithRelationalData_InternalServerError(t *testing.T) {
	_, err := groupRepository.SaveWithRelationalData(
		entity.Group{},
		entity.ServiceGroup{},
		entity.UserGroup{},
		entity.GroupPermission{},
		entity.GroupRole{},
		entity.Policy{},
	)

	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestGroupSaveWithRelationalData_InternalServerError test")
	}
}
