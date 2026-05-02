package employee

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/employee/employment_status"
	"backend-app/internal/modules/master/service/employee/employment_status"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type EmploymentStatusController struct {
	service employment_status.EmploymentStatusService
}

func NewEmploymentStatusController(service employment_status.EmploymentStatusService) *EmploymentStatusController {
	return &EmploymentStatusController{service: service}
}

func (h *EmploymentStatusController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	res, meta, err := h.service.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res, meta)
}

func (h *EmploymentStatusController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *EmploymentStatusController) Create(c *gin.Context) {
	var request req.CreateEmploymentStatusRequest

	if err := c.ShouldBindJSON(&request); err != nil {
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

	res, err := h.service.Create(ctx, request)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *EmploymentStatusController) Update(c *gin.Context) {
	id := c.Param("id")
	var request req.UpdateEmploymentStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
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

	res, err := h.service.Update(ctx, id, request)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *EmploymentStatusController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
