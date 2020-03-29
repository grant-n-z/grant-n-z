package driver

import (
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"net/http"
	"testing"
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
func TestServiceFindAll_InternalServerError(t *testing.T) {
	_, err := serviceRepository.FindAll()
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceFindAll_InternalServerError test")
	}
}

// FindOffSetAndLimit InternalServerError test
func TestServiceFindOffSetAndLimit_InternalServerError(t *testing.T) {
	_, err := serviceRepository.FindOffSetAndLimit(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceFindOffSetAndLimit_InternalServerError test")
	}
}

// FindById InternalServerError test
func TestServiceFindById_InternalServerError(t *testing.T) {
	_, err := serviceRepository.FindById(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceFindById_InternalServerError test")
	}
}

// FindByName InternalServerError test
func TestServiceFindByName_InternalServerError(t *testing.T) {
	_, err := serviceRepository.FindByName("test")
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceFindByName_InternalServerError test")
	}
}

// FindByApiKey InternalServerError test
func TestServiceFindByApiKey_InternalServerError(t *testing.T) {
	_, err := serviceRepository.FindByApiKey("test_api_key")
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceFindByApiKey_InternalServerError test")
	}
}

// FindNameById is nil test
func TestServiceFindNameById_Nil(t *testing.T) {
	name := serviceRepository.FindNameById(1)
	if name != nil {
		t.Errorf("Incorrect TestServiceFindNameById_Nil test")
	}
}

// FindNameByApiKey is nil test
func TestServiceFindNameByApiKey_Nil(t *testing.T) {
	apiKey := serviceRepository.FindNameByApiKey("test_api_key")
	if apiKey != nil {
		t.Errorf("Incorrect TestServiceFindNameByApiKey_Nil test")
	}
}

// FindServicesByUserId  InternalServerError test
func TestServiceFindServicesByUserId_Nil(t *testing.T) {
	_, err := serviceRepository.FindServicesByUserId(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceFindServicesByUserId_Nil test")
	}
}

// Save InternalServerError test
func TestServiceSave_InternalServerError(t *testing.T) {
	_, err := serviceRepository.Save(entity.Service{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceSave_InternalServerError test")
	}
}

// SaveWithRelationalData InternalServerError test
func TestServiceSaveWithRelationalData_InternalServerError(t *testing.T) {
	_, err := serviceRepository.SaveWithRelationalData(entity.Service{}, []entity.Role{{}}, []entity.Permission{{}})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceSaveWithRelationalData_InternalServerError test")
	}
}

// Update InternalServerError test
func TestServiceUpdate_InternalServerError(t *testing.T) {
	_, err := serviceRepository.Update(entity.Service{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestServiceSave_InternalServerError test")
	}
}
