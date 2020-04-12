package middleware

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
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
	ctx.InitContext()
	ctx.SetUserId(1)
	ctx.SetServiceId(1)
	ctx.SetUserUuid(uuid.New())
	ctx.SetApiKey("test")

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

	tokenProcessor = TokenProcessorImpl{
		UserService:           userService,
		OperatorPolicyService: operatorPolicyService,
		Service:               ser,
		PolicyService:         policyService,
		RoleService:           roleService,
		PermissionService:     permissionService,
		ServerConfig:          serviceConfig,
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
		"test",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Gruop id is text.")
		t.FailNow()
	}

	_, err = tokenProcessor.Generate(
		common.AuthOperator,
		"",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Generate operator token")
		t.FailNow()
	}

	_, err = tokenProcessor.Generate(
		common.AuthUser,
		"",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Generate user token")
		t.FailNow()
	}

	_, err = tokenProcessor.Generate(
		"",
		"",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "testa",
		},
	)
	if err == nil {
		t.Errorf("Incorrect TestGenerate_Error test. Generate token")
		t.FailNow()
	}

	_, err = tokenProcessor.Generate(
		"none",
		"",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "testa",
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
		common.AuthOperator,
		"",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "test",
		},
	)
	if err != nil {
		t.Errorf("Incorrect TestGenerate_Success test.")
		t.FailNow()
	}
}

// Test parse token
func TestParseToken_Error(t *testing.T) {
	_, result := tokenProcessor.ParseToken("InvalidToken")
	if result {
		t.Errorf("Incorrect TestParseToken_Error test.")
		t.FailNow()
	}
}

// Test parse token
func TestParseToken_Success(t *testing.T) {
	token, _ := tokenProcessor.Generate(
		common.AuthUser,
		"",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "test",
		},
	)

	_, result := tokenProcessor.ParseToken(token)
	if !result {
		t.Errorf("Incorrect TestParseToken_Success test.")
		t.FailNow()
	}
}

// Test verify operator token
func TestVerifyOperatorToken_Error(t *testing.T) {
	token, _ := tokenProcessor.Generate(
		common.AuthOperator,
		"",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "test",
		},
	)

	_, err := tokenProcessor.VerifyOperatorToken(token)
	if err == nil {
		t.Errorf("Incorrect TestVerifyOperatorToken_Error test.")
		t.FailNow()
	}
}

// Test verify operator token
func TestVerifyOperatorToken_Success(t *testing.T) {
	token, _ := tokenProcessor.Generate(
		common.AuthOperator,
		"",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "test",
		},
	)
	token = "Bearer " + token

	_, err := tokenProcessor.VerifyOperatorToken(token)
	if err != nil {
		t.Errorf("Incorrect TestVerifyOperatorToken_Success test." + err.ToJson())
		t.FailNow()
	}
}

// Test verify user token
func TestVerifyUserToken_Error(t *testing.T) {
	_, err := tokenProcessor.VerifyUserToken("test_token", []string{"test_role"}, "test_permission")
	if err == nil {
		t.Errorf("Incorrect TestVerifyUserToken_Error test." + err.ToJson())
		t.FailNow()
	}
}

