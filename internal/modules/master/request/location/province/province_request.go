package province

type CreateProvinceRequest struct {
	Code     string `json:"code" binding:"required,max=40"`
	Province string `json:"province" binding:"required,max=100"`
}

type UpdateProvinceRequest struct {
	Code     string `json:"code" binding:"required,max=40"`
	Province string `json:"province" binding:"required,max=100"`
}
