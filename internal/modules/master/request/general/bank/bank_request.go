package bank

type CreateBankRequest struct {
	Bank string `json:"bank" binding:"required,max=20"`
}

type UpdateBankRequest struct {
	Bank string `json:"bank" binding:"required,max=20"`
}
