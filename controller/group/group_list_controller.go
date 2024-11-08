package group

import (
	"net/http"
	"school/e"
	"school/util"

	"github.com/gin-gonic/gin"
)

type GroupsListResponse struct {
	ID uint `json:"id"`
}

func (controller *GroupController) ListGroups(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.AccessGroupInformationPermission))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to access groups list", nil))
		return
	}

	groups, err := controller.GroupService.ListGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching groups", nil))
		return
	}

	var groupList []GroupsListResponse
	for _, group := range groups {
		groupList = append(groupList, GroupsListResponse{
			ID: group.ID,
		})
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "groups fetched", groupList))
}
