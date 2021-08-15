package timer

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzcacher/service"
)

var (
	extractorService service.ExtractorService
	updaterService   service.UpdaterService
)

func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})
	etcdClient := cache.EtcdClientImpl{Connection: stubEtcdConnection}

	stubPolicyRepository := driver.RdbmsPolicyRepository{Connection: stubConnection}
	stubPermissionRepository := driver.RdbmsPermissionRepository{Connection: stubConnection}
	stubRoleRepository := driver.RdbmsRoleRepository{Connection: stubConnection}
	stubServiceRepository := driver.RdbmsServiceRepository{Connection: stubConnection}
	stubUserRepository := driver.RdbmsUserRepository{Connection: stubConnection}

	extractorService = service.ExtractorServiceImpl{
		PolicyRepository:     stubPolicyRepository,
		PermissionRepository: stubPermissionRepository,
		RoleRepository:       stubRoleRepository,
		ServiceRepository:    stubServiceRepository,
		UserRepository:       stubUserRepository,
	}

	updaterService = service.UpdaterServiceImpl{EtcdClient: etcdClient}
}

// Test run
func TestRun(t *testing.T) {
	runner := RunnerImpl{
		UpdaterService:   updaterService,
		ExtractorService: extractorService,
	}
	runner.Run()
}
