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

	err := repo.DB.Raw(`
        SELECT DISTINCT p.id, p.name, p.description
        FROM permissions p
        LEFT JOIN group_permissions gp ON p.id = gp.permission_id
        LEFT JOIN user_groups ug ON gp.group_id = ug.group_id
        WHERE ug.user_id = ?
        
        UNION
        
        SELECT DISTINCT p.id, p.permission_name, p.description
        FROM permissions p
        LEFT JOIN user_permissions up ON p.id = up.permission_id
        WHERE up.user_id = ?
    `, userID, userID).Scan(&permissions).Error

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

func (repo *UserRepository) UpdateUser(userID uint, username, password, fullName, phoneNumber, idCode string, gender, isAdmin *bool) error {
	// Prepare a map with non-nil fields to be updated
	updates := map[string]interface{}{}
	if username != "" {
		updates["username"] = username
	}
	if fullName != "" {
		updates["full_name"] = fullName
	}
	if phoneNumber != "" {
		updates["phone_number"] = phoneNumber
	}
	if isAdmin != nil {
		updates["is_admin"] = *isAdmin
	}

	// Perform the update
	return repo.DB.Model(&models.Users{}).Where("id = ?", userID).Updates(updates).Error
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
