package menu

import (
	"backend-app/internal/modules/auth/models"
	"backend-app/internal/modules/auth/repository/menu"
	"backend-app/internal/modules/auth/repository/module"
	res "backend-app/internal/modules/auth/response/menu"

	"github.com/sirupsen/logrus"
)

type menuServiceImpl struct {
	menuRepository   menu.MenuInterface
	moduleRepository module.ModuleRepository
}

func NewMenuService(menuRepository menu.MenuInterface, moduleRepository module.ModuleRepository) MenuService {
	return &menuServiceImpl{menuRepository: menuRepository, moduleRepository: moduleRepository}
}

func (s *menuServiceImpl) GetSideMenu(userID uint32, roleID uint32) ([]res.AppMenuResponse, error) {
	logrus.Infof("Fetching side menu for user_id: %d, role_id: %d", userID, roleID)

	modules, err := s.moduleRepository.GetUserModules(userID, roleID)
	if err != nil {
		return nil, err
	}

	if len(modules) == 0 {
		return []res.AppMenuResponse{}, nil
	}

	moduleIDs := make([]uint32, len(modules))
	for i, mod := range modules {
		moduleIDs[i] = mod.ID
	}

	allMenus, err := s.menuRepository.GetMenusByModuleIDs(moduleIDs)
	if err != nil {
		return nil, err
	}

	// Build nested menus starting from those without parents
	nestedMenus := s.buildNestedMenus(allMenus, nil)

	return nestedMenus, nil
}

func (s *menuServiceImpl) buildNestedMenus(menus []models.AppMenu, parentID *uint32) []res.AppMenuResponse {
	var result []res.AppMenuResponse

	for _, m := range menus {
		// Check if parentID matches
		isMatch := false
		if parentID == nil && m.ParentID == nil {
			isMatch = true
		} else if parentID != nil && m.ParentID != nil && *parentID == *m.ParentID {
			isMatch = true
		}

		if isMatch {
			children := s.buildNestedMenus(menus, &m.ID)

			result = append(result, res.AppMenuResponse{
				ID:          m.ID,
				Code:        m.Code,
				Name:        m.Name,
				Path:        m.Path,
				Icon:        m.Icon,
				Description: m.Description,
				SortOrder:   m.SortOrder,
				Children:    children,
			})
		}
	}

	return result
}
