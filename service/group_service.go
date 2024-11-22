package service

import (
	"school/e"
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

func (service *GroupService) ListUsersInGroup(id uint) ([]uint, error) {
	users, err := service.Repo.ListUsersInGroup(id)
	if users != nil {
		ids := []uint{}
		for _, user := range users {
			ids = append(ids, user.UserID)
		}
		return ids, err
	}
	return []uint{}, err
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

func (service *GroupService) GetPermissionsByGroupID(groupID uint) ([]models.Permission, error) {
	// Call the repository to fetch permissions by group ID
	permissions, err := service.Repo.GetPermissionsByGroupID(groupID)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// IsGroupAdminOnly checks if a group has permissions that make it admin-only
func (service *GroupService) IsGroupAdminOnly(groupID uint) (bool, error) {
	permissions, err := service.Repo.GetPermissionsByGroupID(groupID)
	if err != nil {
		return false, err
	}

	// Check if any of the group's permissions match admin-only permissions
	for _, permission := range permissions {
		for _, adminPermission := range e.ValidPermissions {
			if permission.ID == adminPermission {
				return true, nil
			}
		}
	}
	return false, nil
}

// AssignUserToGroup assigns a user to a group
func (service *GroupService) AssignUserToGroup(groupID, userID uint) error {
	// Check if the user is already in the group
	exists, err := service.Repo.IsUserInGroup(groupID, userID)
	if err != nil {
		return err
	}
	if exists {
		return nil // User is already assigned to the group
	}

	// Assign the user to the group
	return service.Repo.AssignUserToGroup(groupID, userID)
}

func (service *GroupService) IsUserInGroup(groupID, userID uint) (bool, error) {
	return service.Repo.IsUserInGroup(groupID, userID)
}

func (service *GroupService) RemoveUserFromGroup(groupID, userID uint) error {
	// Check if the user is a member of the group
	isMember, err := service.Repo.IsUserInGroup(groupID, userID)
	if err != nil {
		return err
	}
	if !isMember {
		return nil // User is not a member, nothing to remove
	}

	// Remove the user from the group
	return service.Repo.RemoveUserFromGroup(groupID, userID)
}
