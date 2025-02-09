package service

import (
	"github.com/gin-gonic/gin"
	"gocument/app/api/global"
	"gocument/app/api/internal/consts"
	"gocument/app/api/internal/dao"
	"gocument/app/api/internal/model"
	"net/http"
)

func GetMyDocument(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	document := model.Document{UserID: id.(uint)}
	if err := c.ShouldBindUri(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": "参数错误" + err.Error()})
		global.Logger.Info("用户请求错误" + err.Error())
		return
	}
	var wholeDocument model.WholeDocument
	if err := dao.FindDocument(&document, &wholeDocument); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找失败" + err.Error()})
		global.Logger.Error("数据库查找失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "document": wholeDocument})
}

func GetMyDocuments(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	document := model.Document{UserID: id.(uint)}
	var documentList []model.Document
	if err := dao.GetDocumentList(document, &documentList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找失败" + err.Error()})
		global.Logger.Error("数据库查找失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "documentList": documentList})
}

func UpdateMyDocument(c *gin.Context) {
	id, err := c.Get("ID")
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	document := model.Document{UserID: id.(uint)}
	var documentContent model.DocumentContent
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": "参数错误" + err.Error()})
		global.Logger.Info("用户请求错误" + err.Error())
		return
	}
	if err := c.ShouldBindJSON(&documentContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": "参数错误" + err.Error()})
		global.Logger.Info("用户请求错误" + err.Error())
		return
	}
	var err1 error
	err1 = dao.UpdateDocument(&document, &documentContent)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库更新失败" + err1.Error()})
		global.Logger.Error("数据库更新失败" + err1.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success})
}

func DeleteMyDocument(c *gin.Context) {
	id, err := c.Get("ID")
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	document := model.Document{UserID: id.(uint)}
	if err := c.ShouldBindUri(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": "参数错误" + err.Error()})
		global.Logger.Info("用户请求错误" + err.Error())
		return
	}
	if err := dao.DeleteDocument(&document); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库删除失败" + err.Error()})
		global.Logger.Error("数据库删除失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success})
}
