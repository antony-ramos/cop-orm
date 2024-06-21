package database

import (
	"fmt"

	"github.com/antony-ramos/cop-orm/config"
	"github.com/antony-ramos/cop-orm/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Start(config.Postgres) error
	CreateUser(user entity.User) error
	GetUser(userID string) (entity.User, error)
	DeleteUser(userID string) error
	ModifyUser(userID string, user entity.User) error
	GetAllUsers() ([]entity.User, error)
	CreateGroup(group entity.Group) error
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

	repository.db.AutoMigrate(&User{})
	repository.db.AutoMigrate(&Group{})

	return nil
}
