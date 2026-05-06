package app_menu

import (
	resAppModule "backend-app/internal/modules/master/response/app/app_module"
	"time"
)

type AppMenuResponse struct {
	ID            uint32                          `json:"id"`
	UUID          string                          `json:"uuid"`
	AppModuleID   string                          `json:"app_module_id,omitempty"`
	AppModuleName string                          `json:"app_module_name,omitempty"`
	AppModule     *resAppModule.AppModuleResponse `json:"app_module,omitempty"`
	ParentID      *string                         `json:"parent_id,omitempty"`
	ParentName    *string                         `json:"parent_name,omitempty"`
	Parent        *AppMenuResponse                `json:"parent,omitempty"`
	Code          string                          `json:"code"`
	Name          string                          `json:"name"`
	Path          string                          `json:"path"`
	Description   string                          `json:"description"`
	Icon          string                          `json:"icon"`
	SortOrder     int                             `json:"sort_order"`
	IsActive      bool                            `json:"is_active"`
	CreatedAt     time.Time                       `json:"created_at"`
	UpdatedAt     time.Time                       `json:"updated_at"`
	Children      []*AppMenuResponse              `json:"children,omitempty"`
}
