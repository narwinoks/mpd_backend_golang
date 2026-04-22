package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID           uint32         `gorm:"column:id;primaryKey" json:"id"`
	UUID         string         `gorm:"column:uuid;type:char(36);uniqueIndex" json:"uuid"`
	IsActive     bool           `gorm:"column:is_active;default:true" json:"is_active"`
	ProfileID    *uint32        `gorm:"column:profile_id" json:"profile_id"`
	ExternalCode string         `gorm:"column:external_code;type:varchar(20)" json:"external_code"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	CreatedBy    *uint32        `gorm:"column:created_by" json:"created_by"`
	UpdatedBy    *uint32        `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy    *uint32        `gorm:"column:deleted_by" json:"deleted_by"`
}
