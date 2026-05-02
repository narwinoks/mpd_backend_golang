package ward

import (
	"backend-app/internal/modules/master/model/department"
	"time"
)

type WardResponse struct {
	ID                string                    `json:"id"`
	WardName          string                    `json:"ward_name"`
	DepartmentID      *uint32                   `json:"department_id"`
	IsExecutive       bool                      `json:"is_executive"`
	Icon              *string                   `json:"icon"`
	QueueNumberPrefix *string                   `json:"queue_number_prefix"`
	ExternalCode      string                    `json:"external_code"`
	IsActive          bool                      `json:"is_active"`
	CreatedAt         time.Time                 `json:"created_at"`
	UpdatedAt         time.Time                 `json:"updated_at"`
	Department        *DepartmentDetailResponse `json:"department,omitempty"`
}

type DepartmentDetailResponse struct {
	ID             uint32 `json:"id"`
	UUID           string `json:"uuid"`
	DepartmentName string `json:"department_name"`
}

func FromWard(u *department.Ward) *WardResponse {
	if u == nil {
		return nil
	}

	response := &WardResponse{
		ID:                u.UUID,
		WardName:          u.WardName,
		DepartmentID:      u.DepartmentID,
		IsExecutive:       u.IsExecutive,
		Icon:              u.Icon,
		QueueNumberPrefix: u.QueueNumberPrefix,
		ExternalCode:      u.ExternalCode,
		IsActive:          u.IsActive,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
	}

	if u.Department.ID != 0 {
		response.Department = &DepartmentDetailResponse{
			ID:             u.Department.ID,
			UUID:           u.Department.UUID,
			DepartmentName: u.Department.DepartmentName,
		}
	}
	return response
}

func FromWards(wards []department.Ward) []WardResponse {
	res := make([]WardResponse, 0, len(wards))

	for i := range wards {
		wardResponse := FromWard(&wards[i])
		if wardResponse != nil {
			res = append(res, *wardResponse)
		}
	}

	return res
}
