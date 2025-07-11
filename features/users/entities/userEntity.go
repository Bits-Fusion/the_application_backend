package entities

import "time"

type User struct {
	Id           uint32    `json:"id"`
	Username     string    `json:"username"`
	Role         string    `json:"role"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"-"`
	ProfileImage string    `json:"profile_image"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsActive     bool      `json:"is_active"`
}
