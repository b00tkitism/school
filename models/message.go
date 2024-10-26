package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	*gorm.Model
	SenderID    uint           `json:"sender_id"`
	RecipientID uint           `json:"recipient_id"` // Nullable for broadcasts
	Title       string         `json:"title"`
	Content     string         `json:"content"`
	IsBroadcast bool           `json:"is_broadcast"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
