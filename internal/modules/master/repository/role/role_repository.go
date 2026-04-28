package role

import (
	"backend-app/internal/modules/master/model/role"
	"backend-app/pkg/pagination"

	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepositoryImpl{db: db}
}

func (r *roleRepositoryImpl) FindAll(req pagination.Request) ([]role.Role, int64, error) {
	var roles []role.Role
	var total int64

	db := r.db.Model(&role.Role{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Scopes(pagination.PaginateScope(req)).
		Select("id", "role", "created_at", "updated_at").
		Find(&roles).Error

	return roles, total, err
}

func (r *roleRepositoryImpl) FindByID(id uint32) (*role.Role, error) {
	var roleEntity role.Role
	err := r.db.Select("id", "role", "created_at", "updated_at").First(&roleEntity, id).Error
	if err != nil {
		return nil, err
	}
	return &roleEntity, nil
}

func (r *roleRepositoryImpl) Create(roleEntity *role.Role) error {
	return r.db.Create(roleEntity).Error
}

func (r *roleRepositoryImpl) Update(roleEntity *role.Role) error {
	return r.db.Save(roleEntity).Error
}

func (r *roleRepositoryImpl) Delete(id uint32) error {
	return r.db.Delete(&role.Role{}, id).Error
}
