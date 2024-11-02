package user

import (
	"net/http"
	"school/e"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserResponse struct {
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	IDCode      string `json:"id_code"`
	Gender      bool   `json:"gender"`
	IsAdmin     bool   `json:"is_admin"`
}

func (controller *UserController) GetUser(c *gin.Context) {
	permission, err := controller.UserService.UserHasPermission(c.Keys["user_id"].(uint), uint(e.AccessUserInformation))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while checking permission", nil))
		return
	}

	if !permission {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to access user information", nil))
		return
	}

	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "invalid id", nil))
		return
	}

	user, err := controller.UserService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching user", nil))
		return
	}

	response := GetUserResponse{
		Username:    user.Username,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		IDCode:      user.IDCode,
		Gender:      user.Gender,
		IsAdmin:     user.IsAdmin,
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "user fetched", response))
}
