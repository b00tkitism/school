package permission_group

import (
	"net/http"
	"school/models"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdatePermissionGroup updates a permission group by ID
func (ctrl *PermissionGroupController) UpdatePermissionGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var group models.PermissionGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "Invalid input", nil))
		return
	}
	if err := ctrl.Service.UpdatePermissionGroup(uint(id), group); err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "Failed to update permission group", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(true, "Permission group updated", group))
}
