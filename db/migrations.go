package db

import "school/models"

func RunMigrations() {
	DB.AutoMigrate(&models.RefreshToken{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.MessageStatus{})
	DB.AutoMigrate(&models.Group{})
	DB.AutoMigrate(&models.Permission{})
	DB.AutoMigrate(&models.UserGroups{})
	DB.AutoMigrate(&models.UserPermissions{})
}
