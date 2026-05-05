package city

import (
	"backend-app/internal/modules/master/response/location/province"
	"time"
)

type CityResponse struct {
	ID        string                     `json:"id"`
	Code      string                     `json:"code"`
	Province  *province.ProvinceResponse `json:"province,omitempty"`
	City      string                     `json:"city"`
	CreatedAt time.Time                  `json:"created_at"`
	UpdatedAt time.Time                  `json:"updated_at"`
}
