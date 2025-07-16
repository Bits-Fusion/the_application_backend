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
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
