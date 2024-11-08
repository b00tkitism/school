package group

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetGroupResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (controller *GroupController) GetGroup(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.AccessGroupInformationPermission))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to access groups information", nil))
		return
	}

	groupID, err := strconv.ParseUint(c.Param("group_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}

	group, err := controller.GroupService.GetGroup(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching group", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "group fetched", GetGroupResponse{
		ID:          group.ID,
		Name:        group.GroupName,
		Description: group.Description,
	}))
}
