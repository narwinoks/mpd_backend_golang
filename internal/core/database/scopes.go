package database

import (
	"gorm.io/gorm"
)

func FilterProfile(db *gorm.DB) *gorm.DB {
	if db.Statement.Schema != nil {
		if _, ok := db.Statement.Schema.FieldsByDBName["profile_id"]; ok {
			ctx := db.Statement.Context
			if ctx != nil {
				if profileID, ok := ctx.Value("profile_id").(uint32); ok {
					return db.Where("profile_id = ?", profileID)
				}
			}
		}
	}
	return db
}

func DefaultOrder(db *gorm.DB) *gorm.DB {
	if _, ok := db.Statement.Clauses["ORDER BY"]; !ok {
		return db.Order("id DESC")
	}
	return db
}

// WithoutDefaultOrder is a marker scope (optional usage depending on how global callbacks are implemented)
func WithoutDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Set("gorm:without_default_order", true)
}
