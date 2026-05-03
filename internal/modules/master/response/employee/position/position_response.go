package position

import (
	"time"
)

type PositionResponse struct {
	ID        string    `json:"id"`
	Position  string    `json:"position"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
