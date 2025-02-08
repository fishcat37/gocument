package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	ID       uint
	Username string
	jwt.RegisteredClaims
}

type ShareClaims struct {
	ID     uint
	UserID uint
	jwt.RegisteredClaims
}
