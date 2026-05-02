package department

import "backend-app/internal/base/models"

type Room struct {
	models.BaseModel
	WardID             *uint32 `gorm:"column:ward_id"                               json:"ward_id"`
	ClassID            *uint32 `gorm:"column:class_id"                              json:"class_id"`
	RsOnlineCode       *string `gorm:"column:rs_online_code;type:varchar(50)"       json:"rs_online_code"`
	RoomName           string  `gorm:"column:room_name;type:varchar(255);not null"  json:"room_name"`
	BedCount           int     `gorm:"column:bed_count;not null;default:0"          json:"bed_count"`
	OccupiedRoomCount  int     `gorm:"column:occupied_room_count;not null;default:0"  json:"occupied_room_count"`
	AvailableRoomCount int     `gorm:"column:available_room_count;not null;default:0" json:"available_room_count"`
}

func (Room) TableName() string {
	return "rooms_m"
}
