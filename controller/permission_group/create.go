package permission_group

import (
	"net/http"
	"school/models"
	"school/util"

	"github.com/gin-gonic/gin"
)

// CreatePermissionGroup creates a new permission group
func (ctrl *PermissionGroupController) CreatePermissionGroup(c *gin.Context) {
	var group models.PermissionGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "Invalid input", nil))
		return
	}
	if err := ctrl.Service.CreatePermissionGroup(group); err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "Failed to create permission group", nil))
		return
	}
	c.JSON(http.StatusCreated, util.GenerateResponse(true, "Permission group created", group))
}
