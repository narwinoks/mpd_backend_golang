package room

import (
	"backend-app/internal/modules/master/model/department"
	"backend-app/pkg/pagination"
	"context"
)

type RoomRepository interface {
	FindAll(ctx context.Context, req pagination.BaseRequest) ([]department.Room, int64, error)
	FindByID(ctx context.Context, id uint32) (*department.Room, error)
	FindByUuid(ctx context.Context, uuid string) (*department.Room, error)
	Create(ctx context.Context, m *department.Room) error
	Update(ctx context.Context, m *department.Room) error
	Delete(ctx context.Context, id uint32) error
}
