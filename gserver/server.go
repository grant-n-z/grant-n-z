package gserver

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"net/http"
	"os/signal"

	"github.com/gorilla/mux"

	"github.com/tomoyane/grant-n-z/gserver/config"
	"github.com/tomoyane/grant-n-z/gserver/ctx"
	"github.com/tomoyane/grant-n-z/gserver/driver"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/middleware"
	"github.com/tomoyane/grant-n-z/gserver/route"
)

const Port = "8080"

var (
	exitCode   = make(chan int)
	signalCode = make(chan os.Signal, 1)
	server     = &http.Server{Addr: fmt.Sprintf(":%s", Port), Handler: nil}
	banner     = `Start to grant-n-z server :%s
_________________________________________________________________________________________________
     ____                      _           
    / __/ _    ____   _____ __//_      _____   ____     ____  ___     _   _     _   ___     _
   / /__ //__ /__ /  /___ //_ __/     /___ /  /__ /    \ __//  _ \   //__ \\   // /  _ \   //__
  / /_ //___///_//_ //  //  //_  === //  // === //__   _\\  \ /_ /  /___/  \\ //  \ /_ /  / __/
 /____///   /_____///  //  /__/     //  //     /___/  /__/   \____ //       \ /    \____ // 
_________________________________________________________________________________________________
High performance authentication and authorization. version is %s
`
)

type GrantNZServer struct {
	router route.Router
}

func init() {
	ctx.InitContext()
	config.InitGrantNZServerConfig()
	log.InitLogger(config.App.LogLevel)
	driver.InitGrantNZDb()
}

func NewGrantNZServer() GrantNZServer {
	log.Logger.Info("New GrantNZServer")
	signal.Notify(
		signalCode,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	return GrantNZServer{router: route.NewRouter()}
}

// Run GrantNZ server
func (g GrantNZServer) Run() {
	g.migration()
	go g.subscribeSignal(signalCode, exitCode)
	shutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go g.gracefulShutdown(shutdownCtx, exitCode, *server)
	g.runServer(g.runRouter())
}

// Migrate to required initialize data
func (g GrantNZServer) migration() {
	middleware.NewMigration().V1()
}

// Run router
func (g GrantNZServer) runRouter() *mux.Router {
	return g.router.Run()
}

// Run server
func (g GrantNZServer) runServer(router *mux.Router) {
	fmt.Printf(banner, Port, config.App.Version)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", Port), router); err != nil {
		log.Logger.Error("Error run grant-n-z server", err.Error())
		os.Exit(1)
	}
}

// Subscribe signal
func (g GrantNZServer) subscribeSignal(signalCode chan os.Signal, exitCode chan int) {
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
func (g GrantNZServer) gracefulShutdown(ctx context.Context, exitCode chan int, server http.Server) {
	code := <-exitCode
	server.Shutdown(ctx)

	driver.CloseConnection()
	log.Logger.Info("Shutdown gracefully")
	os.Exit(code)
}
