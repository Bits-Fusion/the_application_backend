package repositories

import (
	"errors"

	"github.com/Bits-Fusion/the_application_backend/database"
	"github.com/Bits-Fusion/the_application_backend/features/permissions/entities"

	"gorm.io/gorm"
)

type permissionRepository struct {
	db database.Database
}

func NewPermissionRepository(db database.Database) *permissionRepository {
	return &permissionRepository{
		db: db,
	}
}

func (r *permissionRepository) CreatePermission(action, resource string) (*entities.Permission, error) {
	var perm entities.Permission

	result := r.db.GetDb().Where("action = ? AND resource = ?", action, resource).First(&perm)

	if result.Error == nil {
		return &perm, nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	perm = entities.Permission{
		Action:   entities.Action(action),
		Resource: resource,
	}
	if err := r.db.GetDb().Create(&perm).Error; err != nil {
		return nil, err
	}

	return &perm, nil
}

func (r *permissionRepository) UpdatePermission(id int32, action, resource *string) (entities.Permission, error) {
	var permission entities.Permission

	err := r.db.GetDb().First(&permission, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.Permission{}, err
	}
	if err != nil {
		return entities.Permission{}, errors.Join(errors.New("internal server error"), err)
	}

	if action != nil {
		permission.Action = entities.Action(*action)
	}

	if resource != nil {
		permission.Resource = *resource
	}

	if err := r.db.GetDb().Save(&permission).Error; err != nil {
		return entities.Permission{}, err
	}

	return permission, nil
}

func (r *permissionRepository) GetPermission(id int32) (entities.Permission, error) {
	var permission entities.Permission

	err := r.db.GetDb().First(&permission, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.Permission{}, err
	}

	if err != nil {
		return entities.Permission{}, errors.Join(errors.New("internal server error"), err)
	}

	return permission, nil
}

func (r *permissionRepository) DeletePermission(id int32) error {
	ctx := r.db.GetDb().Delete(&entities.Permission{}, id)

	if ctx.RowsAffected == 0 {
		return errors.New("no recored found with this Id")
	}

	return nil
}
