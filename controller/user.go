package controller

import "school/service"

type UserController struct {
	UserService    *service.UserService
	TokenService   *service.TokenService
	MessageService *service.MessageService
}