package group

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RemovePermissionsRequest struct {
	PermissionIDs []uint `json:"permission_ids"`
}

func (controller *GroupController) RemovePermissionsFromGroup(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), e.RemovePermissionFromGroupPermission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to remove permission from a group", nil))
		return
	}

	groupID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}
	exists, err := controller.GroupService.IsGroupExistsByID(uint(groupID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching group", nil))
		return
	}

	if !exists {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}

	var request AssignPermissionsRequest
	c.BindJSON(&request)

	for _, permission := range request.PermissionIDs {
		if util.IsPermissionValid(permission) {
			controller.GroupService.RemovePermissionFromGroup(uint(groupID), permission)
		}
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "permissions removed", nil))
}
