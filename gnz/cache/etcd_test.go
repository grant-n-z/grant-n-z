package cache

import (
	"context"
	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/tomoyane/grant-n-z/gnz/log"
)

var etcdClient EtcdClient

func init() {
	log.InitLogger("info")
}

// Setup not connected etdc pattern
func setUpNotConnected() {
	c, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	etcdClient = EtcdClientImpl{
		Connection: nil,
		Ctx:        c,
	}
}

// Setup connected etdc, but put is faild pattern
func setUpStubConnected() {
	stubConnection, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:9999"},
	})

	connection = stubConnection
	c, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	etcdClient = EtcdClientImpl{
		Connection: connection,
		Ctx:        c,
	}
}

// Test constructor
func TestGetEtcdClientInstance(t *testing.T) {
	GetEtcdClientInstance()
}

// This is not connected pattern for PUT
// SetUserPolicy failed test
func TestSetPolicy_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetUserPolicy(uuid.New().String(), []structure.UserPolicy{{RoleName: "test"}})
}

// SetPermission failed test
func TestSetPermission_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetPermission(uuid.New().String(), structure.Permission{Name: "test"})
}

// SetRole failed test
func TestSetRole_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetRole(uuid.New().String(), structure.Role{Name: "test"})
}

// SetService failed test
func TestSetService_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetService(uuid.New().String(), structure.Service{Name: "test"})
}

// SetUserService failed test
func TestSetUserService_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetUserService(uuid.New().String(), []structure.UserService{{ServiceName: "test"}})
}

// SetUserService failed test
func TestSetUserGroup_NotConnected(t *testing.T) {
	setUpNotConnected()
	etcdClient.SetUserGroup(uuid.New().String(), []structure.UserGroup{{GroupName: "test"}})
}

// This is connected pattern for PUT
// SetUserPolicy failed test
func TestSetPolicy_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetUserPolicy(uuid.New().String(), []structure.UserPolicy{{RoleName: "test"}})
}

// SetPermission failed test
func TestSetPermission_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetPermission(uuid.New().String(), structure.Permission{Name: "test"})
}

// SetRole failed test
func TestSetRole_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetRole(uuid.New().String(), structure.Role{Name: "test"})
}

// SetService failed test
func TestSetService_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetService(uuid.New().String(), structure.Service{Name: "test"})
}

// SetUserService failed test
func TestSetUserService_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetUserService(uuid.New().String(), []structure.UserService{{ServiceName: "test"}})
}

// SetUserGroup failed test
func TestSetUserGroup_FailedPut(t *testing.T) {
	setUpStubConnected()
	etcdClient.SetUserGroup(uuid.New().String(), []structure.UserGroup{{GroupName: "test"}})
}

// This is not connected pattern for GET
// GetUserPolicy nil test
func TestGetPolicy_NotConnected(t *testing.T) {
	setUpNotConnected()
	policy := etcdClient.GetUserPolicy("policy")
	if policy != nil {
		t.Errorf("Incorrect TestGetPolicy_Nil test")
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

// GetRole nil test
func TestGetRole_NotConnected(t *testing.T) {
	setUpNotConnected()
	role := etcdClient.GetRole("role")
	if role != nil {
		t.Errorf("Incorrect TestGetRole_Nil test")
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

// GetUserService nil test
func TestGetUserService_NotConnected(t *testing.T) {
	setUpNotConnected()
	userService := etcdClient.GetUserService(uuid.New().String())
	if userService != nil {
		t.Errorf("Incorrect TestGetUserService_Nil test")
	}
}

// GetUserGroup nil test
func TestGetUserGroup_NotConnected(t *testing.T) {
	setUpNotConnected()
	userGroup := etcdClient.GetUserGroup(uuid.New().String())
	if userGroup != nil {
		t.Errorf("Incorrect TestGetUserGroup_NotConnected test")
	}
}

// This is connected pattern for GET
// GetUserPolicy nil test
func TestGetPolicy_Nil(t *testing.T) {
	setUpStubConnected()
	policy := etcdClient.GetUserPolicy("policy")
	if policy != nil {
		t.Errorf("Incorrect TestGetPolicy_Nil test")
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

// GetRole nil test
func TestGetRole_Nil(t *testing.T) {
	setUpStubConnected()
	role := etcdClient.GetRole("role")
	if role != nil {
		t.Errorf("Incorrect TestGetRole_Nil test")
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

// GetUserService nil test
func TestGetUserService_Nil(t *testing.T) {
	setUpStubConnected()
	userService := etcdClient.GetUserService(uuid.New().String())
	if userService != nil {
		t.Errorf("Incorrect TestGetUserService_Nil test")
		t.FailNow()
	}
}

// GetUserGroup nil test
func TestGetUserGroup_Nil(t *testing.T) {
	setUpStubConnected()
	userGroup := etcdClient.GetUserGroup(uuid.New().String())
	if userGroup != nil {
		t.Errorf("Incorrect TestGetUserGroup_Nil test")
		t.FailNow()
	}
}