// Test verify user token
func TestVerifyUserToken_Success(t *testing.T) {
	token, _ := tokenProcessor.Generate(
		common.AuthUser,
		"1",
		entity.User{
			Username: "test",
			Email:    "test@gmail.com",
			Password: "test",
		},
	)
	token = "Bearer " + token

	_, err := tokenProcessor.VerifyUserToken(token, []string{"test_role"}, "test_permission")
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

func (uri StubUserRepositoryImpl) FindById(id int) (*entity.User, *model.ErrorResBody) {
	var user entity.User
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindByEmail(email string) (*entity.User, *model.ErrorResBody) {
	user := entity.User{
		Username: "test",
		Email:    email,
		Password: "$2a$10$oqIkJ.fryDacNpVwOkONYe4spwRj9ynuh/YoYOifAlzNa5uWVX/aa",
	}
	return &user, nil
}

func (uri StubUserRepositoryImpl) FindWithOperatorPolicyByEmail(email string) (*model.UserWithOperatorPolicy, *model.ErrorResBody) {
	user := entity.User{
		Id:       1,
		Username: "test",
		Email:    email,
		Password: "$2a$10$oqIkJ.fryDacNpVwOkONYe4spwRj9ynuh/YoYOifAlzNa5uWVX/aa",
	}
	operatorPolicy := entity.OperatorPolicy{
		RoleId: 1,
		UserId: 1,
	}
	uwo := model.UserWithOperatorPolicy{
		user,
		operatorPolicy,
	}
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
	userService := entity.UserService{Id: 1}
	return &userService, nil
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

// Less than stub struct
// OperatorPolicy repository
type StubOperatorPolicyRepositoryImpl struct {
	Connection *gorm.DB
}

func (opr StubOperatorPolicyRepositoryImpl) FindAll() ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	var entities []*entity.OperatorPolicy
	return entities, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindByUserId(userId int) ([]*entity.OperatorPolicy, *model.ErrorResBody) {
	var entities []*entity.OperatorPolicy
	return entities, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindByUserIdAndRoleId(userId int, roleId int) (*entity.OperatorPolicy, *model.ErrorResBody) {
	operatorMemberRole := entity.OperatorPolicy{
		RoleId: 1,
		UserId: 1,
	}
	return &operatorMemberRole, nil
}

func (opr StubOperatorPolicyRepositoryImpl) FindRoleNameByUserId(userId int) ([]string, *model.ErrorResBody) {
	var names []string
	return names, nil
}

func (opr StubOperatorPolicyRepositoryImpl) Save(entity entity.OperatorPolicy) (*entity.OperatorPolicy, *model.ErrorResBody) {
	return &entity, nil
}

// Less than stub struct
// Role repository
type StubRoleRepositoryImpl struct {
	Connection *gorm.DB
}

func (rri StubRoleRepositoryImpl) FindAll() ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindById(id int) (*entity.Role, *model.ErrorResBody) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByName(name string) (*entity.Role, *model.ErrorResBody) {
	var role entity.Role
	return &role, nil
}

func (rri StubRoleRepositoryImpl) FindByNames(names []string) ([]entity.Role, *model.ErrorResBody) {
	var roles []entity.Role
	roles = append(roles, entity.Role{Id: 1, Name: "test_role"})
	roles = append(roles, entity.Role{Id: 2, Name: "test_role"})
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindByGroupId(groupId int) ([]*entity.Role, *model.ErrorResBody) {
	var roles []*entity.Role
	return roles, nil
}

func (rri StubRoleRepositoryImpl) FindNameById(id int) *string {
	role, _ := rri.FindById(id)
	return &role.Name
}

func (rri StubRoleRepositoryImpl) Save(role entity.Role) (*entity.Role, *model.ErrorResBody) {
	return &role, nil
}

func (rri StubRoleRepositoryImpl) SaveWithRelationalData(groupId int, role entity.Role) (*entity.Role, *model.ErrorResBody) {
	return &role, nil
}

// Less than stub struct
// Service repository
type StubServiceRepositoryImpl struct {
	Connection *gorm.DB
}

func (sri StubServiceRepositoryImpl) FindAll() ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) FindOffSetAndLimit(offset int, limit int) ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) FindById(id int) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindByName(name string) (*entity.Service, *model.ErrorResBody) {
	var service entity.Service
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResBody) {
	service := entity.Service{Name: "test"}
	return &service, nil
}

func (sri StubServiceRepositoryImpl) FindNameById(id int) *string {
	service, _ := sri.FindById(id)
	return &service.Name
}

func (sri StubServiceRepositoryImpl) FindNameByApiKey(name string) *string {
	service, _ := sri.FindByName(name)
	return &service.Name
}

func (sri StubServiceRepositoryImpl) FindServicesByUserId(userId int) ([]*entity.Service, *model.ErrorResBody) {
	var services []*entity.Service
	return services, nil
}

func (sri StubServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &service, nil
}

func (sri StubServiceRepositoryImpl) SaveWithRelationalData(service entity.Service, roles []entity.Role, permissions []entity.Permission) (*entity.Service, *model.ErrorResBody) {
	return &service, nil
}

func (sri StubServiceRepositoryImpl) Update(service entity.Service) (*entity.Service, *model.ErrorResBody) {
	return &service, nil
}

// Less than stub struct
// Permission repository
type StubPermissionRepositoryImpl struct {
	Connection *gorm.DB
}

func (pri StubPermissionRepositoryImpl) FindAll() ([]*entity.Permission, *model.ErrorResBody) {
	permissions := []*entity.Permission{{Id: 1, Name: "test_permission"}}
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindById(id int) (*entity.Permission, *model.ErrorResBody) {
	var permission entity.Permission
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByName(name string) (*entity.Permission, *model.ErrorResBody) {
	permission := entity.Permission{Id: 1, Name: "test_permission"}
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) FindByNames(names []string) ([]entity.Permission, *model.ErrorResBody) {
	var permissions []entity.Permission
	permissions = append(permissions, entity.Permission{Id: 1, Name: "test_permission"})
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindByGroupId(groupId int) ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	return permissions, nil
}

func (pri StubPermissionRepositoryImpl) FindNameById(id int) *string {
	permission, _ := pri.FindById(id)
	return &permission.Name
}

func (pri StubPermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	return &permission, nil
}

func (pri StubPermissionRepositoryImpl) SaveWithRelationalData(groupId int, permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	return &permission, nil
}

// Less than stub struct
// Group repository
type StubGroupRepositoryImpl struct {
	Connection *gorm.DB
}

func (gr StubGroupRepositoryImpl) FindAll() ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group
	return groups, nil
}

func (gr StubGroupRepositoryImpl) FindById(id int) (*entity.Group, *model.ErrorResBody) {
	var group entity.Group
	return &group, nil
}

func (gr StubGroupRepositoryImpl) FindByName(name string) (*entity.Group, *model.ErrorResBody) {
	var group *entity.Group
	return group, nil
}

func (gr StubGroupRepositoryImpl) FindGroupsByUserId(userId int) ([]*entity.Group, *model.ErrorResBody) {
	var groups []*entity.Group
	return groups, nil
}

func (gr StubGroupRepositoryImpl) FindGroupWithUserWithPolicyGroupsByUserId(userId int) ([]*model.GroupWithUserGroupWithPolicy, *model.ErrorResBody) {
	var groupWithUserGroupWithPolicies []*model.GroupWithUserGroupWithPolicy
	groupWithUserGroupWithPolicies = append(groupWithUserGroupWithPolicies, &model.GroupWithUserGroupWithPolicy{entity.Group{}, entity.UserGroup{}, entity.Policy{ServiceId: 1}})
	return groupWithUserGroupWithPolicies, nil
}

func (gr StubGroupRepositoryImpl) FindGroupWithPolicyByUserIdAndGroupId(userId int, groupId int) (*model.GroupWithUserGroupWithPolicy, *model.ErrorResBody) {
	group := entity.Group{}
	userGroup := entity.UserGroup{}
	policy := entity.Policy{Id:1}
	groupWithUserGroupWithPolicy := model.GroupWithUserGroupWithPolicy {
		group,
		userGroup,
		policy,
	}
	return &groupWithUserGroupWithPolicy, nil

}

func (gr StubGroupRepositoryImpl) SaveWithRelationalData(
	group entity.Group,
	serviceGroup entity.ServiceGroup,
	userGroup entity.UserGroup,
	groupPermission entity.GroupPermission,
	groupRole entity.GroupRole,
	policy entity.Policy) (*entity.Group, *model.ErrorResBody) {

	return &group, nil
}

// Less than stub struct
// Policy repository
type StubPolicyRepositoryImpl struct {
	Connection *gorm.DB
}

func (pri StubPolicyRepositoryImpl) FindAll() ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindOffSetAndLimit(offsetCnt int, limitCnt int) ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindByRoleId(roleId int) ([]*entity.Policy, *model.ErrorResBody) {
	var policies []*entity.Policy
	return policies, nil
}

func (pri StubPolicyRepositoryImpl) FindById(id int) (entity.Policy, *model.ErrorResBody) {
	policy := entity.Policy{
		RoleId:       1,
		PermissionId: 1,
	}
	return policy, nil
}

func (pri StubPolicyRepositoryImpl) Update(policy entity.Policy) (*entity.Policy, *model.ErrorResBody) {
	return &policy, nil
}
