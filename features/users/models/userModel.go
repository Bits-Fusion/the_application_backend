package models

type UserModel struct {
	Username    string `validate:"required" json:"username"`
	PhoneNumber string `validate:"required" json:"phoneNumber"`
	FirstName   string `validate:"required" json:"firstName"`
	LastName    string `validate:"required" json:"lastName"`
	Email       string `validate:"required,email" json:"email"`
	Password    string `validate:"required" json:"password"`
}

type UserUpdateModel struct {
	Username       string   `json:"username"`
	PhoneNumber    string   `json:"phoneNumber"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Email          string   `json:"email"`
	Permission     []string `validate:"omitempty,dive,permissionformat" json:"permission"`
	Password       string   `json:"password"`
	ProfilePicture string   `json:"profilePicture"`
}

type PasswordReset struct {
	Email       string `json:"email" validate:"required,email"`
	NewPassword string `json:"newPassword" validate:"required,min=7"`
	Otp         string `json:"otp" validate:"required,min=6"`
}

type CreateUserTokenPayload struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=7,max=72"`
}

type ChunkDeletePayload struct {
	UserIds []string `validate:"required" json:"userIds"`
}

type DeleteMode string

const (
	All    DeleteMode = "all"
	Single DeleteMode = "single"
)
