package model

import "time"

type Register struct {
	ID        int64     `json:"-"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Password  string    `json:",omitempty"`
	CreatedAt time.Time `json:"-"`
}

type RegisterFilter struct {
	Phone string
	Name  string
}

type RegisterResponse struct {
	Password string `json:"password"`
}
