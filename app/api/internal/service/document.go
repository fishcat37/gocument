package service

import (
	"github.com/gin-gonic/gin"
	"gocument/app/api/global"
	"gocument/app/api/internal/consts"
	"gocument/app/api/internal/dao"
	"gocument/app/api/internal/model"
	"gocument/app/api/internal/utils"
	"net/http"
)

func Create(c *gin.Context) {
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
	if err := dao.InsertDocument(&document, &documentContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库插入失败" + err.Error()})
		global.Logger.Error("数据库插入失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "ID": document.ID})
}

func GetDocument(c *gin.Context) {
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
	//if err := c.ShouldBindJSON(&document); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": "参数错误" + err.Error()})
	//	global.Logger.Info("用户请求错误" + err.Error())
	//	return
	//}
	var wholeDocument model.WholeDocument

	err1 := dao.FindDocument(&document, &wholeDocument)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找失败" + err1.Error()})
		global.Logger.Error("数据库查找失败" + err1.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "document": wholeDocument})
}

func UpdateDocument(c *gin.Context) {
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

func ShareDocument(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	document := model.Document{UserID: id.(uint)}
	if err := c.ShouldBindJSON(&document); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": "参数错误" + err.Error()})
		global.Logger.Info("用户请求错误" + err.Error())
		return
	}
	shareToken, err := utils.CreateShareToken(document)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "分享失败" + err.Error()})
		global.Logger.Error("分享token生成失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "shareToken": shareToken})
}

func GetSharedDocument(c *gin.Context) {
	authority, get := c.Get("authority")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetAuthorityFailed, "error": "获取权限失败"})
		global.Logger.Info("获取权限失败")
		return
	}
	if authority.(int) == 0 {
		token := c.Query("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.NoAuthority, "mess": "未登录"})
			return
		}
		err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.NoAuthority, "mess": "token出错"})
			return
		}
		shareToken := c.Query("shareToken")
		if shareToken == "" {
			c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.NoAuthority, "mess": "未授权"})
			return
		}
		err = utils.ParseShareToken(shareToken)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.NoAuthority, "mess": "sharetoken有错"})
			return
		}
		id, get := c.Get("ID")
		if !get {
			c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
			global.Logger.Info("获取ID失败")
			return
		}
		document := model.Document{UserID: id.(uint)}
		var wholeDocument model.WholeDocument
		if err := dao.FindDocument(&document, &wholeDocument); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找失败" + err.Error()})
			global.Logger.Error("数据库查找失败" + err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "document": wholeDocument})
		return
	} else {
		id, get := c.Get("ID")
		if !get {
			c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
			global.Logger.Info("获取ID失败")
			return
		}
		document := model.Document{UserID: id.(uint)}
		var wholeDocument model.WholeDocument
		if err := dao.FindDocument(&document, &wholeDocument); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找失败" + err.Error()})
			global.Logger.Error("数据库查找失败" + err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "document": wholeDocument})
		return
	}
}
