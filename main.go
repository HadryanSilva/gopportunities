package main

import (
	"github.com/HadryanSilva/gopportunities/config"
	"github.com/HadryanSilva/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errf("Error initializing config: %v", err)
		return
	}

	router.Initialize()
}
