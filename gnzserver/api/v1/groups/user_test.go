package groups

import (
	"bytes"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"testing"

	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	user User
)

func init() {
	log.InitLogger("info")

	user = UserImpl{
		UserService:  StubUserService{},
		GroupService: StubGroupService{},
	}
}

// Test constructor
func TestGetUserInstance(t *testing.T) {
	GetUserInstance()
}

// Test method not allowed
func TestUser_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodPatch}
	user.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestUser_MethodNotAllowed test. %d", statusCode)
		t.FailNow()
	}
}

// Test put bad request
func TestUser_Put_BadRequest_Body(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"invalid\":\"test@gmail.com\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPut, Body: invalid}
	user.Api(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestRole_Post_BadRequest_Body test. %d", statusCode)
		t.FailNow()
	}
}

// Test put
func TestUser_Put_Ok(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"user_email\":\"test@gmail.com\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPut, Body: invalid}
	user.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestUser_Put_Ok test. %d", statusCode)
		t.FailNow()
	}
}

// Test get
func TestUser_Get_Ok(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	user.Api(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestUser_Get_Ok test. %d", statusCode)
		t.FailNow()
	}
}

// Less than stub struct
// UserService
type StubUserService struct {
}

func (us StubUserService) GenInitialName() string {
	return "1234"
}

func (us StubUserService) EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (us StubUserService) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (us StubUserService) GetUserByUuid(uuid string) (*entity.User, *model.ErrorResBody) {
	return &entity.User{}, nil
}

func (us StubUserService) GetUserByEmail(email string) (*entity.User, *model.ErrorResBody) {
	return &entity.User{}, nil
}

func (us StubUserService) GetUserWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody) {
	return &model.UserWithOperatorPolicy{}, nil
}

func (us StubUserService) GetUserWithUserServiceWithServiceByEmail(email string) (*model.UserWithUserServiceWithService, *model.ErrorResBody) {
	return &model.UserWithUserServiceWithService{}, nil
}

func (us StubUserService) GetUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*entity.UserGroup, *model.ErrorResBody) {
	return &entity.UserGroup{}, nil
}

func (us StubUserService) GetUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	return []*entity.UserService{}, nil
}

func (us StubUserService) GetUserServiceByUserUuidAndServiceUuid(userUuid string, serviceUuid string) (*entity.UserService, *model.ErrorResBody) {
	return &entity.UserService{}, nil
}

func (us StubUserService) GetUserByGroupUuid(groupUuid string) ([]*model.UserResponse, *model.ErrorResBody) {
	return []*model.UserResponse{}, nil
}

func (us StubUserService) GetUserPoliciesByUserUuid(userUuid string) []structure.UserPolicy {
	return []structure.UserPolicy{}
}

func (us StubUserService) GetUserGroupsByUserUuid(userUuid string) []structure.UserGroup {
	return []structure.UserGroup{}
}

func (us StubUserService) InsertUserGroup(userGroupEntity entity.UserGroup) (*entity.UserGroup, *model.ErrorResBody) {
	return &entity.UserGroup{}, nil
}

func (us StubUserService) InsertUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	return &entity.User{}, nil
}

func (us StubUserService) InsertUserWithUserService(user entity.User, userService entity.UserService) (*entity.User, *model.ErrorResBody) {
	return &entity.User{}, nil
}

func (us StubUserService) InsertUserService(userServiceEntity entity.UserService) (*entity.UserService, *model.ErrorResBody) {
	return &entity.UserService{}, nil
}

func (us StubUserService) UpdateUser(user entity.User) (*entity.User, *model.ErrorResBody) {
	return &entity.User{}, nil
}
