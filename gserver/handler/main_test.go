package handler

import (
	"os"
	"testing"

	"github.com/tomoyane/grant-n-z/gserver/common/config"
	"github.com/tomoyane/grant-n-z/gserver/log"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("APP_ENV", "test")
	config.InitConfig()
	log.InitLogger(config.LogLevel)

	ret := m.Run()
	os.Exit(ret)
}
