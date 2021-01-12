package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"go-element-admin/configs"
	"go-element-admin/routers"
)

// @title Go-Admin-Element
// @description Go-Admin-Element后端API文档
// @host 127.0.0.1:8001
// @BasePath
func main()  {
  if configs.ApplicationConfig.Debug {
    gin.SetMode(gin.DebugMode)
  } else {
    gin.SetMode(gin.ReleaseMode)
  }
	r := routers.InitRouter()
	if err := r.Run(configs.ApplicationConfig.Host + ":" + configs.ApplicationConfig.Port); err != nil {
		log.Fatal(err)
	}
}

