package city

type CreateCityRequest struct {
	ProvinceID string `json:"province_id" binding:"required"`
	Code       string `json:"code" binding:"required,max=40"`
	City       string `json:"city" binding:"required,max=100"`
}

type UpdateCityRequest struct {
	ProvinceID string `json:"province_id" binding:"required"`
	Code       string `json:"code" binding:"required,max=40"`
	City       string `json:"city" binding:"required,max=100"`
}
