package main

import (
	"github.com/ilmovilDev/gorm/config"
	"github.com/ilmovilDev/gorm/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
