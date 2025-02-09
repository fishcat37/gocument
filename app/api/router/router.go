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
		user.PUT("/change", middleware.AuthMiddleware(), service.Change)
	}
	document := r.Group("/document", middleware.AuthMiddleware())
	{
		document.POST("/create", service.Create)
		//document.GET("/get", service.GetDocument)
		//document.PUT("/update", service.UpdateDocument)
		document.GET("/share", service.ShareDocument)
	}
	r.GET("/document/share/:id", middleware.FindAuthorityMiddleware(), service.GetSharedDocument)
	myDocuments := r.Group("/myDocuments", middleware.AuthMiddleware())
	{
		myDocuments.GET("/get/id", service.GetMyDocument)
		myDocuments.GET("", service.GetMyDocuments)
		myDocuments.PUT("/update", service.UpdateMyDocument)
		myDocuments.DELETE("/delete", service.DeleteMyDocument)
	}
	err := r.Run()
	if err != nil {
		global.Logger.Fatal("init router failed" + err.Error())
	}
	global.Logger.Info("init router success")
}
