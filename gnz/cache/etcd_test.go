package cache

import (
	"context"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var etcdClient EtcdClient

func init() {
	log.InitLogger("info")
	ctx.InitContext()
}

// Setup not connected etdc pattern
func setUpNotConnected() {
	stubConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{},
		DialTimeout: 5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	connection = stubConnection
	c, _ := context.WithTimeout(ctx.GetCtx(), 10*time.Millisecond)
	etcdClient = EtcdClientImpl{
		Connection: connection,
		Ctx: c,
	}
}

// Setup connected etdc, but put is faild pattern
func setUpStubConnected() {
	stubConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:9999"},
		DialTimeout: 5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})

	connection = stubConnection
	c, _ := context.WithTimeout(ctx.GetCtx(), 10*time.Millisecond)
	etcdClient = EtcdClientImpl{
		Connection: connection,
		Ctx: c,
	}
}

// Test constructor
func TestGetEtcdClientInstance(t *testing.T) {
	GetEtcdClientInstance()
}

// This is not connected pattern for PUT
// SetPolicy failed test
func TestSetPolicy_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetPolicy(entity.Policy{Id: 1, Name: "test"}, 0)
}

// SetPermission failed test
func TestSetPermission_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetPermission(entity.Permission{}, 10)
}

// SetRole failed test
func TestSetRole_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetRole(entity.Role{}, 10)
}

// SetService failed test
func TestSetService_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetService(entity.Service{}, 10)
}

// SetUserService failed test
func TestSetUserService_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetUserService(1, []entity.UserService{{}}, 10)
}

// This is connected pattern for PUT
// SetPolicy failed test
func TestSetPolicy_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetPolicy(entity.Policy{Id: 1, Name: "test"}, 0)
}

// SetPermission failed test
func TestSetPermission_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetPermission(entity.Permission{}, 10)
}

// SetRole failed test
func TestSetRole_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetRole(entity.Role{}, 10)
}

// SetService failed test
func TestSetService_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetService(entity.Service{}, 10)
}

// SetUserService failed test
func TestSetUserService_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetUserService(1, []entity.UserService{{}}, 10)
}

// SetUserGroup failed test
func TestSetUserGroup_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetUserGroup(1, []entity.UserGroup{{}}, 10)
}

// This is not connected pattern for GET
// GetPolicy nil test
func TestGetPolicy_NotConnected(t *testing.T) {
	setUpNotConnected()
	policy := etcdClient.GetPolicy("policy")
	if policy != nil {
		t.Errorf("Incorrect TestGetPolicy_Nil test")
		t.FailNow()
	}
}

// GetPolicyByNames nil test
func TestGetPolicyByNames_NotConnected(t *testing.T) {
	setUpNotConnected()
	policy := etcdClient.GetPolicyByNames([]string{"policy"})
	if policy != nil {
		t.Errorf("Incorrect TestGetPolicyByNames_Nil test")
		t.FailNow()
	}
}

// GetPermission nil test
func TestGetPermission_NotConnected(t *testing.T) {
	setUpNotConnected()
	permission := etcdClient.GetPermission("permission")
	if permission != nil {
		t.Errorf("Incorrect TestGetPermission_Nil test")
		t.FailNow()
	}
}

// GetPermissionByNames nil test
func TestGetPermissionByNames_NotConnected(t *testing.T) {
	setUpNotConnected()
	permission := etcdClient.GetPermissionByNames([]string{"permission"})
	if permission != nil {
		t.Errorf("Incorrect TestGetPermissionByNames_Nil test")
		t.FailNow()
	}
}

// GetRole nil test
func TestGetRole_NotConnected(t *testing.T) {
	setUpNotConnected()
	role := etcdClient.GetRole("role")
	if role != nil {
		t.Errorf("Incorrect TestGetRole_Nil test")
		t.FailNow()
	}
}

// GetRoleByNames nil test
func TestGetRoleByNames_NotConnected(t *testing.T) {
	setUpNotConnected()
	role := etcdClient.GetRoleByNames([]string{"role"})
	if role != nil {
		t.Errorf("Incorrect TestGetRoleByNames_Nil test")
		t.FailNow()
	}
}

