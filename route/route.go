package route

import (
	"school/controller/group"
	"school/controller/message"
	"school/controller/user"
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
		userService := &service.UserService{
			Repo: &repository.UserRepository{
				DB: db.DB,
			},
		}
		messageService := &service.MessageService{
			Repo: &repository.MessageRepository{
				DB: db.DB,
			},
		}

		userController := user.UserController{
			UserService: userService,
			TokenService: &service.TokenService{
				Repo: &repository.TokenRepository{
					DB: db.DB,
				},
			},
			MessageService: messageService,
		}

		messageController := message.MessageController{
			MessageService: messageService,
		}

		groupController := group.GroupController{
			UserService: userService,
			GroupService: &service.GroupService{
				Repo: &repository.GroupRepository{
					DB: db.DB,
				},
			},
		}

		// User Authentication and Account Management (UserController)
		v1.POST("/user/login", userController.Login)
		v1.GET("/user/messages", messageController.GetMessages)  // Moved to MessageController
		v1.POST("/user/messages", messageController.ReadMessage) // Moved to MessageController
		// v1.PATCH("/user/password", userController.ChangePassword)

		// Admin Authentication and Messaging (AdminController)
		v1.POST("/admin/login", userController.AdminLogin)
		v1.POST("/admin/send-message", userController.SendMessage)

		// User Management (UserController)
		v1.POST("/admin/users", userController.NewUser)          // Create a new user
		v1.GET("/admin/users", userController.Users)             // List all users
		v1.GET("/admin/users/:id", userController.GetUser)       // View a specific user
		v1.DELETE("/admin/users/:id", userController.DeleteUser) // Delete a specific user
		v1.PATCH("/admin/users/:id", userController.ModifyUser)  // Update a specific user

		// Group Management (GroupController)
		v1.POST("/admin/groups", groupController.CreateGroup)
		v1.GET("/admin/groups", groupController.ListGroups)
		v1.GET("/admin/groups/:id", groupController.GetGroup)
		v1.DELETE("/admin/groups/:id", groupController.DeleteGroup)
		v1.PATCH("/admin/groups/:id", groupController.UpdateGroup)

		// Assigning Permissions to Groups (PermissionGroupController)
		v1.POST("/admin/groups/:id/permissions", groupController.AssignPermissionsToGroup)
		v1.GET("/admin/groups/:id/permissions", groupController.ListGroupPermissions)
		v1.DELETE("/admin/groups/:id/permissions", groupController.RemovePermissionsFromGroup)

		// Assigning Users to Groups (GroupController)
		v1.POST("/admin/groups/:id/users", groupController.AssignUserToGroup)
		v1.DELETE("/admin/groups/:id/users/:user_id", groupController.RemoveUserFromGroup)

		// Assigning Permissions Directly to Users (UserController)
		v1.POST("/admin/users/:id/permissions", userController.AssignPermissionsToUser)
		v1.DELETE("/admin/users/:id/permissions", userController.RemovePermissionsFromUser)
	}
}
