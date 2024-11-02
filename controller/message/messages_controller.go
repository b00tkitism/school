package message

import (
	"net/http"
	"school/models"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MessagesResponse struct {
	Page     int                        `json:"page"`
	PageSize int                        `json:"page_size"`
	Total    int64                      `json:"total"`
	Messages []models.MessageWithStatus `json:"messages"`
}

// GetMessages retrieves messages for a specific user with pagination and ordered by newest
func (controller *MessageController) GetMessages(c *gin.Context) {
	// Assuming userID is stored in context by JWT middleware
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Get page and pageSize from query parameters, with default values
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size parameter"})
		return
	}

	// Fetch messages with read status and pagination
	messages, err := controller.MessageService.GetMessagesWithStatus(userID.(uint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}

	total, err := controller.MessageService.GetMessagesCount(userID.(uint))
	if err != nil {
		return
	}

	response := MessagesResponse{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Messages: messages,
	}
	// Respond with paginated messages
	c.JSON(http.StatusOK, util.GenerateResponse(true, "messages fetched", response))
}