// GetService nil test
func TestGetService_NotConnected(t *testing.T) {
	setUpNotConnected()
	service := etcdClient.GetService("service")
	if service != nil {
		t.Errorf("Incorrect TestGetService_Nil test")
		t.FailNow()
	}
}

// GetServiceByNames nil test
func TestGetServiceByNames_NotConnected(t *testing.T) {
	setUpNotConnected()
	service := etcdClient.GetServiceByNames([]string{"service"})
	if service != nil {
		t.Errorf("Incorrect TestGetServiceByNames_Nil test")
		t.FailNow()
	}
}

// GetUserService nil test
func TestGetUserService_NotConnected(t *testing.T) {
	setUpNotConnected()
	userService := etcdClient.GetUserService(1, 1)
	if userService != nil {
		t.Errorf("Incorrect TestGetUserService_Nil test")
	}
}

// This is connected pattern for GET
// GetPolicy nil test
func TestGetPolicy_Nil(t *testing.T) {
	setUpStubConnected()
	policy := etcdClient.GetPolicy("policy")
	if policy != nil {
		t.Errorf("Incorrect TestGetPolicy_Nil test")
	}
}

// GetPolicyByNames nil test
func TestGetPolicyByNames_Nil(t *testing.T) {
	setUpStubConnected()
	policy := etcdClient.GetPolicyByNames([]string{"policy"})
	if policy != nil {
		t.Errorf("Incorrect TestGetPolicyByNames_Nil test")
		t.FailNow()
	}
}

// GetPermission nil test
func TestGetPermission_Nil(t *testing.T) {
	setUpStubConnected()
	permission := etcdClient.GetPermission("permission")
	if permission != nil {
		t.Errorf("Incorrect TestGetPermission_Nil test")
		t.FailNow()
	}
}

// GetPermissionByNames nil test
func TestGetPermissionByNames_Nil(t *testing.T) {
	setUpStubConnected()
	permission := etcdClient.GetPermissionByNames([]string{"permission"})
	if permission != nil {
		t.Errorf("Incorrect TestGetPermissionByNames_Nil test")
		t.FailNow()
	}
}

// GetRole nil test
func TestGetRole_Nil(t *testing.T) {
	setUpStubConnected()
	role := etcdClient.GetRole("role")
	if role != nil {
		t.Errorf("Incorrect TestGetRole_Nil test")
		t.FailNow()
	}
}

// GetRoleByNames nil test
func TestGetRoleByNames_Nil(t *testing.T) {
	setUpStubConnected()
	role := etcdClient.GetRoleByNames([]string{"role"})
	if role != nil {
		t.Errorf("Incorrect TestGetRoleByNames_Nil test")
		t.FailNow()
	}
}

// GetService nil test
func TestGetService_Nil(t *testing.T) {
	setUpStubConnected()
	service := etcdClient.GetService("service")
	if service != nil {
		t.Errorf("Incorrect TestGetService_Nil test")
		t.FailNow()
	}
}

// GetServiceByNames nil test
func TestGetServiceByNames_Nil(t *testing.T) {
	setUpStubConnected()
	service := etcdClient.GetServiceByNames([]string{"service"})
	if service != nil {
		t.Errorf("Incorrect TestGetServiceByNames_Nil test")
		t.FailNow()
	}
}

// GetUserService nil test
func TestGetUserService_Nil(t *testing.T) {
	setUpStubConnected()
	userService := etcdClient.GetUserService(1, 1)
	if userService != nil {
		t.Errorf("Incorrect TestGetUserService_Nil test")
		t.FailNow()
	}
}

// GetUserGroup nil test
func TestGetUserGroup_Nil(t *testing.T) {
	setUpStubConnected()
	userGroup := etcdClient.GetUserGroup(1, 1)
	if userGroup != nil {
		t.Errorf("Incorrect TestGetUserGroup_Nil test")
		t.FailNow()
	}
}
