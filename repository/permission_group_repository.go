package repository

import (
	"school/models"

	"gorm.io/gorm"
)

type PermissionGroupRepository struct {
	DB *gorm.DB
}

// CreatePermissionGroup creates a new permission group
func (repo *PermissionGroupRepository) CreatePermissionGroup(group models.PermissionGroup) error {
	return repo.DB.Create(&group).Error
}

// GetPermissionGroupByID retrieves a permission group by its ID
func (repo *PermissionGroupRepository) GetPermissionGroupByID(id uint) (*models.PermissionGroup, error) {
	var group models.PermissionGroup
	if err := repo.DB.Model(&models.PermissionGroup{}).First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

// UpdatePermissionGroup updates a permission group by ID
func (repo *PermissionGroupRepository) UpdatePermissionGroup(id uint, group models.PermissionGroup) error {
	return repo.DB.Model(&models.PermissionGroup{}).Where("id = ?", id).Updates(group).Error
}

// DeletePermissionGroup deletes a permission group by ID
func (repo *PermissionGroupRepository) DeletePermissionGroup(id uint) error {
	return repo.DB.Delete(&models.PermissionGroup{}, id).Error
}

// ListPermissionGroups lists all permission groups with pagination
func (repo *PermissionGroupRepository) ListPermissionGroups(offset, limit int) ([]models.PermissionGroup, error) {
	var groups []models.PermissionGroup
	if err := repo.DB.Offset(offset).Limit(limit).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}
