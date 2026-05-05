package subdistrict

type CreateSubdistrictRequest struct {
	ProvinceID  string `json:"province_id" binding:"required"`
	CityID      string `json:"city_id" binding:"required"`
	Code        string `json:"code" binding:"required,max=40"`
	Subdistrict string `json:"subdistrict" binding:"required,max=100"`
}

type UpdateSubdistrictRequest struct {
	ProvinceID  string `json:"province_id" binding:"required"`
	CityID      string `json:"city_id" binding:"required"`
	Code        string `json:"code" binding:"required,max=40"`
	Subdistrict string `json:"subdistrict" binding:"required,max=100"`
}
