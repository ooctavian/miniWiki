package app

import (
	"os"
	"strings"

	"miniWiki/internal/config"

	"github.com/sirupsen/logrus"
)

func initLogger(cfg config.LoggerConfig) error {
	lvl, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	if strings.ToLower(cfg.Output) == "stdout" {
		logrus.SetOutput(os.Stdout)
	} else {
		f, err := os.OpenFile(cfg.Output, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}
	if strings.ToLower(cfg.Formatter) == "json" {
		logrus.SetFormatter(new(logrus.JSONFormatter))
	}
	return err
}
