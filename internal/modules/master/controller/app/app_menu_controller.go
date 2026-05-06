package app

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/app_menu"
	"backend-app/internal/modules/master/service/app_menu"

	"github.com/gin-gonic/gin"
)

type AppMenuController struct {
	service app_menu.AppMenuService
}

func NewAppMenuController(service app_menu.AppMenuService) *AppMenuController {
	return &AppMenuController{service: service}
}

func (h *AppMenuController) FindAll(c *gin.Context) {
	var paginateReq req.AppMenuFilterRequest
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

func (h *AppMenuController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *AppMenuController) Create(c *gin.Context) {
	var request req.CreateAppMenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	res, err := h.service.Create(c.Request.Context(), request)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *AppMenuController) Update(c *gin.Context) {
	id := c.Param("id")
	var request req.UpdateAppMenuRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	res, err := h.service.Update(c.Request.Context(), id, request)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *AppMenuController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, nil)
}
