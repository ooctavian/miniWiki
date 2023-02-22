package app

import (
	"os"

	"miniWiki/config"

	"github.com/sirupsen/logrus"
)

func initLogger(cfg config.Config) error {
	lvl, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(new(logrus.JSONFormatter))
	return err
}
