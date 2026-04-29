package permission

type CreatePermissionRequest struct {
	Permission string `json:"permission" binding:"required"`
}

type UpdatePermissionRequest struct {
	Permission string `json:"permission" binding:"required"`
}
