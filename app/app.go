package app

import (
	"context"
	"errors"
	"net/http"
	"os"

	"miniWiki/config"

	"github.com/jackc/pgx/v4/pgxpool"
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
}

func New() (*Application, error) {
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

	createDir(cfg.Database.ImageDir + "/resources")
	pool, err := pgxpool.Connect(context.Background(), cfg.Database.DatabaseURL)
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
			Handler:           initRouter(pool, *cfg),
		},
	}
	return app, nil
}

func (app *Application) Start() error {
	err := app.Server.ListenAndServe()
	if err != nil {
		logrus.Fatalf("Error starting server %v", err)
		return err
	}
	app.Database.Close()
	return err
}
