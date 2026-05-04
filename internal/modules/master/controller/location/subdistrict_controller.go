package location

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/location/subdistrict"
	"backend-app/internal/modules/master/service/location/subdistrict"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type SubdistrictController struct {
	subdistrictService subdistrict.SubdistrictService
}

func NewSubdistrictController(subdistrictService subdistrict.SubdistrictService) *SubdistrictController {
	return &SubdistrictController{subdistrictService: subdistrictService}
}

func (h *SubdistrictController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.subdistrictService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *SubdistrictController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.subdistrictService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *SubdistrictController) Create(c *gin.Context) {
	var subdistrictReq req.CreateSubdistrictRequest

	if err := c.ShouldBindJSON(&subdistrictReq); err != nil {
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

	res, err := h.subdistrictService.Create(ctx, subdistrictReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, map[string]string{"id": res})
}

func (h *SubdistrictController) Update(c *gin.Context) {
	id := c.Param("id")
	var subdistrictReq req.UpdateSubdistrictRequest
	if err := c.ShouldBindJSON(&subdistrictReq); err != nil {
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

	res, err := h.subdistrictService.Update(ctx, id, subdistrictReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, map[string]string{"id": res})
}

func (h *SubdistrictController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.subdistrictService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
