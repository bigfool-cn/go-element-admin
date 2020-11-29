package main

import (
  "github.com/gin-gonic/gin"
  "go-element-admin-api/configs"
  "go-element-admin-api/routers"
  "log"
)

// @title Go-Admin-Element
// @description Go-Admin-Element后端API文档
// @host 127.0.0.1:8001
// @BasePath
func main()  {
	gin.SetMode(gin.DebugMode)
	r := routers.InitRouter()
	if err := r.Run(configs.ApplicationConfig.Host + ":" + configs.ApplicationConfig.Port); err != nil {
		log.Fatal(err)
	}
}

