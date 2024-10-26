package service

import (
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

func (service *UserService) GetAdminByID(userID uint) (*models.Users, error) {
	return service.Repo.GetAdminByID(userID)
}

func (service *UserService) UserExists(username string) (bool, error) {
	return service.Repo.UserExists(username)
}

func (service *UserService) NewUser(user models.Users) error {
	user.Password = util.EncodeMD5(user.Password)
	return service.Repo.NewUser(user)
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
