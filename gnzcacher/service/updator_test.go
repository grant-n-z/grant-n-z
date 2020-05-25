package service

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/google/uuid"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/cache/structure"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var (
	stubEtcdConnection *clientv3.Client
	updaterService     UpdaterService
)

func init() {
	log.InitLogger("info")

	stubEtcdConnection, _ = clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})
	etcdClient := cache.EtcdClientImpl{
		Connection: stubEtcdConnection,
	}
	updaterService = UpdaterServiceImpl{EtcdClient: etcdClient}
}

// Test constructor
func TestNewUpdaterService(t *testing.T) {
	NewUpdaterService()
}

// Test update policy
func TestUpdatePolicy(t *testing.T) {
	policies := make(map[string][]structure.UserPolicy)
	policies["test"] = []structure.UserPolicy{{GroupUuid: uuid.New().String()}}
	updaterService.UpdatePolicy(policies)
}

// Test update permission
func TestUpdatePermission(t *testing.T) {
	var permissions []structure.Permission
	permissions = []structure.Permission{{Name: "test"}}
	updaterService.UpdatePermission(permissions)
}

// Test update role
func TestUpdateRole(t *testing.T) {
	var roles []structure.Role
	roles = []structure.Role{{Name: "test"}}
	updaterService.UpdateRole(roles)
}

// Test update service
func TestUpdateService(t *testing.T) {
	var services []structure.Service
	services = []structure.Service{{Name: "test"}}
	updaterService.UpdateService(services)
}

// Test update user service
func TestUpdateUserService(t *testing.T) {
	services := make(map[string][]structure.UserService)
	services["test"] = []structure.UserService{{ServiceName: "service"}}
	updaterService.UpdateUserService(services)
}

// Test update user group
func TestUpdateUserGroup(t *testing.T) {
	groups := make(map[string][]structure.UserGroup)
	groups["test"] = []structure.UserGroup{{GroupName: "group"}}
	updaterService.UpdateUserGroup(groups)
}
