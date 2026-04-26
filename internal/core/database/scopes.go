package database

import (
	"fmt"

	"gorm.io/gorm"
)

func FilterProfile(db *gorm.DB) *gorm.DB {
	if db.Statement.Schema != nil {
		if _, ok := db.Statement.Schema.FieldsByDBName["profile_id"]; ok {
			ctx := db.Statement.Context
			if ctx != nil {
				tableName := db.Statement.Schema.Table
				if profileID, ok := ctx.Value("profile_id").(uint32); ok {
					queryCondition := fmt.Sprintf("%s.profile_id = ?", tableName)
					return db.Statement.Where(queryCondition, profileID)
				}
			}
		}
	}
	return db
}

func DefaultOrder(db *gorm.DB) *gorm.DB {
	if _, ok := db.Statement.Clauses["ORDER BY"]; !ok {
		if db.Statement.Schema != nil {
			tableName := db.Statement.Schema.Table
			orderQuery := fmt.Sprintf("%s.id DESC", tableName)
			return db.Order(orderQuery)
		}

	}

	return db
}
func WithoutDefaultOrder(db *gorm.DB) *gorm.DB {
	return db.Set("gorm:without_default_order", true)
}
