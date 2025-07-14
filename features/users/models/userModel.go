package models

type UserModel struct {
	Username    string `validate:"required" json:"username"`
	PhoneNumber string `validate:"required" json:"phoneNumber"`
	FirstName   string `validate:"required" json:"firstName"`
	LastName    string `validate:"required" json:"lastName"`
	Email       string `validate:"required,email" json:"email"`
	Password    string `validate:"required" json:"password"`
}
