package group

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ModifyGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (controller *GroupController) UpdateGroup(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.ModifyGroupPermission))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to modify group information", nil))
		return
	}

	groupID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}

	exists, err := controller.GroupService.IsGroupExistsByID(uint(groupID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while checking group existness", nil))
		return
	}

	if !exists {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid group_id", nil))
		return
	}

	var request ModifyGroupRequest
	c.BindJSON(&request)

	err = controller.GroupService.UpdateGroup(uint(groupID), request.Name, request.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while updating group information", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "group modified", nil))
}
