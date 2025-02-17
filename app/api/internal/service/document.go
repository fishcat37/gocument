package service

import (
	"gocument/app/api/global"
	"gocument/app/api/internal/consts"
	"gocument/app/api/internal/dao"
	"gocument/app/api/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 只需要分享码和自己的authorization和文档id的路径参数
func GetSharedDocument(c *gin.Context) {
	id := c.Param("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	var document model.Document
	var wholeDocument model.WholeDocument
	document.ID = uint(id64)
	if err := dao.FindDocumentByID(&document, &wholeDocument); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找失败" + err.Error()})
		global.Logger.Error("数据库查找失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "document": wholeDocument})
	return
}
