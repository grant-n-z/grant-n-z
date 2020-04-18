package groups

import (
	"bytes"
	"testing"

	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	user User
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	user = UserImpl{UserService: StubUserService{}}
}

// Test constructor
func TestGetUserInstance(t *testing.T) {
	GetUserInstance()
}

// Test method not allowed
func TestUser_MethodNotAllowed(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	user.Api(response, &request)

	if statusCode != http.StatusMethodNotAllowed {
		t.Errorf("Incorrect TestUser_MethodNotAllowed test.")
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
		t.Errorf("Incorrect TestRole_Post_BadRequest_Body test.")
		t.FailNow()
	}
}

// Test put bad request
func TestUser_Put_BadRequest_QueryParam(t *testing.T) {
	response := StubResponseWriter{}
	invalid := ioutil.NopCloser(bytes.NewReader([]byte("{\"user_email\":\"test@gmail.com\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodPut, Body: invalid}
	user.Api(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestRole_Post_BadRequest_QueryParam test.")
		t.FailNow()
	}
}

// Less than stub struct
// UserService
type StubUserService struct {
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

func (us StubUserService) GetUserById(id int) (*entity.User, *model.ErrorResBody) {
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

func (us StubUserService) GetUserGroupByUserIdAndGroupId(userId int, groupId int) (*entity.UserGroup, *model.ErrorResBody) {
	return &entity.UserGroup{}, nil
}

func (us StubUserService) GetUserServices() ([]*entity.UserService, *model.ErrorResBody) {
	return []*entity.UserService{}, nil
}

func (us StubUserService) GetUserServiceByUserIdAndServiceId(userId int, serviceId int) (*entity.UserService, *model.ErrorResBody) {
	return &entity.UserService{}, nil
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
