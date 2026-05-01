package employee

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/employee/job_title"
	"backend-app/internal/modules/master/service/employee/job_title"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type JobTitleController struct {
	service job_title.JobTitleService
}

func NewJobTitleController(service job_title.JobTitleService) *JobTitleController {
	return &JobTitleController{service: service}
}

func (h *JobTitleController) FindAll(c *gin.Context) {
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

func (h *JobTitleController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *JobTitleController) Create(c *gin.Context) {
	var request req.CreateJobTitleRequest

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

func (h *JobTitleController) Update(c *gin.Context) {
	id := c.Param("id")
	var request req.UpdateJobTitleRequest
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

func (h *JobTitleController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
