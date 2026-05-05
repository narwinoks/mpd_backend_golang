package general

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/general/marital_status"
	"backend-app/internal/modules/master/service/general/marital_status"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type MaritalStatusController struct {
	maritalStatusService marital_status.MaritalStatusService
}

func NewMaritalStatusController(maritalStatusService marital_status.MaritalStatusService) *MaritalStatusController {
	return &MaritalStatusController{maritalStatusService: maritalStatusService}
}

func (h *MaritalStatusController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.maritalStatusService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *MaritalStatusController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.maritalStatusService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *MaritalStatusController) Create(c *gin.Context) {
	var maritalStatusReq req.CreateMaritalStatusRequest

	if err := c.ShouldBindJSON(&maritalStatusReq); err != nil {
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

	res, err := h.maritalStatusService.Create(ctx, maritalStatusReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *MaritalStatusController) Update(c *gin.Context) {
	id := c.Param("id")
	var maritalStatusReq req.UpdateMaritalStatusRequest
	if err := c.ShouldBindJSON(&maritalStatusReq); err != nil {
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

	res, err := h.maritalStatusService.Update(ctx, id, maritalStatusReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *MaritalStatusController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.maritalStatusService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
