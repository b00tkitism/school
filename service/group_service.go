package service

import (
	"school/models"
	"school/repository"
)

type GroupService struct {
	Repo *repository.GroupRepository
}

func (service *GroupService) IsGroupExists(name string) (bool, error) {
	return service.Repo.IsGroupExists(name)
}

func (service *GroupService) IsGroupExistsByID(id uint) (bool, error) {
	return service.Repo.IsGroupExistsByID(id)
}

func (service *GroupService) CreateGroup(name string, description string) (uint, error) {
	return service.Repo.CreateGroup(name, description)
}

func (service *GroupService) DeleteGroup(id uint) error {
	return service.Repo.DeleteGroup(id)
}

func (service *GroupService) ListGroups() ([]models.Group, error) {
	return service.Repo.ListGroups()
}

func (service *GroupService) GetGroup(id uint) (models.Group, error) {
	return service.Repo.GetGroup(id)
}

func (service *GroupService) UpdateGroup(id uint, name, description string) error {
	return service.Repo.UpdateGroup(id, name, description)
}

func (service *GroupService) GroupHasPermission(groupID, permissionID uint) (bool, error) {
	return service.Repo.GroupHasPermission(groupID, permissionID)
}

func (service *GroupService) AssignPermissionToGroup(groupID, permissionID uint) error {
	exists, _ := service.GroupHasPermission(groupID, permissionID)
	if !exists {
		return service.Repo.AssignPermissionToGroup(groupID, permissionID)
	}
	return nil
}

func (service *GroupService) RemovePermissionFromGroup(groupID, permissionID uint) error {
	exists, _ := service.GroupHasPermission(groupID, permissionID)
	if exists {
		return service.Repo.RemovePermissionFromGroup(groupID, permissionID)
	}
	return nil
}
