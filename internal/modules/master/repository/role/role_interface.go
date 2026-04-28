package role

import (
	"backend-app/internal/modules/master/model/role"
	"backend-app/pkg/pagination"
)

type RoleRepository interface {
	FindAll(pagination pagination.Request) ([]role.Role, int64, error)
	FindByID(id uint32) (*role.Role, error)
	Create(role *role.Role) error
	Update(role *role.Role) error
	Delete(id uint32) error
}
