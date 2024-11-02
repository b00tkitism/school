package models

import "gorm.io/gorm"

type UserGroups struct {
	*gorm.Model
	UserID  uint
	GroupID uint
}
