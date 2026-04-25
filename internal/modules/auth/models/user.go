package models

import (
	"backend-app/internal/base/models"
	masterModels "backend-app/internal/modules/master/model/employee"
)

type User struct {
	models.BaseModel
	RoleID     uint32           `gorm:"column:role_id" json:"role_id"`
	EmployeeID *uint32          `gorm:"column:employee_id" json:"employee_id"`
	Username   string           `gorm:"column:username;type:varchar(50)" json:"username"`
	Email      string           `gorm:"column:email;type:varchar(100)" json:"email"`
	Password   string           `gorm:"column:password;type:varchar(255)" json:"password"`
	Role       Role             `gorm:"foreignKey:RoleID;references:ID"`
	Employee   *masterModels.Employee `gorm:"foreignKey:EmployeeID;references:ID"`
}

func (User) TableName() string {
	return "users_m"
}
