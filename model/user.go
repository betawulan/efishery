package model

import "time"

type User struct {
	ID        int64     `json:"-"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Password  string    `json:",omitempty" swaggerignore:"true"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
}

type UserFilter struct {
	Phone string
	Name  string
}

type UserResponse struct {
	Password string `json:"password" example:"1BtL"`
}
