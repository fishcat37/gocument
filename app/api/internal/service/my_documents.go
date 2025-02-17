package service

import (
	"github.com/gin-gonic/gin"
	"gocument/app/api/global"
	"gocument/app/api/internal/consts"
	"gocument/app/api/internal/dao"
	"gocument/app/api/internal/model"
	"net/http"
	"strconv"
)

//全部需要authorization，都会得到用户id

// 需要form-data格式的文档名，正文（html），是否公开
func Create(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	document := model.Document{UserID: id.(uint)}
	var documentContent model.DocumentContent
	var err error
	document.Authority, err = strconv.Atoi(c.PostForm("authority"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetAuthorityFailed, "error": err})
		global.Logger.Info("获取文档权限失败")
		return
	}
	document.Title = c.PostForm("title")
	documentContent.Content = c.PostForm("content")
	if err := dao.InsertDocument(&document, &documentContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库插入失败" + err.Error()})
		global.Logger.Error("数据库插入失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "ID": document.ID})
}

// 需要路径参数文档id
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

// 不需要参数
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

// 需要query格式的文档id
func ShareMyDocument(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	document := model.Document{UserID: id.(uint)}
	if err := c.ShouldBindQuery(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": "参数错误" + err.Error()})
		global.Logger.Info("用户请求错误" + err.Error())
		return
	}
	var wholeDocument model.WholeDocument
	err := dao.FindDocumentByID(&document, &wholeDocument)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": consts.DatabaseFindFailed, "error": "查找文档失败" + err.Error()})
		global.Logger.Error("查找文档失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": consts.Success, "info": "success", "share_id": wholeDocument.Content.ID})
}

// 需要form-data格式的文档id与文档内容，可选参数为title和权限
func UpdateMyDocument(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	document := model.Document{UserID: id.(uint)}
	var documentContent model.DocumentContent
	var err error
	var authority string
	documentid := c.Param("id")
	documentid64, err := strconv.ParseUint(documentid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.TypeConversionFailed, "error": err})
		global.Logger.Info("获取文档ID失败")
		return
	}
	document.ID = uint(documentid64)
	authority = c.PostForm("authority")
	if authority != "" {
		document.Authority, err = strconv.Atoi(authority)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetAuthorityFailed, "error": err})
			global.Logger.Info("获取文档权限失败")
			return
		}
	}
	document.Title = c.PostForm("title")
	documentContent.Content = c.PostForm("content")
	var err1 error
	err1 = dao.UpdateDocument(&document, &documentContent)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库更新失败" + err1.Error()})
		global.Logger.Error("数据库更新失败" + err1.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success})
}

// 路径参数文档id
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
