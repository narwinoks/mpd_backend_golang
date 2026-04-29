package permission

import "time"

type PermissionResponse struct {
	ID         string    `json:"id"`
	Permission string    `json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
