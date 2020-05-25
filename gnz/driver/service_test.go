package driver

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var serviceRepository ServiceRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	serviceRepository = GetServiceRepositoryInstance()
}

// FindAll InternalServerError test
func TestServiceFindAllError(t *testing.T) {
	_, err := serviceRepository.FindAll()
	if err == nil {
		t.Errorf("Incorrect TestServiceFindAllError test")
		t.FailNow()
	}
}

// FindOffSetAndLimit InternalServerError test
func TestServiceFindOffSetAndLimitError(t *testing.T) {
	_, err := serviceRepository.FindOffSetAndLimit(1, 1)
	if err == nil {
		t.Errorf("Incorrect TestServiceFindOffSetAndLimitError test")
		t.FailNow()
	}
}

// FindByUuid InternalServerError test
func TestServiceFindByIdError(t *testing.T) {
	_, err := serviceRepository.FindByUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestServiceFindByIdError test")
		t.FailNow()
	}
}

// FindByName InternalServerError test
func TestServiceFindByNameError(t *testing.T) {
	_, err := serviceRepository.FindByName("test")
	if err == nil {
		t.Errorf("Incorrect TestServiceFindByNameError test")
		t.FailNow()
	}
}

// FindBySecret InternalServerError test
func TestServiceFindByApiKeyError(t *testing.T) {
	_, err := serviceRepository.FindBySecret("test_api_key")
	if err == nil {
		t.Errorf("Incorrect TestServiceFindByApiKeyError test")
		t.FailNow()
	}
}

// FindNameByUuid is nil test
func TestServiceFindNameById_Nil(t *testing.T) {
	name := serviceRepository.FindNameByUuid("uuid")
	if name != nil {
		t.Errorf("Incorrect TestServiceFindNameById_Nil test")
		t.FailNow()
	}
}

// FindServicesByUserUuid  InternalServerError test
func TestServiceFindServicesByUserId_Nil(t *testing.T) {
	_, err := serviceRepository.FindServicesByUserUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestServiceFindServicesByUserId_Nil test")
		t.FailNow()
	}
}

// Save InternalServerError test
func TestServiceSaveError(t *testing.T) {
	_, err := serviceRepository.Save(entity.Service{})
	if err == nil {
		t.Errorf("Incorrect TestServiceSaveError test")
		t.FailNow()
	}
}

// SaveWithRelationalData InternalServerError test
func TestServiceSaveWithRelationalDataError(t *testing.T) {
	_, err := serviceRepository.SaveWithRelationalData(entity.Service{}, []entity.Role{{}}, []entity.Permission{{}})
	if err == nil {
		t.Errorf("Incorrect TestServiceSaveWithRelationalDataError test")
		t.FailNow()
	}
}

// Update InternalServerError test
func TestServiceUpdateError(t *testing.T) {
	_, err := serviceRepository.Update(entity.Service{})
	if err == nil {
		t.Errorf("Incorrect TestServiceSaveError test")
		t.FailNow()
	}
}
