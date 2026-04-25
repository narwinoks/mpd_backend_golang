package menu

import "backend-app/internal/modules/auth/models"

type MenuInterface interface {
	GetMenusByModuleIDs(moduleIDs []uint32) ([]models.AppMenu, error)
}
