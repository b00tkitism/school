package middleware

import (
	"net/http"
	"school/e"
	"school/util"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !util.IsArrayContains(e.NoAuthRoutes, c.Request.RequestURI) {
			token := c.Request.Header.Get("Authorization")
			if token == "" || !strings.Contains(token, "Bearer") {
				c.JSON(http.StatusBadRequest, util.GenerateResponse(false, "no access_token provided", nil))
				c.Abort()
				return
			}
			token = strings.Split(token, "Bearer ")[1]
			userID, err := util.ParseJWT(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, util.GenerateResponse(false, "invalid token", nil))
				c.Abort()
				return
			}

			if userID != 0 {
				c.Set("user_id", userID)
				c.Next()
				return
			}
		}
		c.Next()
	}
}
