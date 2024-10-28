package controller

import (
	"net/http"
	"school/models"
	"school/util"

	"github.com/gin-gonic/gin"
)



type AdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (controller *UserController) AdminLogin(c *gin.Context) {
	var adminRequest AdminRequest
	c.BindJSON(&adminRequest)

	if adminRequest.Username == "" || adminRequest.Password == "" {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "username or password is empty", nil))
		return
	}

	admin, err := controller.UserService.GetAdmin(adminRequest.Username, adminRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching admin", nil))
		return
	}

	if (*admin == models.Users{}) {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "invalid credentials", nil))
		return
	}

	accessToken, err := util.GenerateJWT(admin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while generating jwt", nil))
		return
	}

	refreshToken, err := controller.TokenService.GetRefreshToken(admin.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while getting refresh token", nil))
		return
	}

	if refreshToken == "" {
		refreshToken = util.GenerateUUID()
		err := controller.TokenService.SaveRefreshToken(admin.ID, refreshToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while storing refresh token", nil))
			return
		}
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "login succeed", AdminResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}))
}
