package app

import (
	"os"
	"strings"

	"miniWiki/internal/config"

	"github.com/sirupsen/logrus"
)

func initLogger(cfg config.Config) error {
	lvl, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	logrus.SetOutput(os.Stdout)
	if strings.ToLower(cfg.Logger.Formatter) == "json" {
		logrus.SetFormatter(new(logrus.JSONFormatter))
	}
	return err
}
