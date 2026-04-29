package controller

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/role"
	"backend-app/internal/modules/master/service/role"
	"backend-app/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService role.RoleService
}

func NewRoleController(roleService role.RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}
func (h *RoleController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	roles, meta, err := h.roleService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, roles, meta)
}

func (h *RoleController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.roleService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *RoleController) Create(c *gin.Context) {
	var roleReq req.CreateRoleRequest

	if err := c.ShouldBindJSON(&roleReq); err != nil {
		c.Error(err)
		return
	}

	res, err := h.roleService.Create(c.Request.Context(), roleReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *RoleController) Update(c *gin.Context) {
	id := c.Param("id")
	var roleReq req.UpdateRoleRequest
	if err := c.ShouldBindJSON(&roleReq); err != nil {
		c.Error(err)
		return
	}

	res, err := h.roleService.Update(c.Request.Context(), id, roleReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *RoleController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.roleService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
