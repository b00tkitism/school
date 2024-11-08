package e

var CreateUserPermission uint = 1
var ModifyUserPermission uint = 2
var DeleteUserPermission uint = 3
var SendMessagePermission uint = 4
var AccessUserInformation uint = 5
var CreateGroupPermission uint = 6
var AccessGroupInformationPermission uint = 7
var ModifyGroupPermission uint = 8
var DeleteGroupPermission uint = 9
var AssignPermissionToGroupPermission uint = 10
var RemovePermissionFromGroupPermission uint = 11

var ValidPermissions []uint = []uint{
	CreateUserPermission, ModifyUserPermission, DeleteUserPermission,
	SendMessagePermission, AccessUserInformation, CreateGroupPermission,
	AccessGroupInformationPermission, ModifyGroupPermission, DeleteGroupPermission,
	AssignPermissionToGroupPermission, RemovePermissionFromGroupPermission,
}
