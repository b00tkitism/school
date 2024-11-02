package message

import (
	"net/http"
	"school/util"

	"github.com/gin-gonic/gin"
)

type ReadMessageRequest struct {
	MessageID int `json:"message_id"`
}

func (controller *MessageController) ReadMessage(c *gin.Context) {
	var request ReadMessageRequest
	c.BindJSON(&request)

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Fetch messages with read status and pagination
	err := controller.MessageService.MarkMessageAsRead(userID.(uint), uint(request.MessageID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to mark message as read"})
		return
	}

	// Respond with paginated messages
	c.JSON(http.StatusOK, util.GenerateResponse(true, "message marked as read", nil))
}
