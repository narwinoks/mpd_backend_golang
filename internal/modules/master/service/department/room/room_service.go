package room

import (
	req "backend-app/internal/modules/master/request/department/room"
	res "backend-app/internal/modules/master/response/department/room"
	"backend-app/pkg/pagination"
	"context"
)

type RoomService interface {
	GetAll(ctx context.Context, request pagination.BaseRequest) ([]res.RoomResponse, *pagination.Meta, error)
	GetByID(ctx context.Context, id string) (*res.RoomResponse, error)
	Create(ctx context.Context, request req.CreateRoomRequest) (*res.RoomResponse, error)
	Update(ctx context.Context, id string, request req.UpdateRoomRequest) (*res.RoomResponse, error)
	Delete(ctx context.Context, id string) error
}
