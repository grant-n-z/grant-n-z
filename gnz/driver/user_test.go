package driver

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var userRepository UserRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	userRepository = GetUserRepositoryInstance()
}

// FindByUuid InternalServerError test
func TestUserFindById_Error(t *testing.T) {
	_, err := userRepository.FindByUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestUserFindById_Error test")
		t.FailNow()
	}
}

// FindByEmail InternalServerError test
func TestUserFindByEmail_Error(t *testing.T) {
	_, err := userRepository.FindByEmail("test@gmail.com")
	if err == nil {
		t.Errorf("Incorrect TestUserFindByEmail_Error test")
		t.FailNow()
	}
}

// FindByUuid InternalServerError test
func TestUserFindByGroupId_Error(t *testing.T) {
	_, err := userRepository.FindByGroupUuid("uuid")
	if err == nil {
		t.Errorf("Incorrect TestUserFindByGroupId_Error test")
		t.FailNow()
	}
}

// FindWithOperatorPolicyByEmail InternalServerError test
func TestUserFindWithOperatorPolicyByEmail_Error(t *testing.T) {
	_, err := userRepository.FindWithOperatorPolicyByEmail("test@gmail.com")
	if err == nil {
		t.Errorf("Incorrect TestUserFindWithOperatorPolicyByEmail_Error test")
		t.FailNow()
	}
}

// FindUserGroupByUserUuidAndGroupUuid InternalServerError test
func TestUserFindUserGroupByUserIdAndGroupId_Error(t *testing.T) {
	_, err := userRepository.FindUserGroupByUserUuidAndGroupUuid("uuid", "uuid")
	if err == nil {
		t.Errorf("Incorrect TestUserFindUserGroupByUserIdAndGroupId_Error test")
		t.FailNow()
	}
}

// FindUserServices InternalServerError test
func TestUserFindUserServices_Error(t *testing.T) {
	_, err := userRepository.FindUserServices()
	if err == nil {
		t.Errorf("Incorrect TestUserFindUserServices_Error test")
		t.FailNow()
	}
}

// FindUserServicesOffSetAndLimit InternalServerError test
func TestUserFindUserServicesOffSetAndLimit_Error(t *testing.T) {
	_, err := userRepository.FindUserServicesOffSetAndLimit(1, 1)
	if err == nil {
		t.Errorf("Incorrect TestUserFindUserServicesOffSetAndLimit_Error test")
		t.FailNow()
	}
}

// FindUserGroupsOffSetAndLimit InternalServerError test
func TestUserFindUserGroupsOffSetAndLimit_Error(t *testing.T) {
	_, err := userRepository.FindUserGroupsOffSetAndLimit(1, 1)
	if err == nil {
		t.Errorf("Incorrect TestUserFindUserGroupsOffSetAndLimit_Error test")
		t.FailNow()
	}
}

// FindUserServiceByUserUuidAndServiceUuid InternalServerError test
func TestUserFindUserServiceByUserIdAndServiceId_Error(t *testing.T) {
	_, err := userRepository.FindUserServiceByUserUuidAndServiceUuid("uuid", "uuid")
	if err == nil {
		t.Errorf("Incorrect TestUserFindUserServiceByUserIdAndServiceId_Error test")
		t.FailNow()
	}
}

// SaveUserGroup InternalServerError test
func TestUserSaveUserGroup_Error(t *testing.T) {
	_, err := userRepository.SaveUserGroup(entity.UserGroup{})
	if err == nil {
		t.Errorf("Incorrect TestUserSaveUserGroup_Error test")
		t.FailNow()
	}
}

// SaveUser InternalServerError test
func TestUserSaveUser_Error(t *testing.T) {
	_, err := userRepository.SaveUser(entity.User{})
	if err == nil {
		t.Errorf("Incorrect TestUserSaveUser_Error test")
		t.FailNow()
	}
}

// SaveWithUserService InternalServerError test
func TestUserSaveWithUserService_Error(t *testing.T) {
	_, err := userRepository.SaveWithUserService(entity.User{}, entity.UserService{})
	if err == nil {
		t.Errorf("Incorrect TestSaveWithUserService_Error test")
		t.FailNow()
	}
}

// SaveUserService InternalServerError test
func TestUserSaveUserService_Error(t *testing.T) {
	_, err := userRepository.SaveUserService(entity.UserService{})
	if err == nil {
		t.Errorf("Incorrect TestUserSaveUserService_Error test")
		t.FailNow()
	}
}

// UpdateUser InternalServerError test
func TestUserUpdateUser_Error(t *testing.T) {
	_, err := userRepository.UpdateUser(entity.User{})
	if err == nil {
		t.Errorf("Incorrect TestUserUpdateUser_Error test")
		t.FailNow()
	}
}
