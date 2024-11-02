package permission_group

import (
	"net/http"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeletePermissionGroup deletes a permission group by ID
func (ctrl *PermissionGroupController) DeletePermissionGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.Service.DeletePermissionGroup(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "Failed to delete permission group", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(true, "Permission group deleted", nil))
}