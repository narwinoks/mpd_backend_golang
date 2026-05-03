package database

import (
	"context"

	"gorm.io/gorm"
)

func ResolveUUID(ctx context.Context, db *gorm.DB, tableName string, uuid string) (uint32, error) {
	if uuid == "" {
		return 0, nil
	}
	var id uint32
	query := db.WithContext(ctx).Table(tableName).Select("id").Where("uuid = ?", uuid)

	if profileID, ok := ctx.Value("profile_id").(uint32); ok {
		query = query.Where("profile_id = ?", profileID)
	}
	err := query.Order("id DESC").Limit(1).Scan(&id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}
