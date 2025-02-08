package middleware

import (
	"github.com/gin-gonic/gin"
	"gocument/app/api/global"
	"gocument/app/api/internal/dao"
	"gocument/app/api/internal/model"
)

func FindAuthorityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var document model.Document
		if err := c.ShouldBindUri(&document); err != nil {
			c.JSON(400, gin.H{"error": "参数错误" + err.Error()})
			c.Abort()
			return
		}
		var wholeDocument model.WholeDocument
		if err := dao.FindDocument(&document, &wholeDocument); err != nil {
			c.JSON(500, gin.H{"error": "数据库查找失败" + err.Error()})
			global.Logger.Error("数据库查找失败" + err.Error())
			c.Abort()
			return
		}
		if wholeDocument.Authority == 0 {
			c.Set("ID", wholeDocument.ID)
			c.Set("authority", 0)
		} else {
			c.Set("ID", wholeDocument.ID)
			c.Set("authority", 1)
		}
		c.Next()
	}
}
