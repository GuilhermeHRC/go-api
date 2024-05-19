package main

import (
	"github.com/GuilhermeHRC/go-api/config"
	"github.com/GuilhermeHRC/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}
	// Initialize Router
	router.Init()
}
