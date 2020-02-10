package core

import (
	"context"
	"fmt"
	"github.com/tomoyane/grant-n-z/gnzcache/timer"
	"os"
	"syscall"
	"time"

	"os/signal"

	"github.com/tomoyane/grant-n-z/gnz/config"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

const (
	BannerFilePath = "./grant_n_z_cache.txt"
	ConfigFilePath = "./grant_n_z_cache.yaml"
)

var (
	exitCode   = make(chan int)
	signalCode = make(chan os.Signal, 1)
)

type GrantNZCacheUpdater struct {
	UpdateTimer timer.UpdateTimer
}

func init() {
	log.InitLogger(config.App.LogLevel)
	driver.InitGrantNZDb()
}

func NewGrantNZCacheUpdater() GrantNZCacheUpdater {
	log.Logger.Info("New GrantNZCacheUpdater")
	signal.Notify(
		signalCode,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	return GrantNZCacheUpdater{UpdateTimer: timer.NewUpdateTimer()}
}

// Run GrantNZ cache
func (g GrantNZCacheUpdater) Run() {
	bannerText, err := config.ConvertFileToStr(BannerFilePath)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Could't read %s file", BannerFilePath), err.Error())
		os.Exit(1)
	}
	fmt.Printf(bannerText, config.App.Version)

	g.UpdateTimer.Run()
	go g.subscribeSignal(signalCode, exitCode)
	shutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go g.gracefulShutdown(shutdownCtx, exitCode)
}

// Subscribe signal
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

// Graceful shutdown
func (g GrantNZCacheUpdater) gracefulShutdown(ctx context.Context, exitCode chan int) {
	code := <-exitCode
	driver.CloseConnection()
	log.Logger.Info("Shutdown gracefully GrantNZ Cache")
	os.Exit(code)
}
