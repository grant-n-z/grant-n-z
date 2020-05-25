package middleware

import (
	"bytes"
	"os"
	"testing"
	"time"

	"io/ioutil"
	"net/http"
	"net/url"

	"go.etcd.io/etcd/clientv3"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var interceptor Interceptor

func init() {
	os.Setenv("SERVER_PRIVATE_KEY_PATH", "../../gnz/common/test-private.key")
	os.Setenv("SERVER_PUBLIC_KEY_PATH", "../../gnz/common/test-public.key")
	os.Setenv("SERVER_SIGN_ALGORITHM", "rsa256")
	log.InitLogger("info")
	common.InitGrantNZServerConfig("../grant_n_z_server.yaml")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	userService := service.UserServiceImpl{
		UserRepository: StubUserRepositoryImpl{Connection: stubConnection},
		EtcdClient:     cache.EtcdClientImpl{Connection: stubEtcdConnection},
	}

	operatorPolicyService := service.OperatorPolicyServiceImpl{
		OperatorPolicyRepository: StubOperatorPolicyRepositoryImpl{Connection: stubConnection},
		UserRepository:           StubUserRepositoryImpl{Connection: stubConnection},
		RoleRepository:           StubRoleRepositoryImpl{Connection: stubConnection},
	}

	ser := service.ServiceImpl{
		EtcdClient:           cache.EtcdClientImpl{Connection: stubEtcdConnection},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}

	policyService := service.PolicyServiceImpl{
		EtcdClient:           cache.EtcdClientImpl{Connection: stubEtcdConnection},
		PolicyRepository:     StubPolicyRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		GroupRepository:      StubGroupRepositoryImpl{Connection: stubConnection},
	}

	roleService := service.RoleServiceImpl{
		EtcdClient:     cache.EtcdClientImpl{Connection: stubEtcdConnection},
		RoleRepository: StubRoleRepositoryImpl{Connection: stubConnection},
	}

	permissionService := service.PermissionServiceImpl{
		EtcdClient:           cache.EtcdClientImpl{Connection: stubEtcdConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}

	serviceConfig := common.ServerConfig{
		SignedInPrivateKeyPath: "test key",
	}

	tokenProcessor := TokenProcessorImpl{
		UserService:           userService,
		OperatorPolicyService: operatorPolicyService,
		Service:               ser,
		PolicyService:         policyService,
		RoleService:           roleService,
		PermissionService:     permissionService,
		ServerConfig:          serviceConfig,
	}

	interceptor = InterceptorImpl{tokenProcessor: tokenProcessor}
}

// Test constructor
func TestGetInterceptorInstance(t *testing.T) {
	GetInterceptorInstance()
}

// Test intercept header
func TestInterceptHeader_Error(t *testing.T) {
	writer := StubResponseWriter{}
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Content-Type", "html/text")
	err := interceptHeader(writer, &request)
	if err == nil {
		t.Errorf("Incorrect TestInterceptHeader_Error test.")
		t.FailNow()
	}
}

// Test intercept header
func TestInterceptHeader_Success(t *testing.T) {
	writer := StubResponseWriter{}
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Content-Type", "application/json")
	err := interceptHeader(writer, &request)
	if err != nil {
		t.Errorf("Incorrect TestInterceptHeader_Success test.")
		t.FailNow()
	}
}

// Test intercept api ClientSecret in header
func TestInterceptApiKey_Error(t *testing.T) {
	writer := StubResponseWriter{}
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Client-Secret", "")
	_, err := interceptClientSecret(writer, &request)
	if err == nil {
		t.Errorf("Incorrect TestInterceptApiKey_Error test.")
		t.FailNow()
	}
}

// Test intercept api ClientSecret in header
func TestInterceptApiKey_Success(t *testing.T) {
	writer := StubResponseWriter{}
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Client-Secret", "test_key")
	_, err := interceptClientSecret(writer, &request)
	if err != nil {
		t.Errorf("Incorrect TestInterceptApiKey_Success test.")
		t.FailNow()
	}
}

// Test validate header
func TestValidateHeader_Error(t *testing.T) {
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Content-Type", "html/text")
	err := validateHeader(&request)
	if err == nil {
		t.Errorf("Incorrect TestValidateHeader test.")
		t.FailNow()
	}
}

// Test validate header
func TestValidateHeader_Success(t *testing.T) {
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Content-Type", "application/json")
	err := validateHeader(&request)
	if err != nil {
		t.Errorf("Incorrect TestValidateHeader test.")
		t.FailNow()
	}
}

// Test bind request body
func TestBindBody_Error(t *testing.T) {
	writer := StubResponseWriter{}
	body := ioutil.NopCloser(bytes.NewReader([]byte("")))
	request := http.Request{Header: http.Header{}, Body: body}
	err := BindBody(writer, &request, nil)
	if err == nil {
		t.Errorf("Incorrect TestBindBody_Error test.")
		t.FailNow()
	}

	body = ioutil.NopCloser(bytes.NewReader([]byte("{\"username\":\"test\", \"password\":\"testtest\"}")))
	request = http.Request{Header: http.Header{}, Body: body}
	err = BindBody(writer, &request, nil)
	if err == nil {
		t.Errorf("Incorrect TestBindBody_Error test.")
		t.FailNow()
	}
}

// Test bind request body
func TestBindBody_Success(t *testing.T) {
	writer := StubResponseWriter{}
	body := ioutil.NopCloser(bytes.NewReader([]byte("{\"username\":\"test\", \"email\":\"test@gmail.com\", \"password\":\"testtest\"}")))
	request := http.Request{Header: http.Header{}, Body: body}

	var userEntity *entity.User
	err := BindBody(writer, &request, &userEntity)
	if err != nil {
		t.Errorf("Incorrect TestBindBody_Success test.")
		t.FailNow()
	}
}

// Test bind request body
func TestValidateBody_Error(t *testing.T) {
	writer := StubResponseWriter{}
	user := entity.User{
		Username: "test",
		Email:    "",
		Password: "testtest",
	}
	err := ValidateBody(writer, user)
	if err == nil {
		t.Errorf("Incorrect TestValidateBody_Error test.")
		t.FailNow()
	}
}

// Test bind request body
func TestValidateBody_Success(t *testing.T) {
	writer := StubResponseWriter{}
	user := entity.User{
		Username: "test",
		Email:    "test@gmail.com",
		Password: "testtest",
	}
	err := ValidateBody(writer, user)
	if err != nil {
		t.Errorf("Incorrect TestValidateBody_Success test.")
		t.FailNow()
	}
}

// Test token for password
func TestValidateTokenRequest_Password_Success(t *testing.T) {
	writer := StubResponseWriter{}
	tokenRequest := model.TokenRequest{
		GrantType: "password",
		Email:     "test@gmail.com",
		Password:  "testtest",
	}
	err := ValidateTokenRequest(writer, &tokenRequest)
	if err != nil {
		t.Errorf("Incorrect TestValidateTokenRequest_Password_Success test.")
		t.FailNow()
	}
}

// Test token for refresh type
func TestValidateTokenRequest_RefreshToken_Success(t *testing.T) {
	writer := StubResponseWriter{}
	tokenRequest := model.TokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: "token",
	}
	err := ValidateTokenRequest(writer, &tokenRequest)
	if err != nil {
		t.Errorf("Incorrect TestValidateTokenRequest_RefreshToken_Success test.")
		t.FailNow()
	}
}

// Test token bad request of email
func TestValidateTokenRequest_InvalidEmail_BadRequest(t *testing.T) {
	writer := StubResponseWriter{}
	tokenRequest := model.TokenRequest{
		GrantType: "password",
		Email:     "test@gmail.com",
		Password:  "test",
	}
	err := ValidateTokenRequest(writer, &tokenRequest)
	if err == nil {
		t.Errorf("Incorrect TestValidateTokenRequest_InvalidEmail_BadRequest test.")
		t.FailNow()
	}
}

// Test token bad request of email
func TestValidateTokenRequest_InvalidGrantType_BadRequest(t *testing.T) {
	writer := StubResponseWriter{}
	tokenRequest := model.TokenRequest{
		GrantType: "none",
		Email:     "test@gmail.com",
		Password:  "testtest",
	}
	err := ValidateTokenRequest(writer, &tokenRequest)
	if err == nil {
		t.Errorf("Incorrect TestValidateTokenRequest_InvalidGrantType_BadRequest test.")
		t.FailNow()
	}
}

// Test refresh token bad request
func TestValidateTokenRequest_InvalidRefreshToken_BadRequest(t *testing.T) {
	writer := StubResponseWriter{}
	tokenRequest := model.TokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: "",
	}
	err := ValidateTokenRequest(writer, &tokenRequest)
	if err == nil {
		t.Errorf("Incorrect TestValidateTokenRequest_InvalidRefreshToken_BadRequest test.")
		t.FailNow()
	}
}

// Test param group id
func TestParamGroupId(t *testing.T) {
	request := http.Request{Header: http.Header{}, URL: &url.URL{}}
	request.URL.Host = "localhost:8080"
	request.URL.Path = "/api/v1/groups/62e1a5b6-9ac3-4024-918d-5012375d5108/user"
	id := ParamGroupUuid(&request)
	if id != "" {
		t.Errorf("Incorrect TestParamGroupId test. " + id)
		t.FailNow()
	}
}

type StubResponseWriter struct {
}

func (w StubResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w StubResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (w StubResponseWriter) WriteHeader(statusCode int) {
}
