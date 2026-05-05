package app_menu

import "time"

type AppMenuResponse struct {
	ID            string             `json:"id"`
	AppModuleID   string             `json:"app_module_id"`
	AppModuleName string             `json:"app_module_name"`
	ParentID      *string            `json:"parent_id,omitempty"`
	ParentName    *string            `json:"parent_name,omitempty"`
	Code          string             `json:"code"`
	Name          string             `json:"name"`
	Path          string             `json:"path"`
	Description   string             `json:"description"`
	Icon          string             `json:"icon"`
	SortOrder     int                `json:"sort_order"`
	IsActive      bool               `json:"is_active"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	SubMenus      []*AppMenuResponse `json:"sub_menus,omitempty"`
}
