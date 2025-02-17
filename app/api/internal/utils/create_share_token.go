package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"gocument/app/api/global"
	"gocument/app/api/internal/model"
	"time"
)

func CreateShareToken(document model.Document) (string, error) {
	//TODO
	claims := model.ShareClaims{
		ID:     document.ID,
		UserID: document.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    global.Config.JwtConfig.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.JwtConfig.JwtSecretKey))
}
