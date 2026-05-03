package general

import (
	"backend-app/internal/base/models"
)

type Bank struct {
	models.BaseModel
	Bank string `gorm:"column:bank;type:varchar(20)" json:"bank"`
}

func (Bank) TableName() string {
	return "banks_m"
}
