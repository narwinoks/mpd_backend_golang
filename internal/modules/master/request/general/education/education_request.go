package education

type CreateEducationRequest struct {
	EducationType string `json:"education_type" binding:"required,oneof=FORMAL INFORMAL OTHER"`
	Code          string `json:"code" binding:"required,max=50"`
	Name          string `json:"name" binding:"required,max=100"`
	SortOrder     int    `json:"sort_order"`
}

type UpdateEducationRequest struct {
	EducationType string `json:"education_type" binding:"required,oneof=FORMAL INFORMAL OTHER"`
	Code          string `json:"code" binding:"required,max=50"`
	Name          string `json:"name" binding:"required,max=100"`
	SortOrder     int    `json:"sort_order"`
}
