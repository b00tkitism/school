package db

import "school/models"

func RunMigrations() {
	DB.AutoMigrate(&models.RefreshToken{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Message{})
	DB.AutoMigrate(&models.MessageStatus{})
}
