package gender

import (
	"time"
)

type GenderResponse struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
