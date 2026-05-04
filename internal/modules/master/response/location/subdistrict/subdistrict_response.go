package subdistrict

import (
	"backend-app/internal/modules/master/response/location/city"
	"backend-app/internal/modules/master/response/location/province"
	"time"
)

type SubdistrictResponse struct {
	ID          string                    `json:"id"`
	Province    *province.ProvinceResponse `json:"province,omitempty"`
	City        *city.CityResponse         `json:"city,omitempty"`
	Subdistrict string                    `json:"subdistrict"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
}
