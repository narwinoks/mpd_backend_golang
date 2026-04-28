package role

import (
	"time"
)

type RoleResponse struct {
	ID        uint32    `json:"id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
