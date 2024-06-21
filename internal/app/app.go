package app

import (
	"github.com/antony-ramos/cop-orm/config"
	"github.com/antony-ramos/cop-orm/internal/controllers/database"
	"github.com/antony-ramos/cop-orm/internal/controllers/http"
	"github.com/antony-ramos/cop-orm/internal/usecase"
)

func Start(cfg *config.Config) {
	database := database.Gorm{}
	err := database.Start(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	usecase := usecase.New(&database)

	http.Run(usecase)
}
