package employee

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/employee/job_category"
	"backend-app/internal/modules/master/service/employee/job_category"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type JobCategoryController struct {
	service job_category.JobCategoryService
}

func NewJobCategoryController(service job_category.JobCategoryService) *JobCategoryController {
	return &JobCategoryController{service: service}
}

func (h *JobCategoryController) FindAll(c *gin.Context) {
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

func (h *JobCategoryController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *JobCategoryController) Create(c *gin.Context) {
	var request req.CreateJobCategoryRequest

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

func (h *JobCategoryController) Update(c *gin.Context) {
	id := c.Param("id")
	var request req.UpdateJobCategoryRequest
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

func (h *JobCategoryController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
