package repository

import (
	"school/models"
	"time"

	"gorm.io/gorm"
)

type MessageRepository struct {
	DB *gorm.DB
}

// SendMessage creates a message in the database
func (repo *MessageRepository) SendMessage(senderID uint, recipientID uint, title, content string, isBroadcast bool) error {
	message := models.Message{
		SenderID:    senderID,
		RecipientID: recipientID,
		Title:       title,
		Content:     content,
		IsBroadcast: isBroadcast,
	}
	return repo.DB.Create(&message).Error
}

// MarkMessageAsRead sets a message as read for a user
func (repo *MessageRepository) MarkMessageAsRead(userID, messageID uint) error {
	status := models.MessageStatus{
		UserID:    userID,
		MessageID: messageID,
		IsRead:    true,
		ReadAt:    time.Now(),
	}
	return repo.DB.Where("user_id = ?", userID).Where("message_id = ?", messageID).Save(&status).Error
}

// GetMessagesWithStatus retrieves messages along with read status for a specific user
func (repo *MessageRepository) GetMessagesWithStatus(userID uint, limit, offset int) ([]models.MessageWithStatus, error) {
	var messages []models.MessageWithStatus

	// Retrieve messages with read status for the specified user
	err := repo.DB.Model(&models.Message{}).
		Select("messages.id, messages.sender_id, messages.content, messages.is_broadcast, COALESCE(ms.is_read, false) AS is_read, COALESCE(ms.read_at, null) AS read_at").
		Joins("LEFT JOIN message_statuses AS ms ON ms.message_id = messages.id AND ms.user_id = ?", userID).
		Where("messages.recipient_id = ? OR messages.is_broadcast = ?", userID, true).
		Order("messages.created_at DESC").
		Limit(limit).
		Offset(offset).
		Scan(&messages).Error

	return messages, err
}

func (repo *MessageRepository) GetMessagesCount(userID uint) (int64, error) {
	var count int64
	err := repo.DB.Model(&models.Message{}).Where("recipient_id = ? OR is_broadcast = ?", userID, true).Count(&count).Error
	return count, err
}
