package general

import (
	"backend-app/internal/base/models"
)

type MaritalStatus struct {
	models.BaseModel
	MaterialStatus string `gorm:"column:material_status;type:varchar(50)" json:"material_status"`
}

func (MaritalStatus) TableName() string {
	return "marital_status_m"
}
