package gserver

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"net/http"
	"os/signal"

	"github.com/tomoyane/grant-n-z/gserver/common/config"
	"github.com/tomoyane/grant-n-z/gserver/common/driver"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/migration"
	"github.com/tomoyane/grant-n-z/gserver/router"
)

var (
	exitCode   = make(chan int)
	signalCode = make(chan os.Signal, 1)
	server     = &http.Server{Addr: ":8080", Handler: nil}
	banner     = `Start to grant-n-z server :8080
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

func init() {
	config.InitConfig()
	log.InitLogger(config.App.LogLevel)
	driver.InitDriver()
}

func NewGrantNZServer() GrantNZServer {
	log.Logger.Info("New GrantNZServer")
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
	g.migration()
	g.runRouter()
	go g.subscribeSignal(context.TODO(), signalCode, exitCode)
	shutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go g.gracefulShutdown(shutdownCtx, exitCode, *server)
	g.runServer(*server)
}

func (g GrantNZServer) migration() {
	migration.NewMigration().V1()
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

func (g GrantNZServer) subscribeSignal(ctx context.Context, signalCode chan os.Signal, exitCode chan int) {
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

func (g GrantNZServer) gracefulShutdown(ctx context.Context, exitCode chan int, server http.Server) {
	code := <-exitCode
	server.Shutdown(ctx)

	driver.Db.Close()
	log.Logger.Info("Closed MySQL connection")

	driver.Redis.Close()
	log.Logger.Info("Closed Redis connection")

	log.Logger.Info("Shutdown gracefully")
	os.Exit(code)
}
