package group

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersListResponse struct {
	UserIDs []uint `json:"user_ids"`
}

func (controller *GroupController) ListUsersInGroup(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.AccessGroupInformationPermission))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to access groups list", nil))
		return
	}

	groupID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}

	groups, err := controller.GroupService.ListUsersInGroup(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching groups", nil))
		return
	}

	var groupList UsersListResponse = UsersListResponse{
		UserIDs: groups,
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "groups fetched", groupList))
}
