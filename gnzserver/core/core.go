package core

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"net/http"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/tomoyane/grant-n-z/gnz/cache"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/driver"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
)

const (
	BannerFilePath = "./grant_n_z_server.txt"
	ConfigFilePath = "./grant_n_z_server.yaml"
)

var (
	exitCode   = make(chan int)
	signalCode = make(chan os.Signal, 1)
	port       string
	server     *http.Server
)

type GrantNZServer struct {
	router   Router
	database driver.Database
}

func NewGrantNZServer() GrantNZServer {
	common.InitGrantNZServerConfig(ConfigFilePath)
	log.InitLogger(common.App.LogLevel)
	database := driver.NewDatabase()
	database.Connect()
	cache.InitEtcd()
	log.Logger.Info("New GrantNZServer")

	signal.Notify(
		signalCode,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	return GrantNZServer{
		router:   NewRouter(),
		database: database,
	}
}

// Run
// Start GrantNZ server
func (g GrantNZServer) Run() {
	port = common.GServer.Port
	server = &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: nil}

	g.migration()
	shutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	go g.subscribeSignal(signalCode, exitCode)
	go g.gracefulShutdown(shutdownCtx, exitCode, server)
	go g.database.PingRdbms()

	g.runServer(g.runRouter())
}

// Migrate to required initialize data
func (g GrantNZServer) migration() {
	middleware.NewMigration().V1()
}

// Start router
func (g GrantNZServer) runRouter() *mux.Router {
	return g.router.Run()
}

// Start server
func (g GrantNZServer) runServer(router *mux.Router) {
	bannerText, err := common.ConvertFileToStr(BannerFilePath)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Could't read %s file", BannerFilePath), err.Error())
		os.Exit(1)
	}

	fmt.Printf(bannerText, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
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
func (g GrantNZServer) gracefulShutdown(ctx context.Context, exitCode chan int, server *http.Server) {
	code := <-exitCode
	server.Shutdown(ctx)

	g.database.Close()
	cache.Close()

	log.Logger.Info("Shutdown gracefully GrantNZ Server")
	os.Exit(code)
}
