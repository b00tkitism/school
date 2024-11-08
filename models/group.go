package models

import "gorm.io/gorm"

type Group struct {
	*gorm.Model
	GroupName   string `gorm:"unique;not null"`
	Description string
}
