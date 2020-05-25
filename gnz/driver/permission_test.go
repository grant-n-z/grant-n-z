package driver

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var permissionRepository PermissionRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	permissionRepository = GetPermissionRepositoryInstance()
}

// FindAll InternalServerError test
func TestPermissionFindAll_Error(t *testing.T) {
	_, err := permissionRepository.FindAll()
	if err == nil {
		t.Errorf("Incorrect TestPermissionFindAll_Error test")
		t.FailNow()
	}
}

// FindOffSetAndLimit InternalServerError test
func TestPermissionFindOffSetAndLimit_Error(t *testing.T) {
	_, err := permissionRepository.FindOffSetAndLimit(1, 1)
	if err == nil {
		t.Errorf("Incorrect TestPermissionFindOffSetAndLimit_Error test")
		t.FailNow()
	}
}

// FindByUuid InternalServerError test
func TestPermissionFindById_Error(t *testing.T) {
	_, err := permissionRepository.FindByUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestPermissionFindById_Error test")
		t.FailNow()
	}
}

// FindByName InternalServerError test
func TestPermissionFindByName_Error(t *testing.T) {
	_, err := permissionRepository.FindByName("test")
	if err == nil {
		t.Errorf("Incorrect TestPermissionFindByName_Error test")
		t.FailNow()
	}
}

// FindByNames InternalServerError test
func TestPermissionFindByNames_Error(t *testing.T) {
	_, err := permissionRepository.FindByNames([]string{"test"})
	if err == nil {
		t.Errorf("Incorrect TestPermissionFindByNames_Error test")
		t.FailNow()
	}
}

// FindByGroupUuid InternalServerError test
func TestPermissionFindByGroupId_Error(t *testing.T) {
	_, err := permissionRepository.FindByGroupUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestPermissionFindByGroupId_Error test")
		t.FailNow()
	}
}

// FindNameByUuid name is nil test
func TestPermissionFindNameById_Nil(t *testing.T) {
	name := permissionRepository.FindNameByUuid("uuid")
	if name != nil {
		t.Errorf("Incorrect TestPermissionFindNameById_Nil test")
		t.FailNow()
	}
}

// Save InternalServerError test
func TestPermissionSave_Error(t *testing.T) {
	_, err := permissionRepository.Save(entity.Permission{})
	if err == nil {
		t.Errorf("Incorrect TestPermissionSave_Error test")
		t.FailNow()
	}
}

// SaveWithRelationalData InternalServerError test
func TestPermissionSaveWithRelationalData_Error(t *testing.T) {
	_, err := permissionRepository.SaveWithRelationalData("uuid", entity.Permission{})
	if err == nil {
		t.Errorf("Incorrect TestPermissionSaveWithRelationalData_Error test")
		t.FailNow()
	}
}
