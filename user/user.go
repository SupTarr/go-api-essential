package user

import (
	"github.com/SupTarr/go-api-essential/utils"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	utils.BaseModel
	Email    string `gorm:"unique"`
	Password string
}

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type UserData struct {
	ID    uint
	Email string
}
