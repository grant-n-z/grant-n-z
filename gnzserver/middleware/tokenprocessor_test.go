package middleware

import (
	"fmt"
	"testing"

	"encoding/base64"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var (
	tokenProcessor TokenProcessor
)

func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")

	userService := service.UserServiceImpl{
		UserRepository: StubUserRepositoryImpl{Connection: stubConnection},
		EtcdClient:     StubEtcdlClient{},
	}

	operatorPolicyService := service.OperatorPolicyServiceImpl{
		OperatorPolicyRepository: StubOperatorPolicyRepositoryImpl{Connection: stubConnection},
		UserRepository:           StubUserRepositoryImpl{Connection: stubConnection},
		RoleRepository:           StubRoleRepositoryImpl{Connection: stubConnection},
	}

	ser := service.ServiceImpl{
		EtcdClient:           StubEtcdlClient{},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}

	policyService := service.PolicyServiceImpl{
		EtcdClient:           StubEtcdlClient{},
		PolicyRepository:     StubPolicyRepositoryImpl{Connection: stubConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
		RoleRepository:       StubRoleRepositoryImpl{Connection: stubConnection},
		ServiceRepository:    StubServiceRepositoryImpl{Connection: stubConnection},
		GroupRepository:      StubGroupRepositoryImpl{Connection: stubConnection},
	}

	roleService := service.RoleServiceImpl{
		EtcdClient:     StubEtcdlClient{},
		RoleRepository: StubRoleRepositoryImpl{Connection: stubConnection},
	}

	permissionService := service.PermissionServiceImpl{
		EtcdClient:           StubEtcdlClient{},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}

	privateKey, _ := base64.StdEncoding.DecodeString("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS2dJQkFBS0NBZ0VBd1BWbWFGdjJpZ0lvL1ZITk45dytoWUZDdHNFby8rZ0gwaWtpOEh3RlBybnBpVDN2Ci8wckptRTNyV2pBWExwU1V5NGh2R1dWUzVPMzlYcUgyd3dzcU5uYlNySUd6T3NKWGx4NVNOdy9BT0VGV012YzcKVysvQk9FY0xGUVFLdnJJNTlpUDZqZS91YitkY1F6NDc1bFNFNkpHbFczRlh1dkVJelQrYlliNEZYYTVrRDJJMAppZUVHc2hKV2RmNmFKQjMzeGZlYjI0WjUwSThmdXBsQlAvZnkzSnFJNDlEeGppTm5XSS9TWC9iMzdmRGppd3BqCkRVUjFBMzVxSTJ3VC94QlM3dWpEeG56c2o5N2JocFFWeXl5d2JET3AvNWNqZ3RuUGZqRHRXL2hndExNREVMOFMKWmpKekJXdXpPdlRzWWRNYkUrWU5mcDJLTmh3R0lCODVmc2tSWFFMUll6U2pwMHRrZ3dUNDFidTB2R2F4UC9QawpTUDJtcm9SM1ZMOXVidmJWazVjRUZScnNIVWR5TVYxWGpYYStidEo5bDV2WVp4NDF4OE1teWxGOExQc0V3cFpICk1VNXRXd0l4MWVXTEJvSVpWMjNRNklQMVV1TU50ZkYyT0dkQVlwWmJFM1VaMXp6K0lJeHFoMjhSN2xnVmcyaEwKbmtENE9zNy9GTGRuT2xDK2hDbGVyWXNTbXJibnNTUVBMVitrRHlYaUN4M3NRQmRHSVV6cmI4MnhWK05IaTBHQwpURFVQTGUyOGF6NS9pTXYyZHRJTzU1Mmk1SjN4WFU1YlNvYkVLdktyN09GQmdMUTUzY0FVZWk5S1Ezd3VPL1l6CkdrR1pLd1laR1pBMHRyTXpZcWNYdUYyMkk3ZFd1Ny9xUCtnSlJ1R0tvQitVSXJUUnZ4UllHYUdKRVNzQ0F3RUEKQVFLQ0FnQnE2QXVqQ2tDZjJlNkgyMGlPQ0hLRFdVaHpKTFhvZ3MvQ2VwUW5GUzk2djFwS2RZeUFyeGplVDExMApER3pybTlxTW9ieWNIMjA3OVRlSnRNYVk3Wmluc0ZHc3pmZFZPTk42b0l3QWdiT0g5M3NncGFXM25EQTdVL0VwCjVhRm1ZaXlHMlF3Nms1SlJZYWZXZ2FhQ1NKV1NuUTgyaUtlSXBYNnc2T3JYem9YK2oxNVV2NTVJUGVxTndtY3cKM0t6ZUVkUnpWR3AveUJPNW4zSisyZVl5NE5jbnJsN2xpYUFybGlYdUJWRVRaaXViSXZtTzBXUnJ0MS8xWFpiRQpwVjYyMUg1K2EydjhqcjRxUDlqMHlSc0NCSEtDb0dVeGZMVnFDKzlRQlIrYzh4SHhTN1VKRkRQSGQzQU9zc3NNCmdTSnVXSjRKK2szRHR0a2FmWkJQUDcyRkp6VjhrRHVEcDNSZnlzb1Q0eFBRUGtreFIwK1MvY2pYUXJjcnhzY2gKRHM1SFdwM2ZxTDh2dEtuRlJ2WlMrbm1HaVVYMmdLMXpnUmJLYWdvR3p6Zzl1V1gvaTl0Zk1hZmJyQWdMV21YcwpKblhBRFU1T2duRGdTN01TU1U2Rm5pdlZJSkdKNWJBNFNBRXA1OVdSRi92ckFlcWRDNXVlUkJtcVdtMzNYck1aCkRKUnNZb0hYU041dHlXdGFqaWNpMU1YclZnUkorWlh0cEhTVFBGUUdiSzEwWmZiYzdoQjZGMExXSGpsbngyaTYKcmtRUjVDZ2NxemRndkFCeU5kSmg2bnIyWitoVjFGeUNTbFgvWFZnMllUN1ZpQzdYNTJ4cWluTzNuN0dJOHFOQwpJL0VyZkNmS0pMUGJhTVdmQ2E0TEozaE9uT0d1SU9WdTZCV2VRRHNHU0JFSndjWFdzUUtDQVFFQS90QXp4R01tCjRsT2FSNGFaRkxQRzE0bjFyQUMrbU1OTU9pRzJJeXczSWdsRGVUZkI1S05oampBUWh2NHAyVnJraU9RZDQyOTQKWHBqdnBuSGw1MnZzelpwWlFKeW1zSGFLYjdJNHdCcndaUHBTSTUzd0hYUnpLNTFNQWZpdDhsQlJSWXJ0cWFhLwp6TlRISHdtcTJ0YlloY0krN0dab1NLcndReGUxWUQra1ZCK2REeXNkSkNhQ2VxaDdyQzYrUWdDRG1HeUxMejN2ClZDUWltS0gxaWpCNVJOTDJqM1BEVkwvUVQyYkZtb0FZWlI5ZmxuUUFDV1hPWERjRlpYSzY3SGdmQW13OXR6Z2kKaDU5REllaDZJWW5nWmJ1NzJnNmVyRHo3YVA3T3FrMVhzMHVrdDVVU3p6ZEZON1dXZWIraXdCdTNvZ3g2NFYxQQpCbWJ1bWtzNkN6ZXJsd0tDQVFFQXdkdHp6bkh4dnBML2t6NmgvbjAwVlJndm9vUU9oWTZGbVdDTWdXTWcwNWRECmlMdU9WbmRBazRObXVyVkRScXBYaFNYRFZqbFU4ZnBGSlJNNFRYRFpQaHJEUXlUNEp4L2xDeUxNYkd3RmtPb1cKYXNNZUdoVGdQKzNjc25BVVhCb0FDSmFFbENad3RBaEtYdXI4YnhQdlJMaHAxc2FHTjRjVDR4bVBGb2pOMlB5RQorRXFBWHRrNlhOMCtZVy9oZGFpZkdOL1BwRGFMcEF1MTJzVGFIR1NhOVBDVkNTbzdweHMveG9Tc0xTYW0rL3Y4Cm1lV1ZNVVZaT0RTbjA3a01rbXh3SWtNTXV5eVF5V0VzWWJMM3N1bFphYVBFMU1Ndnk3bUFKR0ppYklBekdERWkKYmJMSk9pcGc1TUMySHZTNFZYeFBGZGVTK3BWUTc4QllvRnBVajhySmpRS0NBUUVBM0RpVWhPWXNkTzVNS0FUcgp5RGlYWVRDYVlrMUNiRVJkWE9CRnlhQXRCZjE3a3dmZFN2enBFem4zRHJRYTl2N1hCSGdpWEsyNkdnZVRGd2JZCjYya2EvNWFtREhGV25xdlVlVFJPVjdqd2lsVE5LSHNYU2wyYUs5ZUdHUzRUSjVqQ3BKZXRUeklPRWJqVFhyKzgKS2VZRXU1VmxUR28xTnBpRmpYYXdDcjcyQnI1THZ4QkQzenBwQ2hrU3lYeWNjZTUvelB3Q1RwSDRoWCsxWnJTUwp3UnVqc3hlZ2Y5cE10cklRRm85N3VFdDh5ZWlUZERSTTA5Sm94c25Hb0NiSDVoYnF0ZTFXYVVMYWxOdlA2VDVDClR6b1o5ZEtLUjZyYTk0Qzh4OEZ3V3o3OHpMaFRZMVl6SzJOWkx3eUJRRGVmTU9qRGpBbTlLWWl1RE5wbzNIQ24KZVlwamdRS0NBUUVBdGxJMEgzU1EzU0NabUIxdTg4OURtY2lPZkhWZ3h3R2M2dnlnQ014M1FpbGdqY2QvL2hoWQpOcVI5eUpuVDlURWQ4UTdzSVRyNGhrQlFLYWRpNjRwMzl1M3F1VXFhelFrMVBIejA3Ly9FV0YrZ3g3Wk1xRkQ3Cis0UTFiZWoxYlEzUy9FQzczaTR0RDFWQXhQYVNoZEdrMWVmdk90MHB2QzJoYVpSUE8rMWNWSGhpZ3JabTkwMnMKazB4TmNBeHVhbDgxaW9wc1dsQW1reG1rWm1WL2tQYVp1a1pPbFBrUWM0Q3dRWC9rQXU3NFc4UEo5ZCt6cWt4RAp0aFhueGJ1amRFN2lRNGIyQVUvUHVHWlkvR1g2aWx6bkIvRExuU01aMzZ2T05lb0dFVytkSG1LUHM4WlRkUTRJClpQeE9ETjB5Uk13T0FVZm5aeDlwcUtNcGQxNmRhME5ZdlFLQ0FRRUFqU29QeWdTY1pCeUNLZ09GaDNiQ3U2czUKZGdsRUl3d3gvWWFuMEVkb0wySkpzVUl1c3JrelFHZGpXODFoSmJFY09vMmRqc24wK3FUbUlBMEdnZmRXTTZFUAo0VUFYUzRSZFZObWR5K0JyNXdiUlQ3YmE5V1JVS0JFY0xGZzFsUEx2SXVsTkNJZ3k0cHlKTy9ENEVqeWNIK1E3CkQrOUVIOVRpTXFXTytCOENrU3pIMFUxN04vR0Z2RHFXU25FWDNXWVZDMS92aG1DWXJQZENYMTVCY3BLV0NkdjAKbW1IaEtJZGl3OWdIaUlPS2daWHVRanZRUW1OT0p5R0ZlMysvT2UycUJlbTVVQ0lhZGhTVFNTcDZ2dzdNZm9teApOWnNCVkg1cS9oV3Q1ZlgyV3ZoZFBIZlJFYTVFYnlUbThPZGNJc0RoN1VpVUJ0dU8rYkFDM2VBRWZIcUZlUT09Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==")
	signKey, _ := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	publicKey, _ := base64.StdEncoding.DecodeString("LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUF3UFZtYUZ2MmlnSW8vVkhOTjl3KwpoWUZDdHNFby8rZ0gwaWtpOEh3RlBybnBpVDN2LzBySm1FM3JXakFYTHBTVXk0aHZHV1ZTNU8zOVhxSDJ3d3NxCk5uYlNySUd6T3NKWGx4NVNOdy9BT0VGV012YzdXKy9CT0VjTEZRUUt2ckk1OWlQNmplL3ViK2RjUXo0NzVsU0UKNkpHbFczRlh1dkVJelQrYlliNEZYYTVrRDJJMGllRUdzaEpXZGY2YUpCMzN4ZmViMjRaNTBJOGZ1cGxCUC9meQozSnFJNDlEeGppTm5XSS9TWC9iMzdmRGppd3BqRFVSMUEzNXFJMndUL3hCUzd1akR4bnpzajk3YmhwUVZ5eXl3CmJET3AvNWNqZ3RuUGZqRHRXL2hndExNREVMOFNaakp6Qld1ek92VHNZZE1iRStZTmZwMktOaHdHSUI4NWZza1IKWFFMUll6U2pwMHRrZ3dUNDFidTB2R2F4UC9Qa1NQMm1yb1IzVkw5dWJ2YlZrNWNFRlJyc0hVZHlNVjFYalhhKwpidEo5bDV2WVp4NDF4OE1teWxGOExQc0V3cFpITVU1dFd3SXgxZVdMQm9JWlYyM1E2SVAxVXVNTnRmRjJPR2RBCllwWmJFM1VaMXp6K0lJeHFoMjhSN2xnVmcyaExua0Q0T3M3L0ZMZG5PbEMraENsZXJZc1NtcmJuc1NRUExWK2sKRHlYaUN4M3NRQmRHSVV6cmI4MnhWK05IaTBHQ1REVVBMZTI4YXo1L2lNdjJkdElPNTUyaTVKM3hYVTViU29iRQpLdktyN09GQmdMUTUzY0FVZWk5S1Ezd3VPL1l6R2tHWkt3WVpHWkEwdHJNellxY1h1RjIySTdkV3U3L3FQK2dKClJ1R0tvQitVSXJUUnZ4UllHYUdKRVNzQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=")
	validateKey, _ := jwt.ParseRSAPublicKeyFromPEM(publicKey)

	serviceConfig := common.ServerConfig{
		SignedInPrivateKey: signKey,
		ValidatePublicKey:  validateKey,
		SigningMethod:      jwt.SigningMethodRS256,
		TokenExpireHour:    100,
	}

	tokenProcessor = TokenProcessorImpl{
		UserService:           userService,
		OperatorPolicyService: operatorPolicyService,
		Service:               ser,
		PolicyService:         policyService,
		RoleService:           roleService,
		PermissionService:     permissionService,
		ServerConfig:          serviceConfig,
		Token:                 jwt.New(serviceConfig.SigningMethod),
	}
}

// Test constructor
func TestGetTokenProcessorInstance(t *testing.T) {
	GetTokenProcessorInstance()
}

// Test Generate token error
func TestGenerate_Error(t *testing.T) {
	_, err := tokenProcessor.Generate(
		common.AuthOperator,
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Gruop id is text.")
		t.FailNow()
	}

	_, err = tokenProcessor.Generate(
		common.AuthOperator,
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Generate operator token")
		t.FailNow()
	}

	_, err = tokenProcessor.Generate(
		common.AuthUser,
		model.TokenRequest{
			Password: "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Generate user token")
		t.FailNow()
	}

	_, err = tokenProcessor.Generate(
		"",
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Generate token")
		t.FailNow()
	}

	_, err = tokenProcessor.Generate(
		"none",
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Generate token")
		t.FailNow()
	}
}

// Test Generate token success
func TestGenerate_Success(t *testing.T) {
	_, err := tokenProcessor.Generate(
		common.AuthUser,
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "test",
		},
	)
	if err != nil {
		t.Errorf("Incorrect TestGenerate_Success test. " + err.ToJson())
		t.FailNow()
	}
}

