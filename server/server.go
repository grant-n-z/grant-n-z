package server

import (
	"context"
	"fmt"
	"github.com/tomoyane/grant-n-z/server/common/driver"
	"os"
	"syscall"
	"time"

	"net/http"
	"os/signal"

	"github.com/tomoyane/grant-n-z/server/common/config"
	"github.com/tomoyane/grant-n-z/server/log"
	"github.com/tomoyane/grant-n-z/server/router"
)

var (
	exitCode   = make(chan int)
	signalCode = make(chan os.Signal, 1)
	server     = &http.Server{Addr: ":8080", Handler: nil}
	banner     = `start grant-n-z server :8080
___________________________________________________
    ____                      _      
   / __/ _    ____   _____ __//_      _____   ____ 
  / /__ //__ /__ /  /___ //_ __/     /___ /  /__ /
 / /_ //___///_//_ //  //  //_  === //  // === //__
/____///   /_____///  //  /__/     //  //     /___/
___________________________________________________
High performance authentication and authorization. version is %s
`
)

type GrantNZServer struct {
	router router.Router
}

func NewGrantNZServer() GrantNZServer {
	log.Logger.Info("Inject `Router`, `CronHandler`, `PolicyService` to `GrantNZServer`")
	signal.Notify(
		signalCode,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	return GrantNZServer{
		router: router.NewRouter(),
	}
}

func (g GrantNZServer) Run() {
	g.runRouter()
	go g.subscribeSignal(signalCode, exitCode)
	go g.gracefulShutdown(exitCode, *server)
	g.runServer(*server)
}

func (g GrantNZServer) runRouter() {
	g.router.Init()
	g.router.V1()
}

func (g GrantNZServer) runServer(server http.Server) {
	fmt.Printf(banner, config.App.Version)
	if err := server.ListenAndServe(); err != nil {
		log.Logger.Error("Error run grant-n-z server", err.Error())
		os.Exit(1)
	}
}

func (g GrantNZServer) subscribeSignal(signalCode chan os.Signal, exitCode chan int) {
	for {
		s := <-signalCode
		switch s {
		case syscall.SIGHUP:
			log.Logger.Info("Catch signal SIGHUP")

		case syscall.SIGINT:
			log.Logger.Info("Catch signal SIGINT")
			exitCode <- 0

		case syscall.SIGTERM:
			log.Logger.Info("Catch signal SIGTERM")
			exitCode <- 0

		case syscall.SIGQUIT:
			log.Logger.Info("Catch signal SIGQUIT")
			exitCode <- 0

		case syscall.SIGKILL:
			log.Logger.Info("Catch signal SIGKILL")
			exitCode <- 0

		default:
			log.Logger.Error("Unknown signal code")
			exitCode <- 1
		}
	}
}

func (g GrantNZServer) gracefulShutdown(exitCode chan int, server http.Server) {
	code := <-exitCode
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)

	driver.Db.Close()
	log.Logger.Info("Closed database connection")

	driver.Redis.Close()
	log.Logger.Info("Closed Redis connection")

	log.Logger.Info("Shutdown gracefully")
	os.Exit(code)
}
