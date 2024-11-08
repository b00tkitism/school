package repository

import "gorm.io/gorm"

type GroupRepository struct {
	DB *gorm.DB
}
