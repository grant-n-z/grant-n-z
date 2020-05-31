package v1

import (
	"bytes"
	"context"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"testing"

	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

var (
	ser Service
)

func init() {
	log.InitLogger("info")

	ser = ServiceImpl{
		ServiceService: StubService{},
		UserService:    StubUserService{},
		TokenProcessor: StubTokenProcessor{},
	}
}

// Test constructor
func TestGetServiceInstance(t *testing.T) {
	GetServiceInstance()
}

// Test get
func TestService_Get(t *testing.T) {
	response := StubResponseWriter{}
	request := http.Request{Header: http.Header{}, Method: http.MethodGet}
	ser.Get(response, &request)

	if statusCode != http.StatusOK {
		t.Errorf("Incorrect TestService_Get test.")
		t.FailNow()
	}
}

// Test post bad request
func TestService_Post_BadRequest(t *testing.T) {
	response := StubResponseWriter{}
	noneEmailBody := ioutil.NopCloser(bytes.NewReader([]byte("{\"username\":\"test\", \"password\":\"testtest\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodGet, Body: noneEmailBody}
	ser.Post(response, &request)

	if statusCode != http.StatusBadRequest {
		t.Errorf("Incorrect TestService_Post_BadRequest test.")
		t.FailNow()
	}
}

// Test post
func TestService_Post(t *testing.T) {
	response := StubResponseWriter{}
	body := ioutil.NopCloser(bytes.NewReader([]byte("{\"email\":\"test@gmail.com\", \"password\":\"testtest\"}")))
	request := http.Request{Header: http.Header{}, Method: http.MethodGet, Body: body}
	ser.Post(response, request.WithContext(context.WithValue(request.Context(), middleware.ScopeSecret, "secret")))

	if statusCode != http.StatusCreated {
		t.Errorf("Incorrect TestService_Post_BadRequest test.")
		t.FailNow()
	}
}

// Less than stub struct
// Service
type StubService struct {
}

func (ss StubService) GetServices() ([]*entity.Service, *model.ErrorResBody) {
	return []*entity.Service{}, nil
}

func (ss StubService) GetServiceByUuid(uuid string) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceByName(name string) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceBySecret(secret string) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GetServiceByUser(userUuid string) ([]*entity.Service, *model.ErrorResBody) {
	return []*entity.Service{}, nil
}

func (ss StubService) InsertService(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) InsertServiceWithRelationalData(service *entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &entity.Service{}, nil
}

func (ss StubService) GenerateSecret() string {
	return ""
}

// Less than stub struct
// Service
type StubUserService struct {
}

func (us StubUserService) GenInitialName() string {
	uid := uuid.New().String()
	return string([]rune(uid)[:6])
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
