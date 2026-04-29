package permission

import (
	"backend-app/internal/base/models"
)

type Permission struct {
	models.BaseModel
	Permission string `gorm:"column:permission;type:varchar(100)" json:"permission"`
}

func (Permission) TableName() string {
	return "app_permission_m"
}
