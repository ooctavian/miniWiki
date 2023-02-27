package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"miniWiki/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func createDir(dirName string) {
	if _, err := os.Stat(dirName); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			logrus.Fatalf("Creating dir error: %v", err)
		}
	}
}

type Application struct {
	Config   *config.Config
	Database *pgxpool.Pool
	Server   *http.Server
	Context  context.Context
}

func New() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
		return nil, err
	}

	err = initLogger(*cfg)
	if err != nil {
		panic(err)
		return nil, err
	}

	ctx := context.Background()

	createDir(cfg.Database.ImageDir + "/resources")
	pool, err := pgxpool.Connect(ctx, cfg.Database.DatabaseURL)
	if err != nil {
		panic(err)
		return nil, err
	}
	app := &Application{
		Config:   cfg,
		Database: pool,
		Server: &http.Server{
			Addr:              ":" + cfg.Server.Port,
			ReadHeaderTimeout: cfg.Server.Timeout,
			Handler:           InitRouter(pool, *cfg),
		},
		Context: ctx,
	}
	return app, nil
}

func (app *Application) Start() error {
	// Shameless stolen from: https://github.com/go-chi/chi/blob/master/_examples/graceful/main.go
	// Server run context
	ctx, cancelFunc := context.WithCancel(app.Context)

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(ctx, app.Config.Server.Timeout)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logrus.Fatal("shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := app.Server.Shutdown(shutdownCtx)
		if err != nil {
			logrus.Fatal(err)
		}
		cancelFunc()
	}()

	// Run the server
	err := app.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	// Wait for server context to be stopped
	<-ctx.Done()
	return err
}
