package position

import (
	"backend-app/internal/modules/master/model/job"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type positionRepositoryImpl struct {
	db *gorm.DB
}

func NewPositionRepository(db *gorm.DB) PositionRepository {
	return &positionRepositoryImpl{db: db}
}

type PositionWithCount struct {
	job.Position
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *positionRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]job.Position, int64, error) {
	var results []PositionWithCount
	var positions []job.Position
	var total int64

	err := r.db.WithContext(ctx).Model(&job.Position{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("position")).
		Find(&results).Error

	if err != nil {
		return nil, 0, err
	}

	if len(results) > 0 {
		total = results[0].TotalCount
		for _, res := range results {
			positions = append(positions, res.Position)
		}
	}

	return positions, total, nil
}

func (r *positionRepositoryImpl) FindByID(ctx context.Context, id uint32) (*job.Position, error) {
	var position job.Position
	err := r.db.WithContext(ctx).Select("id", "uuid", "position", "is_active", "created_at", "updated_at").First(&position, id).Error
	if err != nil {
		return nil, err
	}
	return &position, nil
}

func (r *positionRepositoryImpl) Create(ctx context.Context, position *job.Position) error {
	return r.db.WithContext(ctx).Create(position).Error
}

func (r *positionRepositoryImpl) Update(ctx context.Context, position *job.Position) error {
	return r.db.WithContext(ctx).Updates(position).Error
}

func (r *positionRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var position job.Position
	if err := r.db.WithContext(ctx).First(&position, id).Error; err != nil {
		return err
	}
	return position.SetNonActive(r.db.WithContext(ctx).Model(&position))
}

func (r *positionRepositoryImpl) FindByUuid(ctx context.Context, Uuid string) (*job.Position, error) {
	var position job.Position
	err := r.db.WithContext(ctx).Select("id", "uuid", "position", "is_active", "created_at", "updated_at").Where("uuid = ?", Uuid).First(&position).Error
	if err != nil {
		return nil, err
	}
	return &position, nil
}
