package models

import "gorm.io/gorm"

type PermissionGroup struct {
	*gorm.Model
	GroupID      uint
	PermissionID uint
}
