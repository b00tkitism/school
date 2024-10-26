package models

import (
	"gorm.io/gorm"
)

type RefreshToken struct {
	*gorm.Model
	UserID  uint
	User    Users  `gorm:"foreignKey:UserID"`
	Token   string `gorm:"not null;unique"` // Unique token value
	Revoked bool   `gorm:"default:false"`   // Indicates if the token has been revoked
}
