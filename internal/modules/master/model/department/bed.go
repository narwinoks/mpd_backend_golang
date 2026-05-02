package department

import "backend-app/internal/base/models"

type Bed struct {
	models.BaseModel
	RoomID      *uint32 `gorm:"column:room_id"                               json:"room_id"`
	BedStatusID *uint32 `gorm:"column:bed_status_id"                         json:"bed_status_id"`
	BedNumber   string  `gorm:"column:bed_number;type:varchar(255);not null" json:"bed_number"`
	Description string  `gorm:"column:description;type:text;not null"        json:"description"`
	MergedBedID *uint32 `gorm:"column:merged_bed_id"                         json:"merged_bed_id"`
}

func (Bed) TableName() string {
	return "beds_m"
}
