package cache

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var etcdClient EtcdClient

func init() {
	log.InitLogger("info")

	client, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{},
		DialTimeout: 5 * time.Second,
	})
	connection = client
	etcdClient = NewEtcdClient()
}

// SetPolicy failed test
func TestSetPolicy_Failed(t *testing.T) {
	etcdClient.SetPolicy(entity.Policy{}, 10)
}

// SetPermission failed test
func TestSetPermission_Failed(t *testing.T) {
	etcdClient.SetPermission(entity.Permission{}, 10)
}

// SetRole failed test
func TestSetRole_Failed(t *testing.T) {
	etcdClient.SetRole(entity.Role{}, 10)
}

// SetService failed test
func TestSetService_Failed(t *testing.T) {
	etcdClient.SetService(entity.Service{}, 10)
}

// SetUserService failed test
func TestSetUserService_Failed(t *testing.T) {
	etcdClient.SetUserService(1, []entity.UserService{{}}, 10)
}

// GetPolicy nil test
func TestGetPolicy_Nil(t *testing.T) {
	policy := etcdClient.GetPolicy("policy")
	if policy != nil {
		t.Errorf("Incorrect TestGetPolicy_Nil test")
	}
}

// GetPolicyByNames nil test
func TestGetPolicyByNames_Nil(t *testing.T) {
	policy := etcdClient.GetPolicyByNames([]string{"policy"})
	if policy != nil {
		t.Errorf("Incorrect TestGetPolicyByNames_Nil test")
	}
}

// GetPermission nil test
func TestGetPermission_Nil(t *testing.T) {
	permission := etcdClient.GetPermission("permission")
	if permission != nil {
		t.Errorf("Incorrect TestGetPermission_Nil test")
	}
}

// GetPermissionByNames nil test
func TestGetPermissionByNames_Nil(t *testing.T) {
	permission := etcdClient.GetPermissionByNames([]string{"permission"})
	if permission != nil {
		t.Errorf("Incorrect TestGetPermissionByNames_Nil test")
	}
}

// GetRole nil test
func TestGetRole_Nil(t *testing.T) {
	role := etcdClient.GetRole("role")
	if role != nil {
		t.Errorf("Incorrect TestGetRole_Nil test")
	}
}

// GetRoleByNames nil test
func TestGetRoleByNames_Nil(t *testing.T) {
	role := etcdClient.GetRoleByNames([]string{"role"})
	if role != nil {
		t.Errorf("Incorrect TestGetRoleByNames_Nil test")
	}
}

// GetService nil test
func TestGetService_Nil(t *testing.T) {
	service := etcdClient.GetService("service")
	if service != nil {
		t.Errorf("Incorrect TestGetService_Nil test")
	}
}

// GetServiceByNames nil test
func TestGetServiceByNames_Nil(t *testing.T) {
	service := etcdClient.GetServiceByNames([]string{"service"})
	if service != nil {
		t.Errorf("Incorrect TestGetServiceByNames_Nil test")
	}
}

// GetUserService nil test
func TestGetUserService_Nil(t *testing.T) {
	userService := etcdClient.GetUserService(1, 1)
	if userService != nil {
		t.Errorf("Incorrect TestGetUserService_Nil test")
	}
}
