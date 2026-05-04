package location

import (
	"context"
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/location/province"
	"backend-app/internal/modules/master/service/location/province"
	"backend-app/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type ProvinceController struct {
	provinceService province.ProvinceService
}

func NewProvinceController(provinceService province.ProvinceService) *ProvinceController {
	return &ProvinceController{provinceService: provinceService}
}

func (h *ProvinceController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.provinceService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *ProvinceController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.provinceService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *ProvinceController) Create(c *gin.Context) {
	var provinceReq req.CreateProvinceRequest

	if err := c.ShouldBindJSON(&provinceReq); err != nil {
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

	res, err := h.provinceService.Create(ctx, provinceReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *ProvinceController) Update(c *gin.Context) {
	id := c.Param("id")
	var provinceReq req.UpdateProvinceRequest
	if err := c.ShouldBindJSON(&provinceReq); err != nil {
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

	res, err := h.provinceService.Update(ctx, id, provinceReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *ProvinceController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.provinceService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
