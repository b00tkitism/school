package group

import (
	"net/http"
	"school/e"
	"school/util"

	"github.com/gin-gonic/gin"
)

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateGroupResponse struct {
	GroupID uint `json:"group_id"`
}

func (controller *GroupController) CreateGroup(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.CreateGroupPermission))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to create group", nil))
		return
	}

	var request CreateGroupRequest
	c.BindJSON(&request)

	exists, err := controller.GroupService.IsGroupExists(request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while checking duplicate groups", nil))
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "group name already exists", nil))
		return
	}

	groupID, err := controller.GroupService.CreateGroup(request.Name, request.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while creating group", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "group created", CreateGroupResponse{GroupID: groupID}))
}
