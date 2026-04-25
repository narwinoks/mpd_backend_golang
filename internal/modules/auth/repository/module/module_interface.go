package module

import "backend-app/internal/modules/auth/models"

type ModuleRepository interface {
	GetUserModules(userID uint32, roleID uint32) ([]models.AppModule, error)
}
