package gnzcache

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"os/signal"

	"github.com/tomoyane/grant-n-z/gnzserver/config"
	"github.com/tomoyane/grant-n-z/gnzserver/ctx"
	"github.com/tomoyane/grant-n-z/gnzserver/driver"
	"github.com/tomoyane/grant-n-z/gnzserver/log"
	"github.com/tomoyane/grant-n-z/gnzserver/route"
)

var (
	exitCode   = make(chan int)
	signalCode = make(chan os.Signal, 1)
	banner     = `Start to grant-n-z cache updater
___________________________________________________
    ____                      _      
   / __/ _    ____   _____ __//_      _____   ____ 
  / /__ //__ /__ /  /___ //_ __/     /___ /  /__ /
 / /_ //___///_//_ //  //  //_  === //  // === //__
/____///   /_____///  //  /__/     //  //     /___/
___________________________________________________
Version is %s
`
)

type GrantNZCacheUpdater struct {
	router route.Router
}

func init() {
	ctx.InitContext()
	config.InitGrantNZServerConfig()
	log.InitLogger(config.App.LogLevel)
	driver.InitGrantNZDb()
}

func NewGrantNZCacheUpdater() GrantNZCacheUpdater {
	log.Logger.Info("New GrantNZCacheScheduler")
	signal.Notify(
		signalCode,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	return GrantNZCacheUpdater{router: route.NewRouter()}
}

func (g GrantNZCacheUpdater) Run() {
	fmt.Printf(banner, config.App.Version)
	go g.subscribeSignal(signalCode, exitCode)
	shutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go g.gracefulShutdown(shutdownCtx, exitCode)
}

func (g GrantNZCacheUpdater) subscribeSignal(signalCode chan os.Signal, exitCode chan int) {
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

func (g GrantNZCacheUpdater) gracefulShutdown(ctx context.Context, exitCode chan int) {
	code := <-exitCode
	driver.CloseConnection()
	log.Logger.Info("Shutdown gracefully")
	os.Exit(code)
}
