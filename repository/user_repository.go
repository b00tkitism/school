package repository

import (
	"fmt"
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

func (repo *UserRepository) GetUserByID(userID uint) (*models.Users, error) {
	var user *models.Users
	err := repo.DB.Model(&models.Users{}).Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAdminByID(userID uint) (*models.Users, error) {
	var admin *models.Users
	err := repo.DB.Model(&models.Users{}).Where("id = ?", userID).Where("is_admin = ?", true).Find(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *UserRepository) GetPermissionsByID(userID uint) ([]models.Permission, error) {
	var permissions []models.Permission

	// SQL query to fetch permissions from group and handle overrides
	err := repo.DB.Raw(`
        WITH GroupPermissions AS (
            SELECT DISTINCT p.id, p.name, p.description
            FROM permissions p
            LEFT JOIN group_permissions gp ON p.id = gp.permission_id
            LEFT JOIN user_groups ug ON gp.group_id = ug.group_id
            WHERE ug.user_id = ?
        )
        SELECT DISTINCT gp.id, gp.name, gp.description
        FROM GroupPermissions gp
        LEFT JOIN user_permission_overrides upo ON gp.id = upo.permission_id AND upo.user_id = ?
        WHERE (upo.override IS NULL OR upo.override = TRUE)
        
        UNION
        
        SELECT DISTINCT p.id, p.name, p.description
        FROM permissions p
        INNER JOIN user_permission_overrides upo ON p.id = upo.permission_id
        WHERE upo.user_id = ? AND upo.override = TRUE
    `, userID, userID, userID).Scan(&permissions).Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (repo *UserRepository) UserExists(username string) (bool, error) {
	var count int64
	err := repo.DB.Model(&models.Users{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *UserRepository) UserExistsByID(userID uint) (bool, error) {
	var count int64
	err := repo.DB.Model(&models.Users{}).Where("id = ?", userID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *UserRepository) NewUser(user models.Users) error {
	return repo.DB.Model(&models.Users{}).Create(&user).Error
}

func (repo *UserRepository) DeleteUser(userID uint) error {
	return repo.DB.Model(&models.Users{}).Where("id = ?", userID).Error
}

func (repo *UserRepository) UpdateUser(userID uint, username, password, fullName, phoneNumber, idCode *string, gender, isAdmin *bool) error {
	updateData := models.Users{}

	if username != nil {
		updateData.Username = *username
	}
	if password != nil {
		updateData.Password = *password
	}
	if fullName != nil {
		updateData.FullName = *fullName
	}
	if phoneNumber != nil {
		updateData.PhoneNumber = *phoneNumber
	}
	if idCode != nil {
		updateData.IDCode = *idCode
	}
	if gender != nil {
		updateData.Gender = *gender
	}
	if isAdmin != nil {
		updateData.IsAdmin = *isAdmin
	}

	// Use Select to only update fields that are non-nil
	tx := repo.DB.Model(&models.Users{}).Where("id = ?", userID).Updates(updateData)

	if err := tx.Error; err != nil {
		return fmt.Errorf("failed to update user %d: %w", userID, err)
	}

	return nil
}

func (repo *UserRepository) GetUsersCount() (int64, error) {
	var count int64
	err := repo.DB.Model(&models.Users{}).Count(&count).Error
	return count, err
}

func (repo *UserRepository) GetPaginatedUsers(offset, limit int) ([]models.Users, error) {
	var users []models.Users
	err := repo.DB.Select("id, username, full_name, is_admin").Where("is_hidden=false").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}
