package main

import (
	"miniWiki/internal/app"

	"github.com/sirupsen/logrus"
)

// MiniWiki app
//
// As I find more and more useful resources both for my job and for my hobbies, simply
// bookmarking links becomes inefficient. I need a space to store summaries of the resources and
// categorize them by multiple criteria. This way, no resource is lost and the time spent searching
// for and researching a topic is cut down significantly.
//
// swagger:meta

func main() {
	application, err := app.New()
	if err != nil {
		panic(err)
	}

	err = application.Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
