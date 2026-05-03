package employee

import (
	"github.com/gin-gonic/gin"

	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/employee"
	"backend-app/internal/modules/master/service/employee"
	"backend-app/pkg/pagination"
	"context"
)

type EmployeeController struct {
	employeeService employee.EmployeeService
}

func NewEmployeeController(employeeService employee.EmployeeService) *EmployeeController {
	return &EmployeeController{employeeService: employeeService}
}

func (h *EmployeeController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.employeeService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *EmployeeController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.employeeService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *EmployeeController) Create(c *gin.Context) {
	var employeeReq req.CreateEmployeeRequest

	if err := c.ShouldBindJSON(&employeeReq); err != nil {
		c.Error(err)
		return
	}

	ctx := c.Request.Context()
	if profileID, exists := c.Get("profile_id"); exists {
		ctx = context.WithValue(ctx, "profile_id", profileID)
	}
	if userID, exists := c.Get("user_id"); exists {
		ctx = context.WithValue(ctx, "user_id", userID)
	}

	res, err := h.employeeService.Create(ctx, employeeReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, map[string]string{"id": res})
}

func (h *EmployeeController) Update(c *gin.Context) {
	id := c.Param("id")
	var employeeReq req.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&employeeReq); err != nil {
		c.Error(err)
		return
	}

	ctx := c.Request.Context()
	if profileID, exists := c.Get("profile_id"); exists {
		ctx = context.WithValue(ctx, "profile_id", profileID)
	}
	if userID, exists := c.Get("user_id"); exists {
		ctx = context.WithValue(ctx, "user_id", userID)
	}

	res, err := h.employeeService.Update(ctx, id, employeeReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, map[string]string{"id": res})
}

func (h *EmployeeController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.employeeService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
