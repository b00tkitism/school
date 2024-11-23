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
	return repo.DB.Model(&models.Group{}).Delete(id).Error
}

func (repo *GroupRepository) ListGroups() ([]models.Group, error) {
	var groups []models.Group
	err := repo.DB.Model(&models.Group{}).Find(&groups).Error
	return groups, err
}

func (repo *GroupRepository) ListUsersInGroup(id uint) ([]models.UserGroup, error) {
	var users []models.UserGroup
	err := repo.DB.Model(&models.UserGroup{}).Where("group_id=?", id).Find(&users).Error
	return users, err
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

func (repo *GroupRepository) GetPermissionsByGroupID(groupID uint) ([]models.Permission, error) {
	var permissions []models.Permission

	err := repo.DB.Model(&models.GroupPermission{}).Table("group_permissions").
		Select("permissions.id, permissions.name, permissions.description").
		Joins("JOIN permissions ON group_permissions.permission_id = permissions.id").
		Where("group_permissions.group_id = ?", groupID).
		Scan(&permissions).Error

	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// IsUserInGroup checks if a user is already a member of a group
func (repo *GroupRepository) IsUserInGroup(groupID, userID uint) (bool, error) {
	var count int64

	err := repo.DB.Model(&models.UserGroup{}).
		Where("group_id = ? AND user_id = ?", groupID, userID).
		Count(&count).Error

	return count > 0, err
}

// AssignUserToGroup adds a user to a group
func (repo *GroupRepository) AssignUserToGroup(groupID, userID uint) error {
	userGroup := models.UserGroup{
		GroupID: groupID,
		UserID:  userID,
	}

	err := repo.DB.Create(&userGroup).Error
	return err
}

func (repo *GroupRepository) RemoveUserFromGroup(groupID, userID uint) error {
	err := repo.DB.Where("group_id = ? AND user_id = ?", groupID, userID).
		Delete(&models.UserGroup{}).Error
	return err
}
