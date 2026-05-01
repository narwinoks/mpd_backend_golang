package controller

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/user"
	"backend-app/internal/modules/master/service/user"
	"backend-app/pkg/pagination"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) *UserController {
	return &UserController{userService: userService}
}

func (h *UserController) FindAll(c *gin.Context) {
	var request pagination.BaseRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.Error(err)
		return
	}

	users, meta, err := h.userService.GetAllUsers(c.Request.Context(), request)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, users, meta)
}

func (h *UserController) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	res, err := h.userService.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *UserController) Create(c *gin.Context) {
	var userReq req.UserCreateRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.Error(err)
		return
	}

	res, err := h.userService.CreateUser(c.Request.Context(), &userReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}
