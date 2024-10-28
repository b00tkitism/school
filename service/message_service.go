package service

import (
	"school/models"
	"school/repository"
)

type MessageService struct {
	Repo *repository.MessageRepository
}

// SendMessage handles sending a message to a specific user or as a broadcast
func (service *MessageService) SendMessage(senderID uint, recipientID uint, title, content string) error {
	isBroadcast := recipientID == 0
	return service.Repo.SendMessage(senderID, recipientID, title, content, isBroadcast)
}

// MarkMessageAsRead sets a message as read for a user
func (service *MessageService) MarkMessageAsRead(userID, messageID uint) error {
	return service.Repo.MarkMessageAsRead(userID, messageID)
}

// GetMessagesWithStatus retrieves messages with read status for a specific user
func (service *MessageService) GetMessagesWithStatus(userID uint, page, pageSize int) ([]models.MessageWithStatus, error) {
	// Calculate offset based on page and pageSize
	offset := (page - 1) * pageSize

	// Retrieve messages from the repository
	return service.Repo.GetMessagesWithStatus(userID, pageSize, offset)
}

func (service *MessageService) GetMessagesCount(userID uint) (int64, error) {
	return service.Repo.GetMessagesCount(userID)
}
