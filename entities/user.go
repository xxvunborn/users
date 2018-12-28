package entities

import "time"

// User model(entities)
type User struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// UserQuery ...
type UserQuery struct {
	ID       string
	Email    string
	FullName string
}
