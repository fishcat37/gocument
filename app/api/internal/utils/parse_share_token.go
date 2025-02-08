package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gocument/app/api/global"
	"gocument/app/api/internal/model"
	"time"
)

func ParseShareToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &model.ShareClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JwtConfig.JwtSecretKey), nil
	})
	if err != nil {

		return err
	}
	claims, ok := token.Claims.(*model.ShareClaims)
	if !ok || !token.Valid {

		return fmt.Errorf("invalid token")
	}
	if claims.ExpiresAt.Unix() < time.Now().Unix() {

		return fmt.Errorf("token expired")
	} else if claims.Issuer != global.Config.JwtConfig.Issuer {

		return fmt.Errorf("invalid token")
	}
	return nil
}
