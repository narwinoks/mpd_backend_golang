package bed

import "time"

type BedResponse struct {
	ID           string    `json:"id"`
	BedNumber    string    `json:"bed_number"`
	Description  string    `json:"description"`
	RoomID       *uint32   `json:"room_id"`
	BedStatusID  *uint32   `json:"bed_status_id"`
	MergedBedID  *uint32   `json:"merged_bed_id"`
	ExternalCode string    `json:"external_code"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
