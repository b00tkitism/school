package models

import "gorm.io/gorm"

type UserPermissionOverride struct {
	*gorm.Model
	UserID       uint
	PermissionID uint
	Override     bool // true means explicit grant, false means explicit deny
}
