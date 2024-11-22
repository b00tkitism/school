package group

import (
	"fmt"
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssignUserToGroupRequest struct {
	UserID uint `json:"user_id"`
}

func (controller *GroupController) AssignUserToGroup(c *gin.Context) {
	// Check if the user has permission to assign users to groups
	userID := c.Keys["user_id"].(uint)
	permission, err := controller.UserService.UserHasPermission(userID, e.AssignUserToGroupPermission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while checking permissions", nil))
		return
	}
	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you do not have permission to assign users to groups", nil))
		return
	}

	// Parse group ID from URL
	groupID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}

	// Check if the group exists
	groupExists, err := controller.GroupService.IsGroupExistsByID(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching group", nil))
		return
	}
	if !groupExists {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "group not found", nil))
		return
	}

	// Parse request body
	var request AssignUserToGroupRequest
	// if err := c.BindJSON(&request); err != nil {
	c.BindJSON(&request)
	// c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid request body", nil))
	// return
	// }

	// Check if the user exists
	userExists, err := controller.UserService.UserExistsByID(request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching user", nil))
		return
	}
	if !userExists {
		fmt.Println(request.UserID)
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "user not found", nil))
		return
	}

	// Check if the group has any special permissions (only admins can join)
	isAdminOnly, err := controller.GroupService.IsGroupAdminOnly(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while checking group permissions", nil))
		return
	}
	if isAdminOnly {
		// Check if the user is an admin
		user, err := controller.UserService.GetUserByID(request.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while checking user admin status", nil))
			return
		}
		if !user.IsAdmin {
			c.JSON(http.StatusForbidden, util.GenerateResponse(false, "only admins can join this group", nil))
			return
		}
	}

	// Assign the user to the group
	err = controller.GroupService.AssignUserToGroup(uint(groupID), request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while assigning user to group", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "user assigned to group successfully", nil))
}
