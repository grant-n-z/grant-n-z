package test

import (
	"testing"
	"github.com/tomoyane/embedded-mysql-container/container"
	"os"
	"github.com/tomoyane/grant-n-z/common"
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")

	// Before
	containerDaemon := container.ContainerDaemonImpl{}.NewContainerDaemonImpl()

	containerDaemon.InitDocker()
	containerDaemon.PullImage("docker.io/library/mysql:5.7")

	containerId := containerDaemon.BuildImage(
		"mysql:5.7",
		"embedded_mysql_container")

	containerDaemon.StartContainer(containerId)

	// Init database
	common.InitDB()

	code := m.Run()

	// After
	containerDaemon.StopContainer(containerId)
	containerDaemon.DeleteContainer(containerId)

	os.Exit(code)
}
