//go:build wireinject
// +build wireinject

package master

import (
	"backend-app/config"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeMasterRouter(cfg *config.Config, db *gorm.DB) *MasterRouter {
	wire.Build(
		MasterProviderSet,
	)
	return &MasterRouter{}
}
