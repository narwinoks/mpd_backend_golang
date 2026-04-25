package controller

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/auth/request/user"
	"backend-app/internal/modules/auth/service/user"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService user.UserService
}

func NewAuthController(userService user.UserService) *AuthController {
	return &AuthController{userService: userService}
}

func (h *AuthController) Login(c *gin.Context) {
	var loginReq req.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.Error(err)
		return
	}

	res, err := h.userService.Login(&loginReq)
	if err != nil {
		c.Error(err)
		return
	}

	response.SendSuccess(c, response.Success, res)
}

func (h *AuthController) RefreshToken(c *gin.Context) {
	var refreshReq req.RefreshTokenRequest

	if err := c.ShouldBindJSON(&refreshReq); err != nil {
		c.Error(err)
		return
	}

	res, err := h.userService.RefreshToken(&refreshReq)
	if err != nil {
		c.Error(err)
		return
	}

	response.SendSuccess(c, response.Success, res)
}

func (h *AuthController) Logout(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.Error(exception.NewUnauthorizedError("User not found in context"))
		return
	}

	err := h.userService.Logout(userID.(uint32))
	if err != nil {
		c.Error(err)
		return
	}

	response.SendSuccess(c, response.Success, nil)
}
