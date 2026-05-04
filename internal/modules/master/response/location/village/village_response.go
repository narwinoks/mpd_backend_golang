package village

import (
	"backend-app/internal/modules/master/response/location/city"
	"backend-app/internal/modules/master/response/location/province"
	"backend-app/internal/modules/master/response/location/subdistrict"
	"time"
)

type VillageResponse struct {
	ID            string                          `json:"id"`
	Province      *province.ProvinceResponse      `json:"province,omitempty"`
	City          *city.CityResponse              `json:"city,omitempty"`
	Subdistrict   *subdistrict.SubdistrictResponse `json:"subdistrict,omitempty"`
	Village       string                          `json:"village"`
	PostalCode    string                          `json:"postal_code"`
	Longitude     string                          `json:"longitude"`
	Latitude      string                          `json:"latitude"`
	CreatedAt     time.Time                       `json:"created_at"`
	UpdatedAt     time.Time                       `json:"updated_at"`
}
