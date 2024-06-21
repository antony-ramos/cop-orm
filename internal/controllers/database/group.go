package database

import (
	"github.com/antony-ramos/cop-orm/internal/entity"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	entity.Group
}

func (gorm *Gorm) CreateGroup(group entity.Group) error {
	return gorm.db.Create(&Group{Group: group}).Error
}

func (gorm *Gorm) GetGroup(groupID string) (entity.Group, error) {
	var group Group
	err := gorm.db.First(&group, groupID).Error
	return group.Group, err
}

func (gorm *Gorm) GetAllGroups() ([]entity.Group, error) {
	var groups []Group
	err := gorm.db.Find(&groups).Error
	if err != nil {
		return nil, err
	}

	var entityGroups []entity.Group
	for _, group := range groups {
		entityGroups = append(entityGroups, group.Group)
	}

	return entityGroups, nil
}

func (gorm *Gorm) DeleteGroup(groupID string) error {
	return gorm.db.Delete(&Group{}, groupID).Error
}

func (gorm *Gorm) ModifyGroup(groupID string, group entity.Group) error {
	return gorm.db.Model(&Group{}).Where("id = ?", groupID).Updates(group).Error
}
