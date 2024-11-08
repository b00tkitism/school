package group

import "school/service"

type GroupController struct {
	UserService  *service.UserService
	GroupService *service.GroupService
}
