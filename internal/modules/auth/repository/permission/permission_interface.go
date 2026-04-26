package permission

type PermissionRepository interface {
	GetUserPermissions(userID uint32, roleID uint32) ([]string, error)
}
