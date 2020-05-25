package driver

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var roleRepository RoleRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	roleRepository = GetRoleRepositoryInstance()
}

// FindAll InternalServerError test
func TestRoleFindAll_Error(t *testing.T) {
	_, err := roleRepository.FindAll()
	if err == nil {
		t.Errorf("Incorrect TestRoleFindAll_Error test")
		t.FailNow()
	}
}

// FindOffSetAndLimit InternalServerError test
func TestRoleFindOffSetAndLimit_Error(t *testing.T) {
	_, err := roleRepository.FindOffSetAndLimit(1, 1)
	if err == nil {
		t.Errorf("Incorrect TestRoleFindOffSetAndLimit_Error test")
		t.FailNow()
	}
}

// FindByUuid InternalServerError test
func TestRoleFindById_Error(t *testing.T) {
	_, err := roleRepository.FindByUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestRoleFindById_Error test")
		t.FailNow()
	}
}

// FindByName InternalServerError test
func TestRoleFindByName_Error(t *testing.T) {
	_, err := roleRepository.FindByName("test")
	if err == nil {
		t.Errorf("Incorrect TestRoleFindByName_Error test")
		t.FailNow()
	}
}

// FindByNames InternalServerError test
func TestRoleFindByNames_Error(t *testing.T) {
	_, err := roleRepository.FindByNames([]string{"test"})
	if err == nil {
		t.Errorf("Incorrect TestRoleFindByNames_Error test")
		t.FailNow()
	}
}

// FindByGroupUuid InternalServerError test
func TestRoleFindByGroupId_Error(t *testing.T) {
	_, err := roleRepository.FindByGroupUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestRoleFindByGroupId_Error test")
		t.FailNow()
	}
}

// FindNameByUuid is nil test
func TestRoleFindNameById_Nil(t *testing.T) {
	name := roleRepository.FindNameByUuid("uuid")
	if name != nil {
		t.Errorf("Incorrect TestRoleFindNameById_Nil test")
		t.FailNow()
	}
}

// Save InternalServerError test
func TestRoleSave_Error(t *testing.T) {
	_, err := roleRepository.Save(entity.Role{})
	if err == nil {
		t.Errorf("Incorrect TestRoleSave_Error test")
		t.FailNow()
	}
}

// SaveWithRelationalData InternalServerError test
func TestRoleSaveWithRelationalData_Error(t *testing.T) {
	_, err := roleRepository.SaveWithRelationalData("uuid", entity.Role{})
	if err == nil {
		t.Errorf("Incorrect TestRoleSaveWithRelationalData_Error test")
		t.FailNow()
	}
}
