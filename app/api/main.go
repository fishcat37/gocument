package main

import (
	"gocument/app/api/internal/initialize"
	"gocument/app/api/router"
)

func main() {
	//r := gin.Default()
	//router.Router(r)
	//if err:=r.Run(":8080"); err != nil {
	//	panic(err)
	//}
	initialize.SetupViper()
	initialize.SetupLogger()
	initialize.SetupDatabase()
	router.InitRouter()
}
