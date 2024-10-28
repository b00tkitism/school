package controller

import (
	"net/http"
	"school/util"

	"github.com/gin-gonic/gin"
)

type AdminSendMessageRequest struct {
	RecipientID uint   `json:"recipient_id"`
	Title       string `json:"title"`
	Message     string `json:"message"`
}

func (controller *UserController) SendMessage(c *gin.Context) {
	admin, err := controller.UserService.GetAdminByID(c.Keys["user_id"].(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !admin.CanSendMessages {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to send message", nil))
		return
	}

	var request AdminSendMessageRequest
	c.BindJSON(&request)

	if request.Title == "" || request.Message == "" {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "title or message is empty", nil))
		return
	}

	userID := c.Keys["user_id"].(uint)
	err = controller.MessageService.SendMessage(userID, request.RecipientID, request.Title, request.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while sending message", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "message sent", nil))
}
