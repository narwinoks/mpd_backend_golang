package room

type CreateRoomRequest struct {
	RoomName           string  `json:"room_name"            binding:"required,max=255"`
	WardID             *uint32 `json:"ward_id"`
	ClassID            *uint32 `json:"class_id"`
	RsOnlineCode       *string `json:"rs_online_code"       binding:"omitempty,max=50"`
	BedCount           int     `json:"bed_count"            binding:"min=0"`
	OccupiedRoomCount  int     `json:"occupied_room_count"  binding:"min=0"`
	AvailableRoomCount int     `json:"available_room_count" binding:"min=0"`
	ExternalCode       string  `json:"external_code"        binding:"omitempty,max=20"`
}

type UpdateRoomRequest struct {
	RoomName           string  `json:"room_name"            binding:"required,max=255"`
	WardID             *uint32 `json:"ward_id"`
	ClassID            *uint32 `json:"class_id"`
	RsOnlineCode       *string `json:"rs_online_code"       binding:"omitempty,max=50"`
	BedCount           int     `json:"bed_count"            binding:"min=0"`
	OccupiedRoomCount  int     `json:"occupied_room_count"  binding:"min=0"`
	AvailableRoomCount int     `json:"available_room_count" binding:"min=0"`
	ExternalCode       string  `json:"external_code"        binding:"omitempty,max=20"`
	IsActive           *bool   `json:"is_active"`
}
