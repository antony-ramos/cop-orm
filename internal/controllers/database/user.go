package database

import (
	"github.com/antony-ramos/cop-orm/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	entity.User
}
