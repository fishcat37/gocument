package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"gocument/app/api/global"
	"gocument/app/api/internal/model"
	"time"
)

func CreateToken(user model.User) (string, error) {
	claims := model.CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), //1小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    global.Config.JwtConfig.Issuer, //签发者
			Subject:   user.Username,                  //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.JwtConfig.JwtSecretKey))
}
