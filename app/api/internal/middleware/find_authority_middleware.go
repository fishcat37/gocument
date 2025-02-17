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
