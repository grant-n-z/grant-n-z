package service

import (
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"go.etcd.io/etcd/clientv3"
	"testing"
	"time"
)

var (
	stubEtcdConnection *clientv3.Client
	updaterService     UpdaterService
)

func init() {
	log.InitLogger("info")
	ctx.InitContext()

	stubEtcdConnection, _ = clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})
	etcdClient := cache.EtcdClientImpl{
		Connection: stubEtcdConnection,
		Ctx:        ctx.GetCtx(),
	}
	updaterService = UpdaterServiceImpl{EtcdClient: etcdClient}
}

// Test constructor
func TestNewUpdaterService(t *testing.T) {
	NewUpdaterService()
}

// Test update policy
func TestUpdatePolicy(t *testing.T) {
	var policies []*entity.Policy
	policies = []*entity.Policy{{Id: 1, Name: "test"}}
	updaterService.UpdatePolicy(policies)
}

// Test update permission
func TestUpdatePermission(t *testing.T) {
	var permissions []*entity.Permission
	permissions = []*entity.Permission{{Id: 1, Name: "test"}}
	updaterService.UpdatePermission(permissions)
}

// Test update role
func TestUpdateRole(t *testing.T) {
	var roles []*entity.Role
	roles = []*entity.Role{{Id: 1, Name: "test"}}
	updaterService.UpdateRole(roles)
}

// Test update service
func TestUpdateService(t *testing.T) {
	var services []*entity.Service
	services = []*entity.Service{{Id: 1, Name: "test"}}
	updaterService.UpdateService(services)
}

// Test update user service
func TestUpdateUserService(t *testing.T) {
	var userServices []*entity.UserService
	userServices = []*entity.UserService{{Id: 1, UserId: 1, ServiceId: 1}}
	updaterService.UpdateUserService(userServices)
}

// Test update user group
func TestUpdateUserGroup(t *testing.T) {
	var userGroups []*entity.UserGroup
	userGroups = []*entity.UserGroup{{Id: 1, UserId: 1, GroupId: 1}}
	updaterService.UpdateUserGroup(userGroups)
}
