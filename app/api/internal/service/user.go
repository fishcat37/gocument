package service

import (
	"gocument/app/api/global"
	"gocument/app/api/internal/consts"
	"gocument/app/api/internal/dao"
	"gocument/app/api/internal/model"
	"gocument/app/api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

// 需要JSON格式的用户名与密码
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": consts.DatabaseFindFailed,
			"error":  "数据库查找出错" + err.Error()})
		global.Logger.Error("MySQL查找用户出错" + err.Error())
		return
	}
	err = dao.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": consts.DatabaseAddFailed,
			"error":  "用户插入数据库失败" + err.Error()})
		global.Logger.Error("MySQL插入用户出错" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "ID": user.ID})
}

// 需要JSON格式的用户名与密码
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": consts.DatabaseFindFailed,
			"error":  "数据库查找出错" + err.Error()})
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

// 需要Authorization
func RefreshToken(c *gin.Context) {
	user := model.User{}
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetIDFailed, "error": "获取ID失败"})
		global.Logger.Info("获取ID失败")
		return
	}
	username, get := c.Get("Username")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.GetNameFailed, "error": "获取用户姓名失败"})
		global.Logger.Info("获取用户姓名失败")
		return
	}
	user.ID = id.(uint)
	user.Username = username.(string)
	Token, err := utils.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token创建失败" + err.Error()})
		global.Logger.Error("token创建失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success, "token": Token})
}

// 需要JSON格式的用户信息
// Email    string `json:"email"`
// Gender   string `json:"gender"`
// Avatar   string `json:"avatar"`
// Age      int    `json:"age"·
func Change(c *gin.Context) {
	var user model.User
	username, get := c.Get("Username")
	if !get {
		c.JSON(http.StatusInternalServerError, gin.H{"status": consts.GetNameFailed, "error": "获取用户名失败"})
		global.Logger.Error("获取用户名失败")
		return
	}
	user.Username = username.(string)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": consts.ShouldBindFailed, "error": err.Error()})
		global.Logger.Info("用户修改信息的请求错误" + err.Error())
		return
	}
	//isFind, err := dao.FindUser(&user)
	//if !isFind {
	//	c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.UserNotExist, "mess": "用户不存在或密码错误"})
	//	return
	//} else if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查找出错" + err.Error()})
	//	global.Logger.Error("MySQL查找用户出错" + err.Error())
	//	return
	//}
	err := dao.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "数据库更新出错" + err.Error()})
		global.Logger.Error("MySQL更新用户数据失败" + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success})
}

// 需要authorization
func GetUserInfo(c *gin.Context) {
	var user model.User
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusInternalServerError, gin.H{"status": consts.GetIDFailed, "error": "获取id失败"})
		global.Logger.Error("获取id失败")
		return
	}
	username, get := c.Get("Username")
	if !get {
		c.JSON(http.StatusInternalServerError, gin.H{"status": consts.GetIDFailed, "error": "获取用户名失败"})
		global.Logger.Error("获取用户名失败")
		return
	}
	user.ID = id.(uint)
	user.Username = username.(string)
	isFind, err := dao.FindUserByName(&user)
	if !isFind {
		c.JSON(http.StatusOK, gin.H{"info": "error", "status": consts.UserNotExist, "error": "用户不存在"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": consts.DatabaseFindFailed, "error": "mysql查找用户出错"})
		global.Logger.Error("mysql查找用户失败")
		return
	}
	c.Render(http.StatusOK, render.JSON{Data: gin.H{"info": "success", "status": consts.Success, "data": user}})
}

// query
// oldPassword和newPassword
func ChangePassword(c *gin.Context) {
	username, get := c.Get("Username")
	if !get {
		c.JSON(http.StatusInternalServerError, gin.H{"status": consts.GetIDFailed, "error": "获取用户名失败"})
		global.Logger.Error("获取用户名失败")
		return
	}
	oldPassword := c.Query("oldPassword")
	newPassword := c.Query("newPassword")
	user := model.User{
		Username: username.(string),
		Password: oldPassword,
	}
	err := dao.UpdatePassword(&user, newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": consts.DatabaseUpdateFailed, "error": "更改密码失败"})
		global.Logger.Error("更改密码失败 " + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": consts.Success})
}
