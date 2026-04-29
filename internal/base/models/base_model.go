package models

import (
	"time"

	"github.com/google/uuid"
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

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if m.UUID == "" {
		m.UUID = uuid.New().String()
	}

	ctx := tx.Statement.Context
	if ctx != nil {
		var executorID uint32
		if employeeID, ok := ctx.Value("employee_id").(uint32); ok {
			executorID = employeeID
		} else if userID, ok := ctx.Value("user_id").(uint32); ok {
			executorID = userID
		}

		if executorID != 0 {
			m.CreatedBy = new(uint32)
			*m.CreatedBy = executorID
			m.UpdatedBy = new(uint32)
			*m.UpdatedBy = executorID
		}

		if profileID, ok := ctx.Value("profile_id").(uint32); ok {
			m.ProfileID = new(uint32)
			*m.ProfileID = profileID
		}

		if externalCode, ok := ctx.Value("external_code").(string); ok {
			m.ExternalCode = externalCode
		}
	}
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context
	if ctx != nil {
		var executorID uint32
		if employeeID, ok := ctx.Value("employee_id").(uint32); ok {
			executorID = employeeID
		} else if userID, ok := ctx.Value("user_id").(uint32); ok {
			executorID = userID
		}

		if executorID != 0 {
			m.UpdatedBy = new(uint32)
			*m.UpdatedBy = executorID
		}
	}
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context
	if ctx != nil {
		var executorID uint32
		if employeeID, ok := ctx.Value("employee_id").(uint32); ok {
			executorID = employeeID
		} else if userID, ok := ctx.Value("user_id").(uint32); ok {
			executorID = userID
		}

		if executorID != 0 {
			m.DeletedBy = new(uint32)
			*m.DeletedBy = executorID
		}
	}
	return
}

func (m *BaseModel) SetNonActive(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	if ctx != nil {
		var executorID uint32
		if employeeID, ok := ctx.Value("employee_id").(uint32); ok {
			executorID = employeeID
		} else if userID, ok := ctx.Value("user_id").(uint32); ok {
			executorID = userID
		}

		if executorID != 0 {
			m.DeletedBy = new(uint32)
			*m.DeletedBy = executorID
		}
	}
	m.IsActive = false
	m.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if tx.Statement.Model != nil {
		return tx.Save(tx.Statement.Model).Error
	}

	return tx.Save(m).Error
}
