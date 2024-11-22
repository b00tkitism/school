package models

import "gorm.io/gorm"

type GroupPermission struct {
	*gorm.Model
	GroupID      uint       `gorm:"index"`
	PermissionID uint       `gorm:"index"`
	Group        Group      `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
}
