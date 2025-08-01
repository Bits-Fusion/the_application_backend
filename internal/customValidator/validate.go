package customvalidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var permissionFormatRegex = regexp.MustCompile(`^[a-zA-Z]+_[a-zA-Z]+$`)

func PermissionFormatValidator(fl validator.FieldLevel) bool {
	tag := fl.Field().String()
	return permissionFormatRegex.MatchString(tag)
}
