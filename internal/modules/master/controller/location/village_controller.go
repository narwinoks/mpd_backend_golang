package location

import (
	"context"
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/location/village"
	"backend-app/internal/modules/master/service/location/village"
	"backend-app/pkg/pagination"

	"github.com/gin-gonic/gin"
)

type VillageController struct {
	villageService village.VillageService
}

func NewVillageController(villageService village.VillageService) *VillageController {
	return &VillageController{villageService: villageService}
}

func (h *VillageController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.villageService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *VillageController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.villageService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *VillageController) Create(c *gin.Context) {
	var villageReq req.CreateVillageRequest

	if err := c.ShouldBindJSON(&villageReq); err != nil {
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

	res, err := h.villageService.Create(ctx, villageReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, map[string]string{"id": res})
}

func (h *VillageController) Update(c *gin.Context) {
	id := c.Param("id")
	var villageReq req.UpdateVillageRequest
	if err := c.ShouldBindJSON(&villageReq); err != nil {
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

	res, err := h.villageService.Update(ctx, id, villageReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, map[string]string{"id": res})
}

func (h *VillageController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.villageService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
