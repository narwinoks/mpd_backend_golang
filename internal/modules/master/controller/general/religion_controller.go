package general

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/general/religion"
	"backend-app/internal/modules/master/service/general/religion"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type ReligionController struct {
	religionService religion.ReligionService
}

func NewReligionController(religionService religion.ReligionService) *ReligionController {
	return &ReligionController{religionService: religionService}
}

func (h *ReligionController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	religions, meta, err := h.religionService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, religions, meta)
}

func (h *ReligionController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.religionService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *ReligionController) Create(c *gin.Context) {
	var relReq req.CreateReligionRequest

	if err := c.ShouldBindJSON(&relReq); err != nil {
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

	res, err := h.religionService.Create(ctx, relReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *ReligionController) Update(c *gin.Context) {
	id := c.Param("id")
	var relReq req.UpdateReligionRequest
	if err := c.ShouldBindJSON(&relReq); err != nil {
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

	res, err := h.religionService.Update(ctx, id, relReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *ReligionController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.religionService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
