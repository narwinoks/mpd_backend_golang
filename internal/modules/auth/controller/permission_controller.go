package controller

import (
	"backend-app/internal/core/exception"
	"backend-app/internal/core/response"
	"backend-app/internal/modules/auth/service/permission"

	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	permissionService permission.PermissionService
}

func NewPermissionController(permissionService permission.PermissionService) *PermissionController {
	return &PermissionController{
		permissionService: permissionService,
	}
}

func (h *PermissionController) GetUserPermissions(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.Error(exception.NewUnauthorizedError("User not found in context"))
		return
	}

	userID := userIDVal.(uint32)

	res, err := h.permissionService.GetUserPermissions(userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.SendSuccess(c, response.Success, res)
}
