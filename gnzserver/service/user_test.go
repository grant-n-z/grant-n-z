package service

import (
	"strings"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	userService UserService
)

// Set up
func init() {
	log.InitLogger("info")

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	userService = UserServiceImpl{
		UserRepository: StubUserRepositoryImpl{Connection: stubConnection},
		EtcdClient:     cache.EtcdClientImpl{Connection: stubEtcdConnection},
	}
}

// Test constructor
func TestGetUserServiceInstance(t *testing.T) {
	GetUserServiceInstance()
}

// Test generate initial name
func TestGenInitialName(t *testing.T) {
	name := userService.GenInitialName()
	if strings.EqualFold(name, "") {
		t.Errorf("Incorrect GenInitialName test")
		t.FailNow()
	}
}

// Test encrypt password
func TestEncryptPw(t *testing.T) {
	pw := userService.EncryptPw("test")
	if strings.EqualFold(pw, "") {
		t.Errorf("Incorrect TestEncryptPw test")
		t.FailNow()
	}

	pw = userService.EncryptPw("1@#%&*()-+=:/")
	if strings.EqualFold(pw, "") {
		t.Errorf("Incorrect TestEncryptPw test")
		t.FailNow()
	}
}

// Test compare password success
func TestComparePw_Success(t *testing.T) {
	pw := userService.EncryptPw("test")
	result := userService.ComparePw(pw, "test")
	if !result {
		t.Errorf("Incorrect TestEncryptPw test")
		t.FailNow()
	}
}

// Test compare password failed
func TestComparePw_Failed(t *testing.T) {
	pw := userService.EncryptPw("test")
	result := userService.ComparePw(pw, "test123")
	if result {
		t.Errorf("Incorrect TestEncryptPw test")
		t.FailNow()
	}
}

// Test get user by id
func TestGetUserById_Success(t *testing.T) {
	_, err := userService.GetUserByUuid(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetUserById_Success test")
		t.FailNow()
	}
}

// Test get user by email
func TestGetUserByEmail_Success(t *testing.T) {
	_, err := userService.GetUserByEmail("test@gmail.com")
	if err != nil {
		t.Errorf("Incorrect TestGetUserByEmail_Success test")
		t.FailNow()
	}
}

// Test get user with operator policy by email
func TestGetUserWithOperatorPolicyByEmail_Success(t *testing.T) {
	_, err := userService.GetUserWithOperatorPolicyByEmail("test@gmail.com")
	if err != nil {
		t.Errorf("Incorrect TestGetUserWithOperatorPolicyByEmail_Success test")
		t.FailNow()
	}
}

// Test user with user service with service by email
func TestGetUserWithUserServiceWithServiceByEmail_Success(t *testing.T) {
	_, err := userService.GetUserWithUserServiceWithServiceByEmail("test@gmail.com")
	if err != nil {
		t.Errorf("Incorrect TestGetUserWithUserServiceWithServiceByEmail_Success test")
		t.FailNow()
	}
}

// Test user group by user id and group id
func TestGetUserGroupByUserIdAndGroupId_Success(t *testing.T) {
	_, err := userService.GetUserGroupByUserUuidAndGroupUuid(uuid.New().String(), uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetUserGroupByUserIdAndGroupId_Success test")
		t.FailNow()
	}
}

// Test get user by group id
func TestGetUserByGroupId_Success(t *testing.T) {
	_, err := userService.GetUserByGroupUuid(uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetUserByGroupId_Success test")
		t.FailNow()
	}
}

// Test user services
func TestGetUserServices_Success(t *testing.T) {
	_, err := userService.GetUserServices()
	if err != nil {
		t.Errorf("Incorrect TestGetUserServices_Success test")
		t.FailNow()
	}
}

// Test get user service by user id and service id
func TestGetUserServiceByUserIdAndServiceId_Success(t *testing.T) {
	_, err := userService.GetUserServiceByUserUuidAndServiceUuid(uuid.New().String(), uuid.New().String())
	if err != nil {
		t.Errorf("Incorrect TestGetUserServiceByUserIdAndServiceId_Success test")
		t.FailNow()
	}
}