// Test Generate refresh token success
func TestGenerateRefreshToken_Success(t *testing.T) {
	to, _ := tokenProcessor.Generate(
		common.AuthUser,
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "test",
		},
	)

	_, err := tokenProcessor.Generate(
		common.AuthUser,
		model.TokenRequest{
			GrantType:    model.GrantRefreshToken.String(),
			RefreshToken: to.RefreshToken,
		},
	)

	if err != nil {
		t.Errorf("Incorrect TestGenerateRefreshToken_Success test. %s", err.ToJson())
		t.FailNow()
	}
}

// Test verify operator token
func TestVerifyOperatorToken_Error(t *testing.T) {
	token, _ := tokenProcessor.Generate(
		common.AuthUser,
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "test",
		},
	)

	_, err := tokenProcessor.VerifyOperatorToken(token.Token)
	if err == nil {
		t.Errorf("Incorrect TestVerifyOperatorToken_Error test.")
		t.FailNow()
	}
}

// Test verify operator token
func TestVerifyOperatorToken_Success(t *testing.T) {
	token, _ := tokenProcessor.Generate(
		common.AuthOperator,
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "test",
		},
	)
	_, err := tokenProcessor.VerifyOperatorToken(fmt.Sprintf("Bearer %s", token.Token))
	if err != nil {
		t.Errorf("Incorrect TestVerifyOperatorToken_Success test." + err.ToJson())
		t.FailNow()
	}
}

