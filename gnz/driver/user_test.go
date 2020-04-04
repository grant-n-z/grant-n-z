package driver

import (
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"net/http"
	"testing"
)

var userRepository UserRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	userRepository = GetUserRepositoryInstance()
}

// FindById InternalServerError test
func TestUserFindById_InternalServerError(t *testing.T) {
	_, err := userRepository.FindById(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserFindById_InternalServerError test")
		t.FailNow()
	}
}

// FindByEmail InternalServerError test
func TestUserFindByEmail_InternalServerError(t *testing.T) {
	_, err := userRepository.FindByEmail("test@gmail.com")
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserFindByEmail_InternalServerError test")
		t.FailNow()
	}
}

// FindWithOperatorPolicyByEmail InternalServerError test
func TestUserFindWithOperatorPolicyByEmail_InternalServerError(t *testing.T) {
	_, err := userRepository.FindWithOperatorPolicyByEmail("test@gmail.com")
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserFindWithOperatorPolicyByEmail_InternalServerError test")
		t.FailNow()
	}
}

// FindUserGroupByUserIdAndGroupId InternalServerError test
func TestUserFindUserGroupByUserIdAndGroupId_InternalServerError(t *testing.T) {
	_, err := userRepository.FindUserGroupByUserIdAndGroupId(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserFindUserGroupByUserIdAndGroupId_InternalServerError test")
		t.FailNow()
	}
}

// FindUserServices InternalServerError test
func TestUserFindUserServices_InternalServerError(t *testing.T) {
	_, err := userRepository.FindUserServices()
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserFindUserServices_InternalServerError test")
		t.FailNow()
	}
}

// FindUserServicesOffSetAndLimit InternalServerError test
func TestUserFindUserServicesOffSetAndLimit_InternalServerError(t *testing.T) {
	_, err := userRepository.FindUserServicesOffSetAndLimit(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserFindUserServicesOffSetAndLimit_InternalServerError test")
		t.FailNow()
	}
}

// FindUserServiceByUserIdAndServiceId InternalServerError test
func TestUserFindUserServiceByUserIdAndServiceId_InternalServerError(t *testing.T) {
	_, err := userRepository.FindUserServiceByUserIdAndServiceId(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserFindUserServiceByUserIdAndServiceId_InternalServerError test")
		t.FailNow()
	}
}

// SaveUserGroup InternalServerError test
func TestUserSaveUserGroup_InternalServerError(t *testing.T) {
	_, err := userRepository.SaveUserGroup(entity.UserGroup{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserSaveUserGroup_InternalServerError test")
		t.FailNow()
	}
}

// SaveUser InternalServerError test
func TestUserSaveUser_InternalServerError(t *testing.T) {
	_, err := userRepository.SaveUser(entity.User{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserSaveUser_InternalServerError test")
		t.FailNow()
	}
}

// SaveWithUserService InternalServerError test
func TestUserSaveWithUserService_InternalServerError(t *testing.T) {
	_, err := userRepository.SaveWithUserService(entity.User{}, entity.UserService{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestSaveWithUserService_InternalServerError test")
		t.FailNow()
	}
}

// SaveUserService InternalServerError test
func TestUserSaveUserService_InternalServerError(t *testing.T) {
	_, err := userRepository.SaveUserService(entity.UserService{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserSaveUserService_InternalServerError test")
		t.FailNow()
	}
}

// UpdateUser InternalServerError test
func TestUserUpdateUser_InternalServerError(t *testing.T) {
	_, err := userRepository.UpdateUser(entity.User{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestUserUpdateUser_InternalServerError test")
		t.FailNow()
	}
}
