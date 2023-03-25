package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"miniWiki/pkg/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

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
	}

	err = initLogger(*cfg)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, cfg.Database.DatabaseURL)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(cfg.Database.DatabaseURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	app := &Application{
		Config:   cfg,
		Database: pool,
		Server: &http.Server{
			Addr:              ":" + cfg.Server.Port,
			ReadHeaderTimeout: cfg.Server.Timeout,
			Handler:           InitRouter(pool, db, *cfg),
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
		shutdownCtx, cFunc := context.WithTimeout(ctx, app.Config.Server.Timeout)
		defer cFunc()
		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
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
