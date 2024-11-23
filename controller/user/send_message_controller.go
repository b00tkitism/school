package user

import (
	"net/http"
	"school/e"
	"school/util"

	"github.com/gin-gonic/gin"
)

type AdminSendMessageRequest struct {
	RecipientID   uint   `json:"recipient_id"`
	RecipientType string `json:"recipient_type"`
	Title         string `json:"title"`
	Message       string `json:"message"`
}

func (controller *UserController) SendMessage(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.SendMessagePermission))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
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
	if request.RecipientType == "single" {
		err = controller.MessageService.SendMessage(userID, request.RecipientID, request.Title, request.Message)
	} else {
		list, _ := controller.GroupService.ListUsersInGroup(request.RecipientID)
		for _, user := range list {
			err = controller.MessageService.SendMessage(userID, user, request.Title, request.Message)
		}
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while sending message", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "message sent", nil))
}
