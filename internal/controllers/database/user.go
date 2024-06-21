package database

import (
	"github.com/antony-ramos/cop-orm/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	entity.User
}

func (gorm *Gorm) CreateUser(user entity.User) error {
	return gorm.db.Create(&User{User: user}).Error
}

func (gorm *Gorm) GetUser(userID string) (entity.User, error) {
	var user User
	err := gorm.db.First(&user, userID).Error
	return user.User, err
}

func (gorm *Gorm) GetAllUsers() ([]entity.User, error) {
	var users []User
	err := gorm.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	var entityUsers []entity.User
	for _, user := range users {
		entityUsers = append(entityUsers, user.User)
	}

	return entityUsers, nil
}

func (gorm *Gorm) DeleteUser(userID string) error {
	return gorm.db.Delete(&User{}, userID).Error
}

func (gorm *Gorm) ModifyUser(userID string, user entity.User) error {
	return gorm.db.Model(&User{}).Where("id = ?", userID).Updates(user).Error
}
