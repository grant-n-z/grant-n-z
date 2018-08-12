package test

import (
	"testing"
	"github.com/tomoyane/embedded-mysql-container/container"
	"os"
	"github.com/tomoyane/grant-n-z/common"
	"time"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	containerDaemon := container.ContainerDaemonImpl{}.New()
	embeddedMysql := container.MysqlConfigImpl{}.New()

	containerId := containerDaemon.StartEmbeddedMysql()

	time.Sleep(10 * time.Second)

	embeddedMysql.AddSchema("auth_server")
	embeddedMysql.CreateTable("CREATE TABLE auth_server.users (" +
		"id int(11) NOT NULL AUTO_INCREMENT," +
		"uuid varchar(128) NOT NULL," +
		"username varchar(128) NOT NULL," +
		"email varchar(128) NOT NULL," +
		"password varchar(128) NOT NULL," +
		"created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"PRIMARY KEY (id)," +
		"UNIQUE (email))" +
		" ENGINE=InnoDB DEFAULT CHARSET=utf8")

	// Init database
	common.InitDB()

	code := m.Run()

	// After
	containerDaemon.FinishEmbeddedMysql(containerId)

	os.Exit(code)
}
