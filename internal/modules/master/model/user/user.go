package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Username  string         `gorm:"type:varchar(50);make;not null"`
	Email     string         `gorm:"type:varchar(100);unique;not null"`
	Password  string         `gorm:" type:varchar(255);not null"`
	FullName  string         `gorm:"type:varchar(100);not null"`
	NIP       string         `gorm:"column:nip;type:varchar(20);unique"`
	Role      string         `gorm:"type:varchar(20);not null;default:'staff'"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	LastLogin *time.Time     `json:"last_login"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
