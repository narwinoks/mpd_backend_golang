package religion

import (
	"time"
)

type ReligionResponse struct {
	ID        string    `json:"id"`
	Religion  string    `json:"religion"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
