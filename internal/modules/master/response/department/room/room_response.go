package room

import "time"

type RoomResponse struct {
	ID                 string    `json:"id"`
	RoomName           string    `json:"room_name"`
	WardID             *uint32   `json:"ward_id"`
	ClassID            *uint32   `json:"class_id"`
	RsOnlineCode       *string   `json:"rs_online_code"`
	BedCount           int       `json:"bed_count"`
	OccupiedRoomCount  int       `json:"occupied_room_count"`
	AvailableRoomCount int       `json:"available_room_count"`
	ExternalCode       string    `json:"external_code"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
