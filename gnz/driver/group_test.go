package driver

import (
	"testing"

	"github.com/google/uuid"
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
func TestGroupFindAll_Error(t *testing.T) {
	_, err := groupRepository.FindAll()
	if err == nil {
		t.Errorf("Incorrect TestGroupFindAll_Error test")
		t.FailNow()
	}
}

// FindByUuid InternalServerError test
func TestGroupFindById_Error(t *testing.T) {
	_, err := groupRepository.FindByUuid(uuid.New().String())
	if err == nil {
		t.Errorf("Incorrect TestGroupFindById_Error test")
		t.FailNow()
	}
}

// FindByName InternalServerError test
func TestGroupFindByName_Error(t *testing.T) {
	_, err := groupRepository.FindByName("name")
	if err == nil {
		t.Errorf("Incorrect TestGroupFindByName_Error test")
		t.FailNow()
	}
}

// FindByUserUuid InternalServerError test
func TestGroupFindGroupsByUserId_Error(t *testing.T) {
	_, err := groupRepository.FindByUserUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestGroupFindGroupsByUserId_Error test")
		t.FailNow()
	}
}

// FindGroupWithUserWithPolicyGroupsByUserUuid InternalServerError test
func TestGroupFindGroupWithUserWithPolicyGroupsByUserId_Error(t *testing.T) {
	_, err := groupRepository.FindGroupWithUserWithPolicyGroupsByUserUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestGroupFindGroupWithUserWithPolicyGroupsByUserId_Error test")
		t.FailNow()
	}
}

// FindGroupWithUserWithPolicyGroupsByUserUuid InternalServerError test
func TestGroupFindGroupWithPolicyByUserIdAndGroupId_Error(t *testing.T) {
	_, err := groupRepository.FindGroupWithPolicyByUserUuidAndGroupUuid("uuid", "uuid")
	if err == nil {
		t.Errorf("Incorrect TestGroupFindGroupWithPolicyByUserIdAndGroupId_Error test")
		t.FailNow()
	}
}

// SaveWithRelationalData InternalServerError test
func TestGroupSaveWithRelationalData_Error(t *testing.T) {
	_, err := groupRepository.SaveWithRelationalData(
		entity.Group{},
		entity.ServiceGroup{},
		entity.UserGroup{},
		entity.GroupPermission{},
		entity.GroupRole{},
		entity.Policy{},
	)

	if err == nil {
		t.Errorf("Incorrect TestGroupSaveWithRelationalData_Error test")
		t.FailNow()
	}
}
