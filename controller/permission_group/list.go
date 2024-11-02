package permission_group

import (
	"net/http"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListPermissionGroups lists all permission groups with pagination
func (ctrl *PermissionGroupController) ListPermissionGroups(c *gin.Context) {
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	groups, err := ctrl.Service.ListPermissionGroups(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "Failed to fetch permission groups", nil))
		return
	}
	c.JSON(http.StatusOK, util.GenerateResponse(true, "Permission groups fetched", groups))
}
