package department

import (
	"backend-app/internal/core/response"
	reqRoom "backend-app/internal/modules/master/request/department/room"
	svcRoom "backend-app/internal/modules/master/service/department/room"
	"backend-app/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	service svcRoom.RoomService
}

func NewRoomController(service svcRoom.RoomService) *RoomController {
	return &RoomController{service: service}
}

func (h *RoomController) FindAll(c *gin.Context) {
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

func (h *RoomController) FindByID(c *gin.Context) {
	res, err := h.service.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *RoomController) Create(c *gin.Context) {
	var req reqRoom.CreateRoomRequest
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

func (h *RoomController) Update(c *gin.Context) {
	var req reqRoom.UpdateRoomRequest
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

func (h *RoomController) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
