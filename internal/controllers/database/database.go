package database

import (
	"fmt"

	"github.com/antony-ramos/cop-orm/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Start(config.Postgres) error
}

type Gorm struct {
	Database
	db *gorm.DB
}

func (repository *Gorm) Start(cfg config.Postgres) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d,sslmode=disable TimeZone=Europe/Paris",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)
	var err error
	repository.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
