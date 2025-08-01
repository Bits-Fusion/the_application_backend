package repositories

import "github.com/Bits-Fusion/the_application_backend/features/permissions/entities"

type PermissionRepository interface {
	CreatePermission(action, resource string) (*entities.Permission, error)
	UpdatePermission(id int32, action, resource *string) (entities.Permission, error)
	GetPermission(id int32) (entities.Permission, error)
	DeletePermission(id int32) error
}
