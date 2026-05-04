package app

import (
	"context"
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/app/app_module"
	"backend-app/internal/modules/master/service/app/app_module"
	"backend-app/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type AppModuleController struct {
	appModuleService app_module.AppModuleService
}

func NewAppModuleController(appModuleService app_module.AppModuleService) *AppModuleController {
	return &AppModuleController{appModuleService: appModuleService}
}

func (h *AppModuleController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.appModuleService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *AppModuleController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.appModuleService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *AppModuleController) Create(c *gin.Context) {
	var appModuleReq req.CreateAppModuleRequest

	if err := c.ShouldBindJSON(&appModuleReq); err != nil {
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

	res, err := h.appModuleService.Create(ctx, appModuleReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, map[string]string{"id": res})
}

func (h *AppModuleController) Update(c *gin.Context) {
	id := c.Param("id")
	var appModuleReq req.UpdateAppModuleRequest
	if err := c.ShouldBindJSON(&appModuleReq); err != nil {
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

	res, err := h.appModuleService.Update(ctx, id, appModuleReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, map[string]string{"id": res})
}

func (h *AppModuleController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.appModuleService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
