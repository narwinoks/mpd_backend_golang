package province

import (
	"time"
)

type ProvinceResponse struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Province  string    `json:"province"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
