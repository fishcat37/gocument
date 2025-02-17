package middleware

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gocument/app/api/global"
	"gocument/app/api/internal/consts"
	"gocument/app/api/internal/dao"
	"gocument/app/api/internal/model"
	"net/http"
	"strconv"
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
		if err := dao.FindDocumentByID(&document, &wholeDocument); err != nil {
			c.JSON(500, gin.H{"error": "数据库查找失败" + err.Error()})
			global.Logger.Error("数据库查找失败" + err.Error())
			c.Abort()
			return
		}
		if wholeDocument.Document.Authority == 0 {
			c.Set("id", wholeDocument.Document.ID)
			c.Set("authority", 0)
		} else {
			c.Set("id", wholeDocument.Document.ID)
			c.Set("authority", 1)
		}
		c.Next()
	}
}

func GetShare() gin.HandlerFunc {
	return func(c *gin.Context) {
		var document model.Document
		id := c.Param("id")
		id64, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获得文件id出错"})
			c.Abort()
		}
		var wholeDocument model.WholeDocument
		document.ID = uint(id64)
		err = dao.FindDocumentByID(&document, &wholeDocument)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": consts.DatabaseFindFailed, "error": "查找文档失败"})
			global.Logger.Error("查找文档失败")
			c.Abort()
		}
		if wholeDocument.Document.Authority == 0 {
			AuthMiddleware()(c)
			contentID, get := c.GetQuery("contendID")
			if !get {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": consts.GetIDFailed, "error": "获取内容id错误"})
				c.Abort()
				return
			}
			contentID1, err := primitive.ObjectIDFromHex(contentID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": consts.TypeConversionFailed, "error": "类型转换错误"})
				c.Abort()
				return
			}
			if contentID1 != wholeDocument.Content.ID {
				c.JSON(http.StatusBadRequest, gin.H{"status": consts.NoAuthority, "error": "授权错误"})
				c.Abort()
				return
			}
			c.Next()
		} else {

			c.Next()
		}
	}
}
