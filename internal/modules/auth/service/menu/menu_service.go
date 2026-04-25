package menu

import res "backend-app/internal/modules/auth/response/menu"

type MenuService interface {
	GetSideMenu(userID uint32, roleID uint32) ([]res.AppMenuResponse, error)
}
