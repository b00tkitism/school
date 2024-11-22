package group

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (controller *GroupController) ListGroupPermissions(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), e.AccessGroupPermissionsInformationPermission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to access group permissions information", nil))
		return
	}

	// Get the group ID from the URL parameters
	groupID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}

	// Check if the group exists
	exists, err := controller.GroupService.IsGroupExistsByID(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching group", nil))
		return
	}

	if !exists {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "group not found", nil))
		return
	}

	// Fetch the permissions for the group
	permissions, err := controller.GroupService.GetPermissionsByGroupID(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching group permissions", nil))
		return
	}

	// Return the permissions
	c.JSON(http.StatusOK, util.GenerateResponse(true, "permissions fetched successfully", permissions))
}
