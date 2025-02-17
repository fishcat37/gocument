package router

import (
	"github.com/gin-gonic/gin"
	"gocument/app/api/global"
	"gocument/app/api/internal/middleware"
	"gocument/app/api/internal/service"
)

func InitRouter() {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", service.Register)
		user.GET("/login", service.Login)
		user.GET("/refresh", middleware.AuthMiddleware(), service.RefreshToken)
		user.PUT("/change", middleware.AuthMiddleware(), service.Change)
		user.GET("/info", middleware.AuthMiddleware(), service.GetUserInfo)
		user.PUT("/password", middleware.AuthMiddleware(), service.ChangePassword)
	}

	r.GET("/document/share/:id", middleware.GetShare(), service.GetSharedDocument)
	myDocuments := r.Group("/myDocuments", middleware.AuthMiddleware())
	{
		myDocuments.POST("/create", service.Create)
		myDocuments.GET("/get/:id", service.GetMyDocument)
		myDocuments.GET("", service.GetMyDocuments)
		myDocuments.GET("/share", service.ShareMyDocument)
		myDocuments.PUT("/update/:id", service.UpdateMyDocument)
		myDocuments.DELETE("/delete/:id", service.DeleteMyDocument)
	}
	err := r.Run()
	if err != nil {
		global.Logger.Fatal("init router failed" + err.Error())
	}
	global.Logger.Info("init router success")
}