// Test verify user token
func TestVerifyUserToken_Error(t *testing.T) {
	_, err := tokenProcessor.VerifyUserToken("test_token", "test_role", "test_permission", "")
	if err == nil {
		t.Errorf("Incorrect TestVerifyUserToken_Error test.")
		t.FailNow()
	}
}

// Test verify user token
func TestVerifyUserToken_Success(t *testing.T) {
	token, _ := tokenProcessor.Generate(
		common.AuthUser,
		model.TokenRequest{
			GrantType: "password",
			Email:     "test@gmail.com",
			Password:  "test",
		},
	)
	_, err := tokenProcessor.VerifyUserToken("Bearer "+token.Token, "test_role", "test_permission", "")
	if err != nil {
		t.Errorf("Incorrect TestVerifyUserToken_Success test." + err.ToJson())
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
	user := entity.User{
		Username: "test",
		Email:    email,
		Password: "$2a$10$oqIkJ.fryDacNpVwOkONYe4spwRj9ynuh/YoYOifAlzNa5uWVX/aa",
	}
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindByGroupUuid(groupUuid string) ([]*entity.User, error) {
	var users []*entity.User
	return users, nil
}

func (uri StubUserRepositoryImpl) FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, error) {
	user := entity.User{
		InternalId: "",
		Username:   "test",
		Email:      email,
		Password:   "$2a$10$oqIkJ.fryDacNpVwOkONYe4spwRj9ynuh/YoYOifAlzNa5uWVX/aa",
	}
	operatorPolicy := entity.OperatorPolicy{
		RoleUuid: uuid.New(),
		UserUuid: uuid.New(),
	}
	uwo := model.UserWithOperatorPolicy{
		User:           user,
		OperatorPolicy: operatorPolicy,
		Role:           entity.Role{Name: common.OperatorRole},
	}
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

func (uri StubUserRepositoryImpl) FindUserServicesByUserUuid(userUuid string) ([]*entity.UserService, error) {
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

// Less than stub struct
// OperatorPolicy repository
type StubOperatorPolicyRepositoryImpl struct {
	Connection *gorm.DB
}

func (opr StubOperatorPolicyRepositoryImpl) FindAll() ([]*entity.OperatorPolicy, error) {
	var entities []*entity.OperatorPolicy
	return entities, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindByUserUuid(userUuid string) ([]*entity.OperatorPolicy, error) {
	return []*entity.OperatorPolicy{{RoleUuid: uuid.New(), UserUuid: uuid.New()}}, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindByUserUuidAndRoleUuid(userUuid string, roleUuid string) (*entity.OperatorPolicy, error) {
	operatorMemberRole := entity.OperatorPolicy{
		RoleUuid: uuid.New(),
		UserUuid: uuid.New(),
	}
	return &operatorMemberRole, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindRoleNameByUserUuid(userUuid string) ([]string, error) {
	return []string{""}, nil
}

func (opr StubOperatorPolicyRepositoryImpl) Save(entity entity.OperatorPolicy) (*entity.OperatorPolicy, error) {
	return &entity, nil
}

// Less than stub struct
// Role repository
type StubRoleRepositoryImpl struct {
	Connection *gorm.DB
}

func (rri StubRoleRepositoryImpl) FindAll() ([]*entity.Role, error) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Role, error) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindByUuid(uuid string) (*entity.Role, error) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByName(name string) (*entity.Role, error) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByNames(names []string) ([]entity.Role, error) {
	var roles []entity.Role
	roles = append(roles, entity.Role{InternalId: "", Name: "test_role"})
	roles = append(roles, entity.Role{InternalId: "", Name: "test_role"})
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindByGroupUuid(groupUuid string) ([]*entity.Role, error) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindNameByUuid(uuid string) *string {
	role, err := rri.FindByUuid(uuid)
	if err != nil {
		return nil
	}
	return &role.Name
}

func (rri StubRoleRepositoryImpl) Save(role entity.Role) (*entity.Role, error) {
	return &role, nil
}

func (rri StubRoleRepositoryImpl) SaveWithRelationalData(groupUuid string, role entity.Role) (*entity.Role, error) {
	return &role, nil
}

// Less than stub struct
// Service repository
type StubServiceRepositoryImpl struct {
	Connection *gorm.DB
}

func (sri StubServiceRepositoryImpl) FindAll() ([]*entity.Service, error) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, error) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) FindByUuid(uuid string) (*entity.Service, error) {
	var service entity.Service
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindByName(name string) (*entity.Service, error) {
	var service entity.Service
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindBySecret(secret string) (*entity.Service, error) {
	service := entity.Service{Name: "test"}
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindNameByUuid(uuid string) *string {
	service, err := sri.FindByUuid(uuid)
	if err != nil {
		return nil
	}
	return &service.Name
}

func (sri StubServiceRepositoryImpl) FindServicesByUserUuid(userUuid string) ([]*entity.Service, error) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, error) {
	return &service, nil
}

func (sri StubServiceRepositoryImpl) SaveWithRelationalData(service entity.Service, roles []entity.Role, permissions []entity.Permission) (*entity.Service, error) {
	return &service, nil
}

func (sri StubServiceRepositoryImpl) Update(service entity.Service) (*entity.Service, error) {
	return &service, nil
}

// Less than stub struct
// Permission repository
type StubPermissionRepositoryImpl struct {
	Connection *gorm.DB
}

func (pri StubPermissionRepositoryImpl) FindAll() ([]*entity.Permission, error) {
	permissions := []*entity.Permission{{InternalId: "", Name: "test_permission"}}
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindByUuid(uuid string) (*entity.Permission, error) {
	var permission entity.Permission
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByName(name string) (*entity.Permission, error) {
	permission := entity.Permission{InternalId: "", Name: "test_permission"}
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByNames(names []string) ([]entity.Permission, error) {
	var permissions []entity.Permission
	permissions = append(permissions, entity.Permission{InternalId: "", Name: "test_permission"})
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindByGroupUuid(groupUuid string) ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindNameByUuid(uuid string) *string {
	return nil
}

func (pri StubPermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, error) {
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) SaveWithRelationalData(groupUuid string, permission entity.Permission) (*entity.Permission, error) {
	return &permission, nil
}

// Less than stub struct
// Group repository
type StubGroupRepositoryImpl struct {
	Connection *gorm.DB
}

func (gr StubGroupRepositoryImpl) FindAll() ([]*entity.Group, error) {
	return []*entity.Group{}, nil
}

func (gr StubGroupRepositoryImpl) FindByUuid(uuid string) (*entity.Group, error) {
	var group entity.Group
	return &group, nil
}

func (gr StubGroupRepositoryImpl) FindByName(name string) (*entity.Group, error) {
	var group *entity.Group
	return group, nil
}

func (gr StubGroupRepositoryImpl) FindByUserUuid(userUuid string) ([]*entity.Group, error) {
	var groups []*entity.Group
	return groups, nil
}

func (gr StubGroupRepositoryImpl) FindByServiceUuid(serviceUuid string) ([]*entity.Group, error) {
	var groups []*entity.Group
	return groups, nil
}

func (gr StubGroupRepositoryImpl) FindGroupWithUserWithPolicyGroupsByUserUuid(userUuid string) ([]*model.GroupWithUserGroupWithPolicy, error) {
	uid := uuid.MustParse("20aabf03-7f3f-479c-91d2-909c844de26c")
	var policies []*model.GroupWithUserGroupWithPolicy
	policies = append(policies, &model.GroupWithUserGroupWithPolicy{Policy: entity.Policy{UserGroupUuid: uid, RoleUuid: uid, PermissionUuid: uid, ServiceUuid: uid}})
	return policies, nil
}

func (gr StubGroupRepositoryImpl) FindGroupWithPolicyByUserUuidAndGroupUuid(userUuid string, groupUuid string) (*model.GroupWithUserGroupWithPolicy, error) {
	var groupWithUserGroupWithPolicy model.GroupWithUserGroupWithPolicy
	return &groupWithUserGroupWithPolicy, nil
}

func (gr StubGroupRepositoryImpl) SaveWithRelationalData(
	group entity.Group,
	serviceGroup entity.ServiceGroup,
	userGroup entity.UserGroup,
	groupPermission entity.GroupPermission,
	groupRole entity.GroupRole,
	policy entity.Policy) (*entity.Group, error) {

	return &group, nil
}

// Less than stub struct
// Policy repository
type StubPolicyRepositoryImpl struct {
	Connection *gorm.DB
}

func (pri StubPolicyRepositoryImpl) FindAll() ([]*entity.Policy, error) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Policy, error) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindByRoleUuid(roleUuid string) ([]*entity.Policy, error) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindByUuid(uuid string) (entity.Policy, error) {
	var policy entity.Policy
	return policy, nil
}

func (pri StubPolicyRepositoryImpl) FindPolicyOfUserGroupByUserUuidAndGroupUuid(userUuid string, groupUuid string) (model.UserPolicyOnGroupResponse, error) {
	var policy model.UserPolicyOnGroupResponse
	return policy, nil
}

func (pri StubPolicyRepositoryImpl) FindPolicyOfUserServiceByUserUuidAndServiceUuid(userUuid string) ([]model.UserPolicyOnServiceResponse, error) {
	var policy []model.UserPolicyOnServiceResponse
	return policy, nil
}

func (pri StubPolicyRepositoryImpl) Update(policy entity.Policy) (*entity.Policy, error) {
	return &policy, nil
}

// Less than stub struct
// Policy repository
type StubEtcdlClient struct {
}

func (e StubEtcdlClient) SetUserPolicy(userUuid string, policy []structure.UserPolicy) {
}

func (e StubEtcdlClient) SetPermission(permissionUuid string, permission structure.Permission) {
}

func (e StubEtcdlClient) SetRole(roleUuid string, role structure.Role) {
}

func (e StubEtcdlClient) SetService(serviceUuid string, service structure.Service) {
}

func (e StubEtcdlClient) SetUserService(userUuid string, userServices []structure.UserService) {
}

func (e StubEtcdlClient) SetUserGroup(userUuid string, userGroups []structure.UserGroup) {
}

func (e StubEtcdlClient) GetUserPolicy(userUuid string) []structure.UserPolicy {
	uid := "6311b383-4587-469d-8aa1-8103142d33e8"
	return []structure.UserPolicy{{ServiceUuid: uid, GroupUuid: uid, RoleName: "test_role", PermissionName: "test_permission"}}
}

func (e StubEtcdlClient) GetPermission(permissionUuid string) *structure.Permission {
	var permission structure.Permission
	return &permission
}

func (e StubEtcdlClient) GetRole(roleUuid string) *structure.Role {
	var role structure.Role
	return &role
}

func (e StubEtcdlClient) GetService(serviceUuid string) *structure.Service {
	var service structure.Service
	return &service
}

func (e StubEtcdlClient) GetUserService(userUuid string) []structure.UserService {
	var userServices []structure.UserService
	return userServices
}

func (e StubEtcdlClient) GetUserGroup(userUuid string) []structure.UserGroup {
	var userGroups []structure.UserGroup
	return userGroups
}

func (e StubEtcdlClient) DeleteUserPolicy(userUuid string) {
}
