package group

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (controller *GroupController) DeleteGroup(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.CreateGroupPermission))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to delete group", nil))
		return
	}

	groupID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}

	err = controller.GroupService.DeleteGroup(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while deleting group", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "group deleted", nil))
}
