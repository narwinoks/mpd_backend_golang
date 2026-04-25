package controller

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/core/response"
	"backend-app/internal/modules/auth/service/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) *UserController {
	return &UserController{userService: userService}
}

func (h *UserController) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.Error(exception.NewUnauthorizedError("User not found in context"))
		return
	}

	res, err := h.userService.GetProfile(userID.(uint32))
	if err != nil {
		c.Error(err)
		return
	}

	response.SendSuccess(c, response.Success, res)
}
