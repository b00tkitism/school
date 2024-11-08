package repository

import (
	"school/models"

	"gorm.io/gorm"
)

type GroupRepository struct {
	DB *gorm.DB
}

func (repo *GroupRepository) IsGroupExists(name string) (bool, error) {
	var count int64
	var group models.Group = models.Group{
		GroupName: name,
	}

	err := repo.DB.Model(&models.Group{}).Find(&group).Count(&count).Error
	return count >= 1, err
}

func (repo *GroupRepository) IsGroupExistsByID(id uint) (bool, error) {
	var count int64
	var group models.Group

	err := repo.DB.Model(&models.Group{}).Find(&group, id).Count(&count).Error
	return count >= 1, err
}

func (repo *GroupRepository) CreateGroup(name, description string) (uint, error) {
	var group models.Group = models.Group{
		GroupName:   name,
		Description: description,
	}

	err := repo.DB.Model(&models.Group{}).Create(&group).Error
	return group.ID, err
}

func (repo *GroupRepository) DeleteGroup(id uint) error {
	return repo.DB.Model(&models.Group{}).Delete(&models.Group{}, id).Error
}

func (repo *GroupRepository) ListGroups() ([]models.Group, error) {
	var groups []models.Group
	err := repo.DB.Model(&models.Group{}).Find(&groups).Error
	return groups, err
}

func (repo *GroupRepository) GetGroup(id uint) (models.Group, error) {
	var group models.Group
	err := repo.DB.Model(&models.Group{}).Find(&group, id).Error
	return group, err
}

func (repo *GroupRepository) UpdateGroup(id uint, name, description string) error {
	updates := models.Group{}

	// Apply updates conditionally
	if name != "" {
		updates.GroupName = name
	}
	if description != "" {
		updates.Description = description
	}

	return repo.DB.Model(&models.Group{}).Where("id = ?", id).Updates(updates).Error
}

func (repo *GroupRepository) GroupHasPermission(groupID, permissionID uint) (bool, error) {
	gPermission := models.GroupPermission{
		GroupID:      groupID,
		PermissionID: permissionID,
	}

	var count int64
	err := repo.DB.Model(&models.GroupPermission{}).Where(gPermission).Find(&gPermission).Count(&count).Error
	return count >= 1, err
}

func (repo *GroupRepository) AssignPermissionToGroup(groupID, permissionID uint) error {
	gPermission := models.GroupPermission{
		GroupID:      groupID,
		PermissionID: permissionID,
	}

	err := repo.DB.Model(&models.GroupPermission{}).Create(&gPermission).Error
	return err
}

func (repo *GroupRepository) RemovePermissionFromGroup(groupID, permissionID uint) error {
	gPermission := models.GroupPermission{
		GroupID:      groupID,
		PermissionID: permissionID,
	}

	err := repo.DB.Model(&models.GroupPermission{}).Where(gPermission).Delete(&models.GroupPermission{}).Error
	return err
}
