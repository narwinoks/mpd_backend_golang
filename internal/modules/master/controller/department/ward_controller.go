package department

import (
	"backend-app/internal/core/response"
	reqWard "backend-app/internal/modules/master/request/department/ward"
	svcWard "backend-app/internal/modules/master/service/department/ward"
	"backend-app/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type WardController struct {
	service svcWard.WardService
}

func NewWardController(service svcWard.WardService) *WardController {
	return &WardController{service: service}
}

func (h *WardController) FindAll(c *gin.Context) {
	var req pagination.BaseRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.Error(err)
		return
	}
	data, meta, err := h.service.GetAll(c.Request.Context(), req)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, data, meta)
}

func (h *WardController) FindByID(c *gin.Context) {
	res, err := h.service.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *WardController) Create(c *gin.Context) {
	var req reqWard.CreateWardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.service.Create(enrichContext(c), req)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *WardController) Update(c *gin.Context) {
	var req reqWard.UpdateWardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	res, err := h.service.Update(enrichContext(c), c.Param("id"), req)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *WardController) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
