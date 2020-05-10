package service

import (
	"github.com/google/uuid"
	"strings"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
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
	ctx.InitContext()
	ctx.SetUserId(1)
	ctx.SetServiceId(1)
	ctx.SetUserUuid(uuid.New())

	stubConnection, _ = gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	userService = UserServiceImpl{
		UserRepository: StubUserRepositoryImpl{Connection: stubConnection},
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
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
	_, err := userService.GetUserById(1)
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
	_, err := userService.GetUserGroupByUserIdAndGroupId(1, 1)
	if err != nil {
		t.Errorf("Incorrect TestGetUserGroupByUserIdAndGroupId_Success test")
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
	_, err := userService.GetUserServiceByUserIdAndServiceId(1, 1)
	if err != nil {
		t.Errorf("Incorrect TestGetUserServiceByUserIdAndServiceId_Success test")
		t.FailNow()
	}
}

// Test insert user group
func TestInsertUserGroup_Success(t *testing.T) {
	_, err := userService.InsertUserGroup(entity.UserGroup{UserId: 1, GroupId: 1})
	if err != nil {
		t.Errorf("Incorrect TestInsertUserGroup_Success test")
		t.FailNow()
	}
}

// Test insert user
func TestInsertUser_Success(t *testing.T) {
	_, err := userService.InsertUser(entity.User{Id:1})
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

func (uri StubUserRepositoryImpl) FindById(id int) (*entity.User, *model.ErrorResBody) {
	var user entity.User
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindByEmail(email string) (*entity.User, *model.ErrorResBody) {
	var user entity.User
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody) {
	var uwo model.UserWithOperatorPolicy
	return &uwo, nil
}

func (uri StubUserRepositoryImpl) FindWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody) {
	var uus model.UserWithUserServiceWithService
	return &uus, nil
}

func (uri StubUserRepositoryImpl) FindUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody) {
	return nil, nil
}

func (uri StubUserRepositoryImpl) FindUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	var userServices []*entity.UserService
	return userServices, nil
}

func (uri StubUserRepositoryImpl) FindUserServicesOffSetAndLimit(offset int, limit int) ([]*entity.UserService, *model.ErrorResBody) {
	var userServices []*entity.UserService
	return userServices, nil
}

func (uri StubUserRepositoryImpl) FindUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	return nil, nil
}

func (uri StubUserRepositoryImpl) FindUserGroupsOffSetAndLimit(offset int, limit int) ([]*entity.UserGroup, *model.ErrorResBody) {
	return nil, nil
}

func (uri StubUserRepositoryImpl) SaveUserGroup(userGroup entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	return &userGroup, nil
}

func (uri StubUserRepositoryImpl) SaveUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	return &user, nil
}

func (uri StubUserRepositoryImpl) SaveWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody) {
	return &user, nil
}

func (uri StubUserRepositoryImpl) UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	return &user, nil
}

func (uri StubUserRepositoryImpl) SaveUserService(userService entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	return &userService, nil
}
