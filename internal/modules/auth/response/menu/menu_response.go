package menu

type AppMenuResponse struct {
	ID          uint32            `json:"id"`
	Code        string            `json:"code"`
	Name        string            `json:"name"`
	Path        string            `json:"path"`
	Icon        string            `json:"icon"`
	Description string            `json:"description"`
	SortOrder   int               `json:"sort_order"`
	Children    []AppMenuResponse `json:"children,omitempty"`
}
