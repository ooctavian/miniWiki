package main

import (
	"miniWiki/app"

	"github.com/sirupsen/logrus"
)

func main() {
	application, err := app.New()
	if err != nil {
		logrus.Fatal(err)
	}

	err = application.Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
