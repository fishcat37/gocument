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

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": "参数错误" + err.Error()})
		global.Logger.Info("用户请求错误" + err.Error())
		return
	}
	isFind, err := dao.FindUser(&user)
	if isFind {
		c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.UserExist, "mess": "用户已存在"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找出错" + err.Error()})
		global.Logger.Error("MySQL查找用户出错" + err.Error())
		return
	}
	err = dao.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户插入数据库失败" + err.Error()})
		global.Logger.Error("MySQL插入用户出错" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "ID": user.ID})
}

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": err.Error()})
		global.Logger.Info("用户请求错误" + err.Error())
		return
	}
	isFind, err := dao.FindUser(&user)
	if !isFind {
		c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.UserNotExist, "mess": "用户不存在或密码错误"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找出错" + err.Error()})
		global.Logger.Error("MySQL查找用户出错" + err.Error())
		return
	}
	token, err := utils.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token创建失败" + err.Error()})
		global.Logger.Error("token创建失败" + err.Error())
		return
	}
	refreshToken, err := utils.CreateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "refreshToken创建失败" + err.Error()})
		global.Logger.Error("refreshToken创建失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "token": token, "refreshToken": refreshToken})
}

func Change(c *gin.Context) {
	var user model.User
	//id,_:=c.Get("ID")
	//var ok bool;
	//if user.ID,ok=id.(int);!ok{
	//
	//}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": err.Error()})
		global.Logger.Info("用户修改信息的请求错误" + err.Error())
		return
	}
	isFind, err := dao.FindUser(&user)
	if !isFind {
		c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.UserNotExist, "mess": "用户不存在或密码错误"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找出错" + err.Error()})
		global.Logger.Error("MySQL查找用户出错" + err.Error())
		return
	}
	err = dao.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库更新出错" + err.Error()})
		global.Logger.Error("MySQL更新用户数据失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success})
}
