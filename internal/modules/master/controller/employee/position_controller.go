package employee

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/employee/position"
	"backend-app/internal/modules/master/service/employee/position"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type PositionController struct {
	positionService position.PositionService
}

func NewPositionController(positionService position.PositionService) *PositionController {
	return &PositionController{positionService: positionService}
}

func (h *PositionController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.positionService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *PositionController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.positionService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *PositionController) Create(c *gin.Context) {
	var positionReq req.CreatePositionRequest

	if err := c.ShouldBindJSON(&positionReq); err != nil {
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

	res, err := h.positionService.Create(ctx, positionReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *PositionController) Update(c *gin.Context) {
	id := c.Param("id")
	var positionReq req.UpdatePositionRequest
	if err := c.ShouldBindJSON(&positionReq); err != nil {
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

	res, err := h.positionService.Update(ctx, id, positionReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *PositionController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.positionService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
