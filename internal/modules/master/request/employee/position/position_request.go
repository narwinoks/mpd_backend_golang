package position

type CreatePositionRequest struct {
	Position string `json:"position" binding:"required,max=100"`
}

type UpdatePositionRequest struct {
	Position string `json:"position" binding:"required,max=100"`
}
