package general

import (
	"backend-app/internal/core/response"
	req "backend-app/internal/modules/master/request/general/bank"
	"backend-app/internal/modules/master/service/general/bank"
	"backend-app/pkg/pagination"
	"context"

	"github.com/gin-gonic/gin"
)

type BankController struct {
	bankService bank.BankService
}

func NewBankController(bankService bank.BankService) *BankController {
	return &BankController{bankService: bankService}
}

func (h *BankController) FindAll(c *gin.Context) {
	var paginateReq pagination.BaseRequest
	if err := c.ShouldBindQuery(&paginateReq); err != nil {
		c.Error(err)
		return
	}

	items, meta, err := h.bankService.GetAll(c.Request.Context(), paginateReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, items, meta)
}

func (h *BankController) FindByID(c *gin.Context) {
	id := c.Param("id")

	res, err := h.bankService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.Success, res)
}

func (h *BankController) Create(c *gin.Context) {
	var bankReq req.CreateBankRequest

	if err := c.ShouldBindJSON(&bankReq); err != nil {
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

	res, err := h.bankService.Create(ctx, bankReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessCreate, res)
}

func (h *BankController) Update(c *gin.Context) {
	id := c.Param("id")
	var bankReq req.UpdateBankRequest
	if err := c.ShouldBindJSON(&bankReq); err != nil {
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

	res, err := h.bankService.Update(ctx, id, bankReq)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessUpdate, res)
}

func (h *BankController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.bankService.Delete(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	response.SendSuccess(c, response.SuccessDelete, "")
}
