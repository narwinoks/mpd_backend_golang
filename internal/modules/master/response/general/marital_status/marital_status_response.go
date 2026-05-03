package marital_status

import (
	"time"
)

type MaritalStatusResponse struct {
	ID            string    `json:"id"`
	MaritalStatus string    `json:"material_status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
