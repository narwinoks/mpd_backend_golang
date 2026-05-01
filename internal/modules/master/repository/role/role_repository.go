package role

import (
	"backend-app/internal/modules/master/model/role"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type roleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepositoryImpl{db: db}
}

type RoleWithCount struct {
	role.Role
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *roleRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]role.Role, int64, error) {
	var results []RoleWithCount
	var roles []role.Role
	var total int64

	err := r.db.WithContext(ctx).Model(&role.Role{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("role")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			roles = append(roles, res.Role)
		}
	}

	return roles, total, nil
}

func (r *roleRepositoryImpl) FindByID(ctx context.Context, id uint32) (*role.Role, error) {
	var roleEntity role.Role
	err := r.db.WithContext(ctx).Select("id", "role", "is_active", "created_at", "updated_at").First(&roleEntity, id).Error
	if err != nil {
		return nil, err
	}
	return &roleEntity, nil
}

func (r *roleRepositoryImpl) Create(ctx context.Context, roleEntity *role.Role) error {
	return r.db.WithContext(ctx).Create(roleEntity).Error
}

func (r *roleRepositoryImpl) Update(ctx context.Context, roleEntity *role.Role) error {
	return r.db.WithContext(ctx).Updates(roleEntity).Error
}

func (r *roleRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var roleEntity role.Role
	if err := r.db.WithContext(ctx).First(&roleEntity, id).Error; err != nil {
		return err
	}
	return roleEntity.SetNonActive(r.db.WithContext(ctx).Model(&roleEntity))
}

func (r *roleRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*role.Role, error) {
	var roleEntity role.Role
	err := r.db.WithContext(ctx).Select("id", "uuid", "role", "is_active", "created_at", "updated_at").Where("uuid = ?", Uuid).First(&roleEntity).Error
	if err != nil {
		return nil, err
	}
	return &roleEntity, nil
}
