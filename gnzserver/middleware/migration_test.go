package middleware

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var migration Migration

func init() {
	log.InitLogger("info")

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

	roleService := service.RoleServiceImpl{
		EtcdClient:     cache.EtcdClientImpl{Connection: stubEtcdConnection},
		RoleRepository: StubRoleRepositoryImpl{Connection: stubConnection},
	}

	permissionService := service.PermissionServiceImpl{
		EtcdClient:           cache.EtcdClientImpl{Connection: stubEtcdConnection},
		PermissionRepository: StubPermissionRepositoryImpl{Connection: stubConnection},
	}

	migration = Migration{
		service:               ser,
		userService:           userService,
		roleService:           roleService,
		operatorPolicyService: operatorPolicyService,
		permissionService:     permissionService,
	}
}

func TestNewMigration(t *testing.T) {
	NewMigration()
}

func TestMigration_V1(t *testing.T) {
	migration.V1()
}
