package main

import (
	"log"

	"github.com/1140251/go-template/config"
	"github.com/1140251/go-template/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
