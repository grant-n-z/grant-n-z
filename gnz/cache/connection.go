package cache

import (
	"fmt"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"

	"github.com/tomoyane/grant-n-z/gnz/config"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var connection *clientv3.Client

// Initialize cache database driver
func InitEtcd() {
	if strings.EqualFold(config.Etcd.Host, "") || strings.EqualFold(config.Etcd.Port, "") {
		log.Logger.Info("Not use etcd")
		return
	}

	// 10millisecond timeout
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("%s:%s", config.Etcd.Host, config.Etcd.Port)},
		DialTimeout: 20 * time.Millisecond,
	})

	if err != nil {
		log.Logger.Warn(err.Error())
		log.Logger.Warn("Cannot connect etcd. If needs to high performance, run GrantNZ cache server with etcd")
		Close()
		return
	}
	log.Logger.Info("Connected etcd", config.Etcd.Host)
	connection = client
}

// Close etcd
func Close() {
	if connection != nil {
		_ = connection.Close()
		log.Logger.Info("Closed etcd connection")
	} else {
		log.Logger.Info("Already closed etcd connection")
	}
}
