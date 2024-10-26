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
func (service *MessageService) GetMessagesWithStatus(userID uint) ([]models.MessageWithStatus, error) {
	return service.Repo.GetMessagesWithStatus(userID)
}
