package department

import (
	"backend-app/internal/core/response"
	reqDept "backend-app/internal/modules/master/request/department/department"
	svcDept "backend-app/internal/modules/master/service/department/department"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	service svcDept.DepartmentService
}

func NewDepartmentController(service svcDept.DepartmentService) *DepartmentController {
	return &DepartmentController{service: service}
}

func (h *DepartmentController) FindAll(c *gin.Context) {
	var req pagination.BaseRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		return
	}
	data, meta, err := h.service.GetAll(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, data, meta)
}

func (h *DepartmentController) FindByID(c *gin.Context) {
	res, err := h.service.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *DepartmentController) Create(c *gin.Context) {
	var req reqDept.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.service.Create(enrichContext(c), req)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *DepartmentController) Update(c *gin.Context) {
	var req reqDept.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.service.Update(enrichContext(c), c.Param("id"), req)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *DepartmentController) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}

func enrichContext(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	if v, exists := c.Get("profile_id"); exists {
		ctx = context.WithValue(ctx, "profile_id", v)
	}
	if v, exists := c.Get("user_id"); exists {
		ctx = context.WithValue(ctx, "user_id", v)
	}
	return ctx
}
