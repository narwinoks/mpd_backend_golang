package province

type CreateProvinceRequest struct {
	Province string `json:"province" binding:"required,max=100"`
}

type UpdateProvinceRequest struct {
	Province string `json:"province" binding:"required,max=100"`
}