// Test get user policy in etcd by user uuid
func TestGetUserPoliciesByUserUuid_Success(t *testing.T) {
	policies := userService.GetUserPoliciesByUserUuid(uuid.New().String())
	if policies != nil {
		t.Errorf("Incorrect TestGetUserPoliciesByUserUuid_Success test")
		t.FailNow()
	}
}

// Test get user service in etcd by user uuid
func TestGetUserGroupsByUserUuid_Success(t *testing.T) {
	policies := userService.GetUserGroupsByUserUuid(uuid.New().String())
	if policies != nil {
		t.Errorf("Incorrect TestGetUserGroupsByUserUuid_Success test")
		t.FailNow()
	}
}

// Test insert user group
func TestInsertUserGroup_Success(t *testing.T) {
	_, err := userService.InsertUserGroup(entity.UserGroup{UserUuid: uuid.New(), GroupUuid: uuid.New()})
	if err != nil {
		t.Errorf("Incorrect TestInsertUserGroup_Success test")
		t.FailNow()
	}
}

// Test insert user
func TestInsertUser_Success(t *testing.T) {
	_, err := userService.InsertUser(entity.User{Id: 1})
	if err != nil {
		t.Errorf("Incorrect TestInsertUser_Success test")
		t.FailNow()
	}
}

// Test insert user group with user service
func TestInsertUserWithUserService_Success(t *testing.T) {
	_, err := userService.InsertUserWithUserService(entity.User{}, entity.UserService{})
	if err != nil {
		t.Errorf("Incorrect TestInsertUserWithUserService_Success test")
		t.FailNow()
	}
}

// Test insert user service
func TestInsertUserService_Success(t *testing.T) {
	_, err := userService.InsertUserService(entity.UserService{})
	if err != nil {
		t.Errorf("Incorrect TestInsertUserService_Success test")
		t.FailNow()
	}
}

// Test update user
func TestUpdateUser_Success(t *testing.T) {
	var user entity.User
	_, err := userService.UpdateUser(user)
	if err != nil {
		t.Errorf("Incorrect TestUpdateUser_Success test")
		t.FailNow()
	}
}

// Less than stub struct
// OperatorPolicy repository
type StubUserRepositoryImpl struct {
	Connection *gorm.DB
}

func (uri StubUserRepositoryImpl) FindByUuid(uuid string) (*entity.User, error) {
	var user entity.User
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindByGroupUuid(groupUuid string) ([]*entity.User, error) {
	var users []*entity.User
	return users, nil
}

func (uri StubUserRepositoryImpl) FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, error) {
	var uwo model.UserWithOperatorPolicy
	return &uwo, nil
}

func (uri StubUserRepositoryImpl) FindWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, error) {
	var uus model.UserWithUserServiceWithService
	return &uus, nil
}

func (uri StubUserRepositoryImpl) FindUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*entity.UserGroup, error) {
	var userGroup entity.UserGroup
	return &userGroup, nil
}

func (uri StubUserRepositoryImpl) FindUserServices() ([]*entity.UserService, error) {
	var userServices []*entity.UserService
	return userServices, nil
}

func (uri StubUserRepositoryImpl) FindUserServicesOffSetAndLimit(offset int, limit int) ([]*entity.UserService, error) {
	var userServices []*entity.UserService
	return userServices, nil
}

func (uri StubUserRepositoryImpl) FindUserGroupsOffSetAndLimit(offset int, limit int) ([]*entity.UserGroup, error) {
	var userGroups []*entity.UserGroup
	return userGroups, nil
}

func (uri StubUserRepositoryImpl) FindUserServiceByUserUuidAndServiceUuid(userUuid string, serviceUuid string) (*entity.UserService, error) {
	var userService entity.UserService
	return &userService, nil
}

func (uri StubUserRepositoryImpl) SaveUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, error) {
	return &userGroup, nil
}

func (uri StubUserRepositoryImpl) SaveUser(user entity.User) (*entity.User, error) {
	return &user, nil
}

func (uri StubUserRepositoryImpl) SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, error) {
	return &user, nil
}

func (uri StubUserRepositoryImpl) SaveUserService(userService entity.UserService) (*entity.UserService, error) {
	return &userService, nil
}

func (uri StubUserRepositoryImpl) UpdateUser(user entity.User) (*entity.User, error) {
	return &user, nil
}
