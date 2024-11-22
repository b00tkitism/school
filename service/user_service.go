package service

import (
	"errors"
	"school/models"
	"school/repository"
	"school/util"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (service *UserService) GetAdmin(username, password string) (*models.Users, error) {
	return service.Repo.GetAdmin(username, util.EncodeMD5(password))
}

func (service *UserService) GetUser(username, password string) (*models.Users, error) {
	return service.Repo.GetUser(username, util.EncodeMD5(password))
}

func (service *UserService) GetUserByID(userID uint) (*models.Users, error) {
	return service.Repo.GetUserByID(userID)
}

func (service *UserService) GetAdminByID(userID uint) (*models.Users, error) {
	return service.Repo.GetAdminByID(userID)
}

func (service *UserService) GetPermissionsByID(userID uint) ([]models.Permission, error) {
	return service.Repo.GetPermissionsByID(userID)
}

func (service *UserService) UserHasPermission(userID uint, permissionID uint) (bool, error) {
	permissions, err := service.GetPermissionsByID(userID)
	for _, v := range permissions {
		if permissionID == v.ID {
			return true, nil
		}
	}
	return false, err
}

func (service *UserService) UserExists(username string) (bool, error) {
	return service.Repo.UserExists(username)
}

func (service *UserService) UserExistsByID(userID uint) (bool, error) {
	return service.Repo.UserExistsByID(userID)
}

func (service *UserService) NewUser(user models.Users) error {
	user.Password = util.EncodeMD5(user.Password)
	return service.Repo.NewUser(user)
}

func (service *UserService) DeleteUser(userID uint) error {
	return service.Repo.DeleteUser(userID)
}

func (service *UserService) ModifyUser(userID uint, username, password, fullName, phoneNumber, idCode string, gender, isAdmin *bool) error {
	exists, err := service.Repo.UserExistsByID(userID)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("user not found")
	}

	hashedPassword := util.EncodeMD5(password)
	return service.Repo.UpdateUser(userID, &username, &hashedPassword, &fullName, &phoneNumber, &idCode, gender, isAdmin)
}

func (s *UserService) GetUsersCount() (int64, error) {
	// Calculate the offset based on the page and pageSize
	return s.Repo.GetUsersCount()
}

func (s *UserService) GetPaginatedUsers(page, pageSize int) ([]models.Users, error) {
	// Calculate the offset based on the page and pageSize
	offset := (page - 1) * pageSize

	// Fetch users from the repository with the calculated offset and limit
	users, err := s.Repo.GetPaginatedUsers(offset, pageSize)
	if err != nil {
		return nil, err
	}
	return users, nil
}
