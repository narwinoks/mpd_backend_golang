package controller

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/role"
	"backend-app/internal/modules/master/service/role"
	"backend-app/pkg/pagination"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleService role.RoleService
}

func NewRoleController(roleService role.RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

func (h *RoleController) FindAll(c *gin.Context) {
	var paginateReq pagination.Request
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	roles, meta, err := h.roleService.GetAll(paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, roles, meta)
}

func (h *RoleController) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.roleService.GetByID(uint32(id))
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

	res, err := h.roleService.Create(roleReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *RoleController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	var roleReq req.UpdateRoleRequest
	if err := c.ShouldBindJSON(&roleReq); err != nil {
		c.Error(err)
		return
	}

	res, err := h.roleService.Update(uint32(id), roleReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *RoleController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	err = h.roleService.Delete(uint32(id))
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, nil)
}
