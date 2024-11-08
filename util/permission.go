package util

import "school/e"

func IsPermissionValid(permissionID uint) bool {
	return IsUintArrayContains(e.ValidPermissions, permissionID)
}
