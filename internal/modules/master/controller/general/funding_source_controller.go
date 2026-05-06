package general

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/funding_source"
	"backend-app/internal/modules/master/service/funding_source"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type FundingSourceController struct {
	fundingSourceService funding_source.FundingSourceService
}

func NewFundingSourceController(fundingSourceService funding_source.FundingSourceService) *FundingSourceController {
	return &FundingSourceController{fundingSourceService: fundingSourceService}
}

func (h *FundingSourceController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	fundingSources, meta, err := h.fundingSourceService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, fundingSources, meta)
}

func (h *FundingSourceController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.fundingSourceService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *FundingSourceController) Create(c *gin.Context) {
	var fsReq req.CreateFundingSourceRequest

	if err := c.ShouldBindJSON(&fsReq); err != nil {
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

	res, err := h.fundingSourceService.Create(ctx, fsReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *FundingSourceController) Update(c *gin.Context) {
	id := c.Param("id")
	var fsReq req.UpdateFundingSourceRequest
	if err := c.ShouldBindJSON(&fsReq); err != nil {
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

	res, err := h.fundingSourceService.Update(ctx, id, fsReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *FundingSourceController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.fundingSourceService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
