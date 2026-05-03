package general

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/general/education"
	"backend-app/internal/modules/master/service/general/education"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type EducationController struct {
	educationService education.EducationService
}

func NewEducationController(educationService education.EducationService) *EducationController {
	return &EducationController{educationService: educationService}
}

func (h *EducationController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.educationService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *EducationController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.educationService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *EducationController) Create(c *gin.Context) {
	var educationReq req.CreateEducationRequest

	if err := c.ShouldBindJSON(&educationReq); err != nil {
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

	res, err := h.educationService.Create(ctx, educationReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *EducationController) Update(c *gin.Context) {
	id := c.Param("id")
	var educationReq req.UpdateEducationRequest
	if err := c.ShouldBindJSON(&educationReq); err != nil {
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

	res, err := h.educationService.Update(ctx, id, educationReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *EducationController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.educationService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
