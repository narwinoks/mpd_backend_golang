package app_module

type CreateAppModuleRequest struct {
	Code      string `json:"code" binding:"required,max=20"`
	Name      string `json:"name" binding:"required,max=100"`
	Category  string `json:"category" binding:"required,max=50"`
	SortOrder int    `json:"sort_order"`
}

type UpdateAppModuleRequest struct {
	Code      string `json:"code" binding:"required,max=20"`
	Name      string `json:"name" binding:"required,max=100"`
	Category  string `json:"category" binding:"required,max=50"`
	SortOrder int    `json:"sort_order"`
}
