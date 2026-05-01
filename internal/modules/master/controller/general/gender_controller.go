package general

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/general/gender"
	"backend-app/internal/modules/master/service/general/gender"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type GenderController struct {
	genderService gender.GenderService
}

func NewGenderController(genderService gender.GenderService) *GenderController {
	return &GenderController{genderService: genderService}
}

func (h *GenderController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	genders, meta, err := h.genderService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, genders, meta)
}

func (h *GenderController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.genderService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *GenderController) Create(c *gin.Context) {
	var genderReq req.CreateGenderRequest

	if err := c.ShouldBindJSON(&genderReq); err != nil {
		c.Error(err)
		return
	}

	// Ensure profile_id and user_id are in the context for GORM hooks
	ctx := c.Request.Context()
	if profileID, exists := c.Get("profile_id"); exists {
		ctx = context.WithValue(ctx, "profile_id", profileID)
	}
	if userID, exists := c.Get("user_id"); exists {
		ctx = context.WithValue(ctx, "user_id", userID)
	}

	res, err := h.genderService.Create(ctx, genderReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *GenderController) Update(c *gin.Context) {
	id := c.Param("id")
	var genderReq req.UpdateGenderRequest
	if err := c.ShouldBindJSON(&genderReq); err != nil {
		c.Error(err)
		return
	}

	// Ensure profile_id and user_id are in the context for GORM hooks
	ctx := c.Request.Context()
	if profileID, exists := c.Get("profile_id"); exists {
		ctx = context.WithValue(ctx, "profile_id", profileID)
	}
	if userID, exists := c.Get("user_id"); exists {
		ctx = context.WithValue(ctx, "user_id", userID)
	}

	res, err := h.genderService.Update(ctx, id, genderReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *GenderController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.genderService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
