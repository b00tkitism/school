package models

import "gorm.io/gorm"

type UserGroup struct {
	*gorm.Model
	UserID  uint  `gorm:"index"`
	GroupID uint  `gorm:"index"`
	User    Users `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Group   Group `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE"`
}
