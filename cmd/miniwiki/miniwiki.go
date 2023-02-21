package main

import (
	"context"
	"errors"
	"net/http"
	"os"

	"miniWiki"
	"miniWiki/utils"

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

func main() {
	utils.InitLogger()
	cfg, err := miniWiki.InitConfig()
	if err != nil {
		panic(err)
	}
	createDir(cfg.Database.ImageDir + "/resources")
	pool, err := pgxpool.Connect(context.Background(), cfg.Database.DatabaseUrl)

	err = http.ListenAndServe(":"+cfg.Server.Port, miniWiki.InitRouter(pool))
	if err != nil {
		logrus.Fatalf("Error while starting server %v", err)
	}
	pool.Close()
}
