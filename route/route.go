package route

import (
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
		messageService := &service.MessageService{
			Repo: &repository.MessageRepository{
				DB: db.DB,
			},
		}

		userController := user.UserController{
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
			MessageService: messageService,
		}

		messageController := message.MessageController{
			MessageService: messageService,
		}

		// User Authentication and Account Management (UserController)
		v1.POST("/user/login", userController.Login)
		v1.GET("/user/messages", messageController.GetMessages)  // Moved to MessageController
		v1.POST("/user/messages", messageController.ReadMessage) // Moved to MessageController

		// Admin Authentication and Messaging (AdminController)
		v1.POST("/admin/login", userController.AdminLogin)
		v1.POST("/admin/send-message", userController.SendMessage)

		// User Management (UserController)
		v1.POST("/admin/users", userController.NewUser)          // Create a new user
		v1.GET("/admin/users", userController.Users)             // List all users
		v1.GET("/admin/users/:id", userController.GetUser)       // View a specific user
		v1.DELETE("/admin/users/:id", userController.DeleteUser) // Delete a specific user
		v1.PATCH("/admin/users/:id", userController.ModifyUser)  // Update a specific user

		// Permission Group Management (PermissionGroupController)
		// v1.POST("/admin/permission-groups", permissionGroupController.CreatePermissionGroup)
		// v1.GET("/admin/permission-groups", permissionGroupController.ListPermissionGroups)
		// v1.GET("/admin/permission-groups/:id", permissionGroupController.GetPermissionGroup)
		// v1.DELETE("/admin/permission-groups/:id", permissionGroupController.DeletePermissionGroup)
		// v1.PATCH("/admin/permission-groups/:id", permissionGroupController.UpdatePermissionGroup)

		// Assigning Permissions to Permission Groups (PermissionGroupController)
		// v1.POST("/admin/permission-groups/:id/permissions", permissionGroupController.AssignPermissionToGroup)
		// v1.DELETE("/admin/permission-groups/:id/permissions/:permission_id", permissionGroupController.RemovePermissionFromGroup)

		// Group Management (GroupController)
		// v1.POST("/admin/groups", groupController.CreateGroup)
		// v1.GET("/admin/groups", groupController.ListGroups)
		// v1.GET("/admin/groups/:id", groupController.GetGroup)
		// v1.DELETE("/admin/groups/:id", groupController.DeleteGroup)
		// v1.PATCH("/admin/groups/:id", groupController.UpdateGroup)

		// Assigning Users to Groups (GroupController)
		// v1.POST("/admin/groups/:id/users", groupController.AssignUserToGroup)
		// v1.DELETE("/admin/groups/:id/users/:user_id", groupController.RemoveUserFromGroup)

		// Permission Management (PermissionController)
		// v1.POST("/admin/permissions", permissionController.CreatePermission)
		// v1.GET("/admin/permissions", permissionController.ListPermissions)
		// v1.GET("/admin/permissions/:id", permissionController.GetPermission)
		// v1.DELETE("/admin/permissions/:id", permissionController.DeletePermission)
		// v1.PATCH("/admin/permissions/:id", permissionController.UpdatePermission)

		// Assigning Permissions Directly to Users (UserController)
		// v1.POST("/admin/users/:id/permissions", userController.AssignPermissionToUser)
		// v1.DELETE("/admin/users/:id/permissions/:permission_id", userController.RemovePermissionFromUser)

	}
}
