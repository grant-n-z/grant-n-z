package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzcacher/timer"
)

const (
	BannerFilePath = "./grant_n_z_cacher.txt"
	ConfigFilePath = "./grant_n_z_cacher.yaml"
)

var (
	exitCode   = make(chan int)
	signalCode = make(chan os.Signal, 1)
)

type GrantNZCacher struct {
	UpdateTimer timer.UpdateTimer
	Database    driver.Database
}

func init() {
}

func NewGrantNZCacher() GrantNZCacher {
	log.InitLogger(common.App.LogLevel)
	common.InitGrantNZCacherConfig(ConfigFilePath)
	database := driver.NewDatabase()
	database.Connect()
	cache.InitEtcd()
	log.Logger.Info("New GrantNZCacher")

	signal.Notify(
		signalCode,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	return GrantNZCacher{UpdateTimer: timer.NewUpdateTimer()}
}

// Start GrantNZ cache
func (g GrantNZCacher) Run() {
	bannerText, err := common.ConvertFileToStr(BannerFilePath)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Could't read %s file", BannerFilePath), err.Error())
		os.Exit(1)
	}
	fmt.Printf(bannerText)

	go g.subscribeSignal(signalCode, exitCode)
	go g.Database.PingRdbms()

	exitCode := g.UpdateTimer.Start(exitCode)
	g.gracefulShutdown(exitCode)
}

// Subscribe signal
func (g GrantNZCacher) subscribeSignal(signalCode chan os.Signal, exitCode chan int) {
	for {
		s := <-signalCode
		switch s {
		case syscall.SIGHUP:
			log.Logger.Info("Caught signal SIGHUP")

		case syscall.SIGINT:
			log.Logger.Info("Caught signal SIGINT")
			exitCode <- 0

		case syscall.SIGTERM:
			log.Logger.Info("Caught signal SIGTERM")
			exitCode <- 0

		case syscall.SIGQUIT:
			log.Logger.Info("Caught signal SIGQUIT")
			exitCode <- 0

		case syscall.SIGKILL:
			log.Logger.Info("Caught signal SIGKILL")
			exitCode <- 0

		default:
			log.Logger.Error("Unknown signal code")
			exitCode <- 1
		}
	}
}

// Graceful shutdown
func (g GrantNZCacher) gracefulShutdown(code int) {
	g.Database.Close()
	cache.Close()

	log.Logger.Info("Shutdown gracefully GrantNZ Cacher")
	os.Exit(code)
}
