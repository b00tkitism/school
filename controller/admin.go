package controller

import "school/service"

type AdminController struct {
	UserService    *service.UserService
	TokenService   *service.TokenService
	MessageService *service.MessageService
}
