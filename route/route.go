package route

import (
	"school/controller"
	"school/db"
	"school/middleware"
	"school/repository"
	"school/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/api/v1")
	v1.Use(middleware.JWTMiddleware())
	{
		adminController := controller.AdminController{
			UserService: &service.UserService{
				Repo: &repository.UserRepository{
					DB: db.DB,
				},
			},
			TokenService: &service.TokenService{
				Repo: &repository.TokenRepository{
					DB: db.DB,
				},
			},
			MessageService: &service.MessageService{
				Repo: &repository.MessageRepository{
					DB: db.DB,
				},
			},
		}

		v1.POST("/admin/login", adminController.Login)
		v1.POST("/admin/users", adminController.NewUser)
		v1.POST("/admin/send-message", adminController.SendMessage)

		v1.GET("/admin/users", adminController.Users)

		v1.POST("/user/login")
		v1.GET("/user/messages")
	}
}
