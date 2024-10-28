package repository

import (
	"school/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) GetAdmin(username, password string) (*models.Users, error) {
	var admin *models.Users
	err := repo.DB.Model(&models.Users{}).Where("username = ?", username).Where("password = ?", password).Where("is_admin = ?", true).Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *UserRepository) GetUser(username, password string) (*models.Users, error) {
	var admin *models.Users
	err := repo.DB.Model(&models.Users{}).Where("username = ?", username).Where("password = ?", password).Where("is_admin = ?", false).Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *UserRepository) GetAdminByID(userID uint) (*models.Users, error) {
	var admin *models.Users
	err := repo.DB.Model(&models.Users{}).Where("id = ?", userID).Where("is_admin = ?", true).Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *UserRepository) UserExists(username string) (bool, error) {
	var count int64
	err := repo.DB.Model(&models.Users{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *UserRepository) NewUser(user models.Users) error {
	return repo.DB.Model(&models.Users{}).Create(&user).Error
}

func (repo *UserRepository) GetUsersCount() (int64, error) {
	var count int64
	err := repo.DB.Model(&models.Users{}).Count(&count).Error
	return count, err
}

func (repo *UserRepository) GetPaginatedUsers(offset, limit int) ([]models.Users, error) {
	var users []models.Users
	err := repo.DB.Select("id, username, full_name").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}
