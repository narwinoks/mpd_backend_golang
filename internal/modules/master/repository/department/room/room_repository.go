package room

import (
	model "backend-app/internal/modules/master/model/department"
	"backend-app/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type roomRepositoryImpl struct{ db *gorm.DB }

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepositoryImpl{db: db}
}

type roomWithCount struct {
	model.Room
	TotalCount int64 `gorm:"column:total_count"`
}

func (r *roomRepositoryImpl) FindAll(ctx context.Context, req pagination.BaseRequest) ([]model.Room, int64, error) {
	var rows []roomWithCount
	err := r.db.WithContext(ctx).Model(&model.Room{}).
		Scopes(pagination.PaginateScope(req)).
		Scopes(req.SearchScope("room_name", "external_code", "rs_online_code")).
		Find(&rows).Error
	if err != nil {
		return nil, 0, err
	}
	var results []model.Room
	var total int64
	for _, row := range rows {
		total = row.TotalCount
		results = append(results, row.Room)
	}
	return results, total, nil
}

func (r *roomRepositoryImpl) FindByID(ctx context.Context, id uint32) (*model.Room, error) {
	var m model.Room
	err := r.db.WithContext(ctx).
		Select("id", "uuid", "room_name", "ward_id", "class_id", "rs_online_code", "bed_count", "occupied_room_count", "available_room_count", "is_active", "external_code", "created_at", "updated_at").
		First(&m, id).Error
	return &m, err
}

func (r *roomRepositoryImpl) FindByUuid(ctx context.Context, uuid string) (*model.Room, error) {
	var m model.Room
	err := r.db.WithContext(ctx).
		Select("id", "uuid", "room_name", "ward_id", "class_id", "rs_online_code", "bed_count", "occupied_room_count", "available_room_count", "is_active", "external_code", "created_at", "updated_at").
		Where("uuid = ?", uuid).First(&m).Error
	return &m, err
}

func (r *roomRepositoryImpl) Create(ctx context.Context, m *model.Room) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *roomRepositoryImpl) Update(ctx context.Context, m *model.Room) error {
	return r.db.WithContext(ctx).Updates(m).Error
}

func (r *roomRepositoryImpl) Delete(ctx context.Context, id uint32) error {
	var m model.Room
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return err
	}
	return m.SetNonActive(r.db.WithContext(ctx).Model(&m))
}
