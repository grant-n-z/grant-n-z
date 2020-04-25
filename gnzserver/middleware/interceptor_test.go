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

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var interceptor Interceptor

func init() {
	os.Setenv("SERVER_PRIVATE_KEY", "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS2dJQkFBS0NBZ0VBd1BWbWFGdjJpZ0lvL1ZITk45dytoWUZDdHNFby8rZ0gwaWtpOEh3RlBybnBpVDN2Ci8wckptRTNyV2pBWExwU1V5NGh2R1dWUzVPMzlYcUgyd3dzcU5uYlNySUd6T3NKWGx4NVNOdy9BT0VGV012YzcKVysvQk9FY0xGUVFLdnJJNTlpUDZqZS91YitkY1F6NDc1bFNFNkpHbFczRlh1dkVJelQrYlliNEZYYTVrRDJJMAppZUVHc2hKV2RmNmFKQjMzeGZlYjI0WjUwSThmdXBsQlAvZnkzSnFJNDlEeGppTm5XSS9TWC9iMzdmRGppd3BqCkRVUjFBMzVxSTJ3VC94QlM3dWpEeG56c2o5N2JocFFWeXl5d2JET3AvNWNqZ3RuUGZqRHRXL2hndExNREVMOFMKWmpKekJXdXpPdlRzWWRNYkUrWU5mcDJLTmh3R0lCODVmc2tSWFFMUll6U2pwMHRrZ3dUNDFidTB2R2F4UC9QawpTUDJtcm9SM1ZMOXVidmJWazVjRUZScnNIVWR5TVYxWGpYYStidEo5bDV2WVp4NDF4OE1teWxGOExQc0V3cFpICk1VNXRXd0l4MWVXTEJvSVpWMjNRNklQMVV1TU50ZkYyT0dkQVlwWmJFM1VaMXp6K0lJeHFoMjhSN2xnVmcyaEwKbmtENE9zNy9GTGRuT2xDK2hDbGVyWXNTbXJibnNTUVBMVitrRHlYaUN4M3NRQmRHSVV6cmI4MnhWK05IaTBHQwpURFVQTGUyOGF6NS9pTXYyZHRJTzU1Mmk1SjN4WFU1YlNvYkVLdktyN09GQmdMUTUzY0FVZWk5S1Ezd3VPL1l6CkdrR1pLd1laR1pBMHRyTXpZcWNYdUYyMkk3ZFd1Ny9xUCtnSlJ1R0tvQitVSXJUUnZ4UllHYUdKRVNzQ0F3RUEKQVFLQ0FnQnE2QXVqQ2tDZjJlNkgyMGlPQ0hLRFdVaHpKTFhvZ3MvQ2VwUW5GUzk2djFwS2RZeUFyeGplVDExMApER3pybTlxTW9ieWNIMjA3OVRlSnRNYVk3Wmluc0ZHc3pmZFZPTk42b0l3QWdiT0g5M3NncGFXM25EQTdVL0VwCjVhRm1ZaXlHMlF3Nms1SlJZYWZXZ2FhQ1NKV1NuUTgyaUtlSXBYNnc2T3JYem9YK2oxNVV2NTVJUGVxTndtY3cKM0t6ZUVkUnpWR3AveUJPNW4zSisyZVl5NE5jbnJsN2xpYUFybGlYdUJWRVRaaXViSXZtTzBXUnJ0MS8xWFpiRQpwVjYyMUg1K2EydjhqcjRxUDlqMHlSc0NCSEtDb0dVeGZMVnFDKzlRQlIrYzh4SHhTN1VKRkRQSGQzQU9zc3NNCmdTSnVXSjRKK2szRHR0a2FmWkJQUDcyRkp6VjhrRHVEcDNSZnlzb1Q0eFBRUGtreFIwK1MvY2pYUXJjcnhzY2gKRHM1SFdwM2ZxTDh2dEtuRlJ2WlMrbm1HaVVYMmdLMXpnUmJLYWdvR3p6Zzl1V1gvaTl0Zk1hZmJyQWdMV21YcwpKblhBRFU1T2duRGdTN01TU1U2Rm5pdlZJSkdKNWJBNFNBRXA1OVdSRi92ckFlcWRDNXVlUkJtcVdtMzNYck1aCkRKUnNZb0hYU041dHlXdGFqaWNpMU1YclZnUkorWlh0cEhTVFBGUUdiSzEwWmZiYzdoQjZGMExXSGpsbngyaTYKcmtRUjVDZ2NxemRndkFCeU5kSmg2bnIyWitoVjFGeUNTbFgvWFZnMllUN1ZpQzdYNTJ4cWluTzNuN0dJOHFOQwpJL0VyZkNmS0pMUGJhTVdmQ2E0TEozaE9uT0d1SU9WdTZCV2VRRHNHU0JFSndjWFdzUUtDQVFFQS90QXp4R01tCjRsT2FSNGFaRkxQRzE0bjFyQUMrbU1OTU9pRzJJeXczSWdsRGVUZkI1S05oampBUWh2NHAyVnJraU9RZDQyOTQKWHBqdnBuSGw1MnZzelpwWlFKeW1zSGFLYjdJNHdCcndaUHBTSTUzd0hYUnpLNTFNQWZpdDhsQlJSWXJ0cWFhLwp6TlRISHdtcTJ0YlloY0krN0dab1NLcndReGUxWUQra1ZCK2REeXNkSkNhQ2VxaDdyQzYrUWdDRG1HeUxMejN2ClZDUWltS0gxaWpCNVJOTDJqM1BEVkwvUVQyYkZtb0FZWlI5ZmxuUUFDV1hPWERjRlpYSzY3SGdmQW13OXR6Z2kKaDU5REllaDZJWW5nWmJ1NzJnNmVyRHo3YVA3T3FrMVhzMHVrdDVVU3p6ZEZON1dXZWIraXdCdTNvZ3g2NFYxQQpCbWJ1bWtzNkN6ZXJsd0tDQVFFQXdkdHp6bkh4dnBML2t6NmgvbjAwVlJndm9vUU9oWTZGbVdDTWdXTWcwNWRECmlMdU9WbmRBazRObXVyVkRScXBYaFNYRFZqbFU4ZnBGSlJNNFRYRFpQaHJEUXlUNEp4L2xDeUxNYkd3RmtPb1cKYXNNZUdoVGdQKzNjc25BVVhCb0FDSmFFbENad3RBaEtYdXI4YnhQdlJMaHAxc2FHTjRjVDR4bVBGb2pOMlB5RQorRXFBWHRrNlhOMCtZVy9oZGFpZkdOL1BwRGFMcEF1MTJzVGFIR1NhOVBDVkNTbzdweHMveG9Tc0xTYW0rL3Y4Cm1lV1ZNVVZaT0RTbjA3a01rbXh3SWtNTXV5eVF5V0VzWWJMM3N1bFphYVBFMU1Ndnk3bUFKR0ppYklBekdERWkKYmJMSk9pcGc1TUMySHZTNFZYeFBGZGVTK3BWUTc4QllvRnBVajhySmpRS0NBUUVBM0RpVWhPWXNkTzVNS0FUcgp5RGlYWVRDYVlrMUNiRVJkWE9CRnlhQXRCZjE3a3dmZFN2enBFem4zRHJRYTl2N1hCSGdpWEsyNkdnZVRGd2JZCjYya2EvNWFtREhGV25xdlVlVFJPVjdqd2lsVE5LSHNYU2wyYUs5ZUdHUzRUSjVqQ3BKZXRUeklPRWJqVFhyKzgKS2VZRXU1VmxUR28xTnBpRmpYYXdDcjcyQnI1THZ4QkQzenBwQ2hrU3lYeWNjZTUvelB3Q1RwSDRoWCsxWnJTUwp3UnVqc3hlZ2Y5cE10cklRRm85N3VFdDh5ZWlUZERSTTA5Sm94c25Hb0NiSDVoYnF0ZTFXYVVMYWxOdlA2VDVDClR6b1o5ZEtLUjZyYTk0Qzh4OEZ3V3o3OHpMaFRZMVl6SzJOWkx3eUJRRGVmTU9qRGpBbTlLWWl1RE5wbzNIQ24KZVlwamdRS0NBUUVBdGxJMEgzU1EzU0NabUIxdTg4OURtY2lPZkhWZ3h3R2M2dnlnQ014M1FpbGdqY2QvL2hoWQpOcVI5eUpuVDlURWQ4UTdzSVRyNGhrQlFLYWRpNjRwMzl1M3F1VXFhelFrMVBIejA3Ly9FV0YrZ3g3Wk1xRkQ3Cis0UTFiZWoxYlEzUy9FQzczaTR0RDFWQXhQYVNoZEdrMWVmdk90MHB2QzJoYVpSUE8rMWNWSGhpZ3JabTkwMnMKazB4TmNBeHVhbDgxaW9wc1dsQW1reG1rWm1WL2tQYVp1a1pPbFBrUWM0Q3dRWC9rQXU3NFc4UEo5ZCt6cWt4RAp0aFhueGJ1amRFN2lRNGIyQVUvUHVHWlkvR1g2aWx6bkIvRExuU01aMzZ2T05lb0dFVytkSG1LUHM4WlRkUTRJClpQeE9ETjB5Uk13T0FVZm5aeDlwcUtNcGQxNmRhME5ZdlFLQ0FRRUFqU29QeWdTY1pCeUNLZ09GaDNiQ3U2czUKZGdsRUl3d3gvWWFuMEVkb0wySkpzVUl1c3JrelFHZGpXODFoSmJFY09vMmRqc24wK3FUbUlBMEdnZmRXTTZFUAo0VUFYUzRSZFZObWR5K0JyNXdiUlQ3YmE5V1JVS0JFY0xGZzFsUEx2SXVsTkNJZ3k0cHlKTy9ENEVqeWNIK1E3CkQrOUVIOVRpTXFXTytCOENrU3pIMFUxN04vR0Z2RHFXU25FWDNXWVZDMS92aG1DWXJQZENYMTVCY3BLV0NkdjAKbW1IaEtJZGl3OWdIaUlPS2daWHVRanZRUW1OT0p5R0ZlMysvT2UycUJlbTVVQ0lhZGhTVFNTcDZ2dzdNZm9teApOWnNCVkg1cS9oV3Q1ZlgyV3ZoZFBIZlJFYTVFYnlUbThPZGNJc0RoN1VpVUJ0dU8rYkFDM2VBRWZIcUZlUT09Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==")
	os.Setenv("SERVER_PUBLIC_KEY", "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUF3UFZtYUZ2MmlnSW8vVkhOTjl3KwpoWUZDdHNFby8rZ0gwaWtpOEh3RlBybnBpVDN2LzBySm1FM3JXakFYTHBTVXk0aHZHV1ZTNU8zOVhxSDJ3d3NxCk5uYlNySUd6T3NKWGx4NVNOdy9BT0VGV012YzdXKy9CT0VjTEZRUUt2ckk1OWlQNmplL3ViK2RjUXo0NzVsU0UKNkpHbFczRlh1dkVJelQrYlliNEZYYTVrRDJJMGllRUdzaEpXZGY2YUpCMzN4ZmViMjRaNTBJOGZ1cGxCUC9meQozSnFJNDlEeGppTm5XSS9TWC9iMzdmRGppd3BqRFVSMUEzNXFJMndUL3hCUzd1akR4bnpzajk3YmhwUVZ5eXl3CmJET3AvNWNqZ3RuUGZqRHRXL2hndExNREVMOFNaakp6Qld1ek92VHNZZE1iRStZTmZwMktOaHdHSUI4NWZza1IKWFFMUll6U2pwMHRrZ3dUNDFidTB2R2F4UC9Qa1NQMm1yb1IzVkw5dWJ2YlZrNWNFRlJyc0hVZHlNVjFYalhhKwpidEo5bDV2WVp4NDF4OE1teWxGOExQc0V3cFpITVU1dFd3SXgxZVdMQm9JWlYyM1E2SVAxVXVNTnRmRjJPR2RBCllwWmJFM1VaMXp6K0lJeHFoMjhSN2xnVmcyaExua0Q0T3M3L0ZMZG5PbEMraENsZXJZc1NtcmJuc1NRUExWK2sKRHlYaUN4M3NRQmRHSVV6cmI4MnhWK05IaTBHQ1REVVBMZTI4YXo1L2lNdjJkdElPNTUyaTVKM3hYVTViU29iRQpLdktyN09GQmdMUTUzY0FVZWk5S1Ezd3VPL1l6R2tHWkt3WVpHWkEwdHJNellxY1h1RjIySTdkV3U3L3FQK2dKClJ1R0tvQitVSXJUUnZ4UllHYUdKRVNzQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=")
	os.Setenv("SERVER_SIGN_ALGORITHM", "rsa256")
	log.InitLogger("info")
	ctx.InitContext()
	ctx.SetUserId(1)
	ctx.SetServiceId(1)
	ctx.SetUserUuid(uuid.New())
	ctx.SetApiKey("test")
	common.InitGrantNZServerConfig("../grant_n_z_server.yaml")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	userService := service.UserServiceImpl{
		UserRepository: StubUserRepositoryImpl{Connection: stubConnection},
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
	}

	operatorPolicyService := service.OperatorPolicyServiceImpl{
		OperatorPolicyRepository: StubOperatorPolicyRepositoryImpl{Connection: stubConnection},
		UserRepository:           StubUserRepositoryImpl{Connection: stubConnection},
		RoleRepository:           StubRoleRepositoryImpl{Connection: stubConnection},
	}

	ser := service.ServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}

	policyService := service.PolicyServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
		PolicyRepository:     StubPolicyRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		GroupRepository:      StubGroupRepositoryImpl{Connection: stubConnection},
	}

	roleService := service.RoleServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
		RoleRepository: StubRoleRepositoryImpl{Connection: stubConnection},
	}

	permissionService := service.PermissionServiceImpl{
		EtcdClient: cache.EtcdClientImpl{
			Connection: stubEtcdConnection,
			Ctx:        ctx.GetCtx(),
		},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}

	serviceConfig := common.ServerConfig{
		SignedInPrivateKeyBase64: "test key",
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

// Test intercept api Key in header
func TestInterceptApiKey_Error(t *testing.T) {
	writer := StubResponseWriter{}
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Api-Key", "")
	err := interceptApiKey(writer, &request)
	if err == nil {
		t.Errorf("Incorrect TestInterceptApiKey_Error test.")
		t.FailNow()
	}
}

// Test intercept api Key in header
func TestInterceptApiKey_Success(t *testing.T) {
	writer := StubResponseWriter{}
	request := http.Request{Header: http.Header{}}
	request.Header.Set("Api-Key", "test_key")
	err := interceptApiKey(writer, &request)
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

// Test param group id
func TestParamGroupId_Error(t *testing.T) {
	request := http.Request{Header: http.Header{}, URL: &url.URL{}}
	request.URL.Host = "localhost:8080"
	request.URL.Path = "/api/v1/groups/1/user"
	_, err := ParamGroupId(&request)
	if err == nil {
		t.Errorf("Incorrect TestParamGroupId_Error test.")
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
