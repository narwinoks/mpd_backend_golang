package registry

import "time"

type RegistryResponse struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Path      string             `json:"path"`
	Icon      string             `json:"icon"`
	HeadID    *uint32            `json:"head_id"`
	SortOrder int                `json:"sort_order"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Children  []RegistryResponse `json:"children,omitempty"`
}
