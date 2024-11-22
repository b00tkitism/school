package group

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RemoveUserFromGroupRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}

func (controller *GroupController) RemoveUserFromGroup(c *gin.Context) {
	// Check if the user has permission to remove users from groups
	userID := c.Keys["user_id"].(uint)
	permission, err := controller.UserService.UserHasPermission(userID, e.RemoveUserFromGroupPermission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while checking permissions", nil))
		return
	}
	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you do not have permission to remove users from groups", nil))
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
	userID2, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid user_id", nil))
		return
	}

	// Check if the user exists
	userExists, err := controller.UserService.UserExistsByID(uint(userID2))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching user", nil))
		return
	}
	if !userExists {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "user not found", nil))
		return
	}

	// Check if the user is a member of the group
	isMember, err := controller.GroupService.IsUserInGroup(uint(groupID), uint(userID2))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while checking user membership", nil))
		return
	}
	if !isMember {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "user is not a member of this group", nil))
		return
	}

	// Remove the user from the group
	err = controller.GroupService.RemoveUserFromGroup(uint(groupID), uint(userID2))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while removing user from group", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "user removed from group successfully", nil))
}
