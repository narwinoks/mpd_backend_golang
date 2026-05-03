package marital_status

import (
	"backend-app/internal/modules/master/model/general"
	"backend-app/pkg/pagination"
	"context"
)

type MaritalStatusRepository interface {
	FindAll(ctx context.Context, pagination pagination.BaseRequest) ([]general.MaritalStatus, int64, error)
	FindByID(ctx context.Context, id uint32) (*general.MaritalStatus, error)
	Create(ctx context.Context, maritalStatus *general.MaritalStatus) error
	Update(ctx context.Context, maritalStatus *general.MaritalStatus) error
	Delete(ctx context.Context, id uint32) error
	FindByUuid(ctx context.Context, Uuid string) (*general.MaritalStatus, error)
}
