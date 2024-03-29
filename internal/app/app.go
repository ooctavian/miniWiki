package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"miniWiki/internal/config"

	"github.com/joho/godotenv"
	"github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Application struct {
	Config  *config.Config
	Server  *http.Server
	Context context.Context
}

func New() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg, err := config.InitConfig()
	if err != nil {
		return nil, err
	}
	err = initLogger(cfg.Logger)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	db, err := gorm.Open(postgres.Open(cfg.Database.DatabaseURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gorm_logrus.New(),
	})

	if err != nil {
		return nil, err
	}

	app := &Application{
		Config: cfg,
		Server: &http.Server{
			Addr:              ":" + cfg.Server.Port,
			ReadHeaderTimeout: cfg.Server.Timeout,
			Handler:           InitRouter(db, *cfg),
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
