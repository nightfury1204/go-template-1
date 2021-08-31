package cmd

import (
	"context"
	"fmt"
	"go-template/api"
	"go-template/config"
	"go-template/infra/db"
	"go-template/logger"
	"go-template/repo"
	"go-template/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var cfgFile string

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the http server",
	RunE:  serve,
}

func init() {
	ServeCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".env", "the config file")
}

func serve(cmd *cobra.Command, args []string) error {
	err := godotenv.Load(cfgFile)
	if err != nil {
		return err
	}
	log.Printf("Environment: %s", os.Getenv("ENV"))

	appConfig := config.GetApp()

	log := logger.DefaultOutStructLogger

	DB, err := db.NewDB()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(DB)
	if err != nil {
		return err
	}

	bookRepo := repo.NewBookRepo(DB, log)

	bookService := service.NewBookService(bookRepo, log)

	bookCtrl := api.NewBookController(bookService, log)

	router := api.NewRouter(bookCtrl)

	return startApiServer(appConfig, router)
}

func startApiServer(appConfig *config.Application, router chi.Router) error {
	srv := &http.Server{
		Addr:    getAddressFromHostAndPort(appConfig.Host, appConfig.Port),
		Handler: router,
		//ErrorLog: logger.DefaultErrLogger,
		//WriteTimeout: cfg.WriteTimeout,
		//ReadTimeout:  cfg.ReadTimeout,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	stopCh := make(chan os.Signal, 1)
	// notify the interrupt signal
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	log.Println("api server started", "listening port", appConfig.Port)

	<-stopCh
	log.Println("api server stopping")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(appConfig.GracefulTimeout)*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed:%v", err)
	}
	log.Println("server exited properly")
	return nil
}

func getAddressFromHostAndPort(host string, port int) string {
	addr := host
	if port != 0 {
		addr = addr + ":" + strconv.Itoa(port)
	}
	return addr
}
