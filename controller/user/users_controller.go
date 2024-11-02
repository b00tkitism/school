package user

import (
	"net/http"
	"school/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersList struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

type UsersListResponse struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Users    []UsersList `json:"users"`
	Total    int64       `json:"total"`
}

func (controller *UserController) Users(c *gin.Context) {
	// Get page and page size from query parameters, with default values
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// Calculate the offset for pagination
	offset := (page - 1) * pageSize

	// Fetch users with pagination from UserService
	users, err := controller.UserService.GetPaginatedUsers(offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.GenerateResponse(false, "failed to fetch users", nil))
		return
	}

	// Build response with user ID, Username, and FullName only
	var userResponses []UsersList
	for _, user := range users {
		userResponses = append(userResponses, UsersList{
			ID:       user.ID,
			Username: user.Username,
			IsAdmin:  user.IsAdmin,
		})
	}

	count, _ := controller.UserService.GetUsersCount()

	c.JSON(http.StatusOK, util.GenerateResponse(true, "users fetched", UsersListResponse{
		Page:     page,
		PageSize: pageSize,
		Users:    userResponses,
		Total:    count,
	}))
}
