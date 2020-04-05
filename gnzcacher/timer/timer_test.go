package timer

import (
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/ctx"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzcacher/service"
)

var (
	runner Runner
)

func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	stubEtcdConnection, _ := clientv3.New(clientv3.Config{
		Endpoints:            []string{},
		DialTimeout:          5 * time.Millisecond,
		DialKeepAliveTimeout: 5 * time.Millisecond,
	})
	etcdClient := cache.EtcdClientImpl{
		Connection: stubEtcdConnection,
		Ctx:        ctx.GetCtx(),
	}

	stubPolicyRepository := driver.PolicyRepositoryImpl{Connection: stubConnection}
	stubPermissionRepository := driver.PermissionRepositoryImpl{Connection: stubConnection}
	stubRoleRepository := driver.RoleRepositoryImpl{Connection: stubConnection}
	stubServiceRepository := driver.ServiceRepositoryImpl{Connection: stubConnection}
	stubUserRepository := driver.UserRepositoryImpl{Connection: stubConnection}

	extractorService = service.ExtractorServiceImpl{
		PolicyRepository:     stubPolicyRepository,
		PermissionRepository: stubPermissionRepository,
		RoleRepository:       stubRoleRepository,
		ServiceRepository:    stubServiceRepository,
		UserRepository:       stubUserRepository,
	}

	updaterService = service.UpdaterServiceImpl{EtcdClient: etcdClient}
	runner = RunnerImpl{
		UpdaterService:   updaterService,
		ExtractorService: extractorService,
	}
}

// Test start
func TestStart(t *testing.T) {
	updateTimer := UpdateTimerImpl{
		Runner: runner,
		Ticker: time.NewTicker(1 * time.Second),
	}

	exitCode := make(chan int)
	go updateTimer.Start(exitCode)

	time.Sleep(2 * time.Second)

	exitCode <- 1

	updateTimer.Stop()
}
