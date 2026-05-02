package bed

type CreateBedRequest struct {
	BedNumber    string  `json:"bed_number"    binding:"required,max=255"`
	Description  string  `json:"description"   binding:"required"`
	RoomID       *uint32 `json:"room_id"`
	BedStatusID  *uint32 `json:"bed_status_id"`
	MergedBedID  *uint32 `json:"merged_bed_id"`
	ExternalCode string  `json:"external_code" binding:"omitempty,max=20"`
}

type UpdateBedRequest struct {
	BedNumber    string  `json:"bed_number"    binding:"required,max=255"`
	Description  string  `json:"description"   binding:"required"`
	RoomID       *uint32 `json:"room_id"`
	BedStatusID  *uint32 `json:"bed_status_id"`
	MergedBedID  *uint32 `json:"merged_bed_id"`
	ExternalCode string  `json:"external_code" binding:"omitempty,max=20"`
	IsActive     *bool   `json:"is_active"`
}
