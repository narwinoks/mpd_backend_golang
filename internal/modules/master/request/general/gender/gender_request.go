package gender

type CreateGenderRequest struct {
	Code   string `json:"code" binding:"required,max=4"`
	Gender string `json:"gender" binding:"required,max=20"`
}

type UpdateGenderRequest struct {
	Code   string `json:"code" binding:"required,max=4"`
	Gender string `json:"gender" binding:"required,max=20"`
}
