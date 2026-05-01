package religion

type CreateReligionRequest struct {
	Religion string `json:"religion" binding:"required,max=100"`
}

type UpdateReligionRequest struct {
	Religion string `json:"religion" binding:"required,max=100"`
}
