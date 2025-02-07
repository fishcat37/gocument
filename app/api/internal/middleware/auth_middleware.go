package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gocument/app/api/global"
	"gocument/app/api/internal/model"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.Config.JwtConfig.JwtSecretKey), nil
		})
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*model.CustomClaims)
		if !ok || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		} else if claims.Issuer != global.Config.JwtConfig.Issuer || claims.Subject != claims.Username {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("ID", claims.ID)
		c.Next()
		return
	}
}
