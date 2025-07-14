package entities

import "time"

type Role string

type FilterField string

const (
	FilterByAll         FilterField = "all"
	FilterByID          FilterField = "id"
	FilterByUsername    FilterField = "username"
	FilterByEmail       FilterField = "email"
	FilterByPhoneNumber FilterField = "phone_number"
)

const (
	AdminRole  Role = "admin"
	UserRole   Role = "user"
	EditorRole Role = "editor"
)

type User struct {
	Id           uint32    `json:"id"`
	Username     string    `json:"username"`
	Role         Role      `json:"role" gorm:"type:role_enum"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"-"`
	ProfileImage string    `json:"profile_image"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsActive     bool      `json:"is_active" gorm:"default:0"`
}

type InsertUserData struct {
	Username    string
	Role        Role
	Email       string
	PhoneNumber string
	Password    string
	FirstName   string
	LastName    string
}

type InsertUserDTO struct {
	Username    string
	Role        string
	Email       string
	PhoneNumber string
	Password    string
	FirstName   string
	LastName    string
}

type FilterParams struct {
	Page    int32
	Limit   int32
	OrderBy string
}
