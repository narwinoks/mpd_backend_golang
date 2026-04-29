package registry

type CreateRegistryRequest struct {
	Name      string `json:"name" binding:"required"`
	Path      string `json:"path"`
	Icon      string `json:"icon"`
	HeadID    string `json:"head_id"`
	SortOrder int    `json:"sort_order"`
}

type UpdateRegistryRequest struct {
	Name      string `json:"name" binding:"required"`
	Path      string `json:"path"`
	Icon      string `json:"icon"`
	HeadID    string `json:"head_id"`
	SortOrder int    `json:"sort_order"`
}
