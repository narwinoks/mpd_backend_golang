package app_menu

import "backend-app/pkg/pagination"

type AppMenuFilterRequest struct {
	pagination.BaseRequest
	AppModuleID string `form:"app_module_id"`
	HeadID      string `form:"head_id"`
}

type CreateAppMenuRequest struct {
	AppModuleID string  `json:"app_module_id" binding:"required"`
	ParentID    *string `json:"parent_id"`
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Path        string  `json:"path" binding:"required"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	SortOrder   int     `json:"sort_order"`
}

type UpdateAppMenuRequest struct {
	AppModuleID string  `json:"app_module_id" binding:"required"`
	ParentID    *string `json:"parent_id"`
	Code        string  `json:"code" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Path        string  `json:"path" binding:"required"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	SortOrder   int     `json:"sort_order"`
	IsActive    *bool   `json:"is_active"`
}
