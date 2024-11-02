package models

import "gorm.io/gorm"

type UserPermissions struct {
	*gorm.Model
	UserID       uint
	PermissionID uint
}
