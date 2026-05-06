package location

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/location/city"
	"backend-app/internal/modules/master/service/location/city"
	"context"

	"github.com/gin-gonic/gin"
)

type CityController struct {
	cityService city.CityService
}

func NewCityController(cityService city.CityService) *CityController {
	return &CityController{cityService: cityService}
}

func (h *CityController) FindAll(c *gin.Context) {
	var findAllReq req.FindAllRequest
	if err := c.ShouldBindQuery(&findAllReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.cityService.GetAll(c.Request.Context(), findAllReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *CityController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.cityService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *CityController) Create(c *gin.Context) {
	var cityReq req.CreateCityRequest

	if err := c.ShouldBindJSON(&cityReq); err != nil {
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

	res, err := h.cityService.Create(ctx, cityReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, map[string]string{"id": res})
}

func (h *CityController) Update(c *gin.Context) {
	id := c.Param("id")
	var cityReq req.UpdateCityRequest
	if err := c.ShouldBindJSON(&cityReq); err != nil {
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

	res, err := h.cityService.Update(ctx, id, cityReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, map[string]string{"id": res})
}

func (h *CityController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.cityService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
