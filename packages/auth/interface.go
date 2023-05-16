package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claim struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	jwt.StandardClaims
}