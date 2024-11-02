package models

import (
	"time"

	"gorm.io/gorm"
)

type MessageStatus struct {
	*gorm.Model
	UserID    uint      `json:"user_id"`
	MessageID uint      `json:"message_id"`
	IsRead    bool      `json:"is_read"`
	ReadAt    time.Time `json:"read_at,omitempty"`
}

type MessageWithStatus struct {
	ID          uint       `json:"id"`
	SenderID    uint       `json:"sender_id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	IsBroadcast bool       `json:"is_broadcast"`
	IsRead      bool       `json:"is_read"`
	ReadAt      *time.Time `json:"read_at,omitempty"`
	CreatedAt   *time.Time `json:"send_time"`
}
