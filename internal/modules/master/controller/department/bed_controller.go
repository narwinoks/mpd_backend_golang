package department

import (
	"backend-app/internal/core/response"
	reqBed "backend-app/internal/modules/master/request/department/bed"
	svcBed "backend-app/internal/modules/master/service/department/bed"
	"backend-app/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type BedController struct {
	service svcBed.BedService
}

func NewBedController(service svcBed.BedService) *BedController {
	return &BedController{service: service}
}

func (h *BedController) FindAll(c *gin.Context) {
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

func (h *BedController) FindByID(c *gin.Context) {
	res, err := h.service.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *BedController) Create(c *gin.Context) {
	var req reqBed.CreateBedRequest
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

func (h *BedController) Update(c *gin.Context) {
	var req reqBed.UpdateBedRequest
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

func (h *BedController) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
