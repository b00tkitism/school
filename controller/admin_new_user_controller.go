package controller

import (
	"net/http"
	"school/models"
	"school/util"

	"github.com/gin-gonic/gin"
)

func (controller *UserController) NewUser(c *gin.Context) {
	admin, err := controller.UserService.GetAdminByID(c.Keys["user_id"].(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if !admin.CanCreateUsers {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "you don't have permission to create a new user", nil))
		return
	}

	var request models.Users
	c.BindJSON(&request)

	if request.Username == "" || request.Password == "" || request.FullName == "" || request.PhoneNumber == "" || request.IDCode == "" {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "username or password is empty", nil))
		return
	}

	userExists, err := controller.UserService.UserExists(request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching users", nil))
		return
	}

	if userExists {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "username already exists", nil))
		return
	}

	controller.UserService.NewUser(request)
	c.JSON(http.StatusOK, util.GenerateResponse(true, "user created", nil))
}
