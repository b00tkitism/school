package user

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ModifyUserRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	IDCode      string `json:"id_code"`
	Gender      bool   `json:"gender"`
	IsAdmin     bool   `json:"is_admin"`
}

func (controller *UserController) ModifyUser(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.ModifyUserPermission))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to modify user", nil))
		return
	}

	var request ModifyUserRequest
	c.BindJSON(&request)

	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid id", nil))
		return
	}

	err = controller.UserService.ModifyUser(uint(userID), request.Username, request.Password, request.FullName, request.PhoneNumber, request.IDCode, &request.Gender, &request.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while modifying user", nil))
		return
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "user modified", nil))
}
