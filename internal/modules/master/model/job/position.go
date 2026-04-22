package job

import (
	"backend-app/internal/base/models"
)

type Position struct {
	models.BaseModel
	Position string `gorm:"column:position;type:varchar(100)" json:"position"`
}

func (Position) TableName() string {
	return "positions_m"
}
