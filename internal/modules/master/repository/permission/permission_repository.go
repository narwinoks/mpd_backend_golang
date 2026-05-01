package permission

import (
	"backend-app/internal/modules/master/model/permission"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type permissionRepositoryImpl struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepositoryImpl{db: db}
}

type PermissionWithCount struct {
	permission.Permission
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *permissionRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]permission.Permission, int64, error) {
	var results []PermissionWithCount
	var permissions []permission.Permission
	var total int64

	err := r.db.WithContext(ctx).Model(&permission.Permission{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("permission")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			permissions = append(permissions, res.Permission)
		}
	}

	return permissions, total, nil
}

func (r *permissionRepositoryImpl) FindByID(ctx context.Context, id uint32) (*permission.Permission, error) {
	var p permission.Permission
	err := r.db.WithContext(ctx).First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *permissionRepositoryImpl) Create(ctx context.Context, p *permission.Permission) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *permissionRepositoryImpl) Update(ctx context.Context, p *permission.Permission) error {
	return r.db.WithContext(ctx).Updates(p).Error
}

func (r *permissionRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var p permission.Permission
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return err
	}
	return p.SetNonActive(r.db.WithContext(ctx).Model(&p))
}

func (r *permissionRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*permission.Permission, error) {
	var p permission.Permission
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}
