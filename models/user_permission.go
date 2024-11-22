package models

import "gorm.io/gorm"

type UserPermission struct {
	*gorm.Model
	UserID       uint       `gorm:"index"`
	PermissionID uint       `gorm:"index"`
	IsGranted    bool       `gorm:"default:true"`
	Source       string     `gorm:"size:20"` // group or independent
	User         Users      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Permission   Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
}
