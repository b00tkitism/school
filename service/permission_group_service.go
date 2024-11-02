package service

import (
	"school/models"
	"school/repository"
)

type PermissionGroupService struct {
	Repo *repository.PermissionGroupRepository
}

// CreatePermissionGroup creates a new permission group
func (s *PermissionGroupService) CreatePermissionGroup(group models.PermissionGroup) error {
	return s.Repo.CreatePermissionGroup(group)
}

// GetPermissionGroup retrieves a permission group by ID
func (s *PermissionGroupService) GetPermissionGroup(id uint) (*models.PermissionGroup, error) {
	return s.Repo.GetPermissionGroupByID(id)
}

// UpdatePermissionGroup updates a permission group by ID
func (s *PermissionGroupService) UpdatePermissionGroup(id uint, group models.PermissionGroup) error {
	return s.Repo.UpdatePermissionGroup(id, group)
}

// DeletePermissionGroup deletes a permission group by ID
func (s *PermissionGroupService) DeletePermissionGroup(id uint) error {
	return s.Repo.DeletePermissionGroup(id)
}

// ListPermissionGroups lists all permission groups with pagination
func (s *PermissionGroupService) ListPermissionGroups(offset, limit int) ([]models.PermissionGroup, error) {
	return s.Repo.ListPermissionGroups(offset, limit)
}
