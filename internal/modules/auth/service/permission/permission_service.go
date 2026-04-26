package permission

type PermissionService interface {
	GetUserPermissions(userID uint32) ([]string, error)
}
