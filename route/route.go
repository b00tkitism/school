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
		userController := controller.UserController{
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

		v1.POST("/admin/login", userController.AdminLogin)
		v1.POST("/admin/users", userController.NewUser)
		v1.POST("/admin/send-message", userController.SendMessage)

		v1.GET("/admin/users", userController.Users)

		v1.POST("/user/login", userController.Login)
		v1.POST("/user/messages", userController.ReadMessage)
		v1.GET("/user/messages", userController.GetMessages)

	}
}
