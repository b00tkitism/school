package models

import "gorm.io/gorm"

type GroupPermission struct {
	*gorm.Model
	GroupID      uint
	PermissionID uint
}
