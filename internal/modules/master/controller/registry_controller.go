package controller

import (
	"context"
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/registry"
	"backend-app/internal/modules/master/service/registry"
	"backend-app/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type RegistryController struct {
	service registry.RegistryService
}

func NewRegistryController(service registry.RegistryService) *RegistryController {
	return &RegistryController{service: service}
}

func (h *RegistryController) FindAll(c *gin.Context) {
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

func (h *RegistryController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *RegistryController) Create(c *gin.Context) {
	var request req.CreateRegistryRequest

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

func (h *RegistryController) Update(c *gin.Context) {
	id := c.Param("id")
	var request req.UpdateRegistryRequest
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

func (h *RegistryController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}

func (h *RegistryController) GetMenu(c *gin.Context) {
	res, err := h.service.GetNestedMenu(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}
