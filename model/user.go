package model

import "time"

type User struct {
	ID        int64     `json:"-"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Password  string    `json:",omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type UserFilter struct {
	Phone string
	Name  string
}

type UserResponse struct {
	Password string `json:"password"`
}
