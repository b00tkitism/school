package permission_group

import (
	"net/http"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPermissionGroup retrieves a permission group by ID
func (ctrl *PermissionGroupController) GetPermissionGroup(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	group, err := ctrl.Service.GetPermissionGroup(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, util.GenerateResponse(false, "Permission group not found", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(true, "Permission group fetched", group))
}
