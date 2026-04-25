package controller

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/core/response"
	"backend-app/internal/modules/auth/repository/user"
	"backend-app/internal/modules/auth/service/menu"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	menuService menu.MenuService
	userRepo    user.UserRepository
}

func NewMenuController(menuService menu.MenuService, userRepo user.UserRepository) *MenuController {
	return &MenuController{
		menuService: menuService,
		userRepo:    userRepo,
	}
}

func (h *MenuController) GetSideMenu(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.Error(exception.NewUnauthorizedError("User not found in context"))
		return
	}

	userID := userIDVal.(uint32)

	// We need role_id, let's fetch user from repo
	user, err := h.userRepo.FindByID(userID)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.menuService.GetSideMenu(userID, user.RoleID)
	if err != nil {
		c.Error(err)
		return
	}

	response.SendSuccess(c, response.Success, res)
}
