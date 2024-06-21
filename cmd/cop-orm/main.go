package main

import (
	"os"

	"github.com/antony-ramos/cop-orm/config"
	"github.com/antony-ramos/cop-orm/internal/app"
)

func main() {

	// App Configuration
	configPath := os.Getenv("COP_ORM_CONFIG_PATH")
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	app.Start(cfg)
}
