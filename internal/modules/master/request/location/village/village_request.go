package village

type CreateVillageRequest struct {
	ProvinceID    string `json:"province_id" binding:"required"`
	CityID        string `json:"city_id" binding:"required"`
	SubdistrictID string `json:"subdistrict_id" binding:"required"`
	Village       string `json:"village" binding:"required,max=100"`
	PostalCode    string `json:"postal_code" binding:"max=10"`
	Longitude     string `json:"longitude" binding:"max=50"`
	Latitude      string `json:"latitude" binding:"max=50"`
}

type UpdateVillageRequest struct {
	ProvinceID    string `json:"province_id" binding:"required"`
	CityID        string `json:"city_id" binding:"required"`
	SubdistrictID string `json:"subdistrict_id" binding:"required"`
	Village       string `json:"village" binding:"required,max=100"`
	PostalCode    string `json:"postal_code" binding:"max=10"`
	Longitude     string `json:"longitude" binding:"max=50"`
	Latitude      string `json:"latitude" binding:"max=50"`
}
