package db

import "school/models"

func RunMigrations() {
	DB.AutoMigrate(&models.RefreshToken{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.MessageStatus{})
	DB.AutoMigrate(&models.Group{})
	DB.AutoMigrate(&models.GroupPermission{})
	DB.AutoMigrate(&models.Permission{})
	DB.AutoMigrate(&models.UserGroup{})
	DB.AutoMigrate(&models.UserPermission{})
}
