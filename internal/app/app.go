package app

import (
	"github.com/antony-ramos/cop-orm/internal/controllers/http"
	"github.com/antony-ramos/cop-orm/internal/usecase"
)

func Start() {
	usecase := usecase.New()
	http.Run(usecase)
}
