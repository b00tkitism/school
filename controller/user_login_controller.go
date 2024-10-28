package controller

import (
	"net/http"
	"school/models"
	"school/util"

	"github.com/gin-gonic/gin"
)



type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (controller *UserController) Login(c *gin.Context) {
	var loginRequest LoginRequest
	c.BindJSON(&loginRequest)

	if loginRequest.Username == "" || loginRequest.Password == "" {
		c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "username or password is empty", nil))
		return
	}

	user, err := controller.UserService.GetUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while fetching user", nil))
		return
	}

	if (*user == models.Users{}) {
		c.JSON(http.StatusForbidden, util.GenerateResponse(false, "invalid credentials", nil))
		return
	}

	accessToken, err := util.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while generating jwt", nil))
		return
	}

	refreshToken, err := controller.TokenService.GetRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while getting refresh token", nil))
		return
	}

	if refreshToken == "" {
		refreshToken = util.GenerateUUID()
		err := controller.TokenService.SaveRefreshToken(user.ID, refreshToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "error while storing refresh token", nil))
			return
		}
	}

	c.JSON(http.StatusOK, util.GenerateResponse(true, "login succeed", LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}))
}
