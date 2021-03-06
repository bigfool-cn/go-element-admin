package routers

import (
	_ "go-element-admin/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go-element-admin/apis"
	"go-element-admin/middlewares"
)

func InitRouter() *gin.Engine  {
	r := gin.New()

	r.Use(Cors())

	r.GET("/chat", apis.Chat)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/user/login", apis.UserLogin)

	auth := r.Group("/")
	auth.Use(middlewares.JwtMiddleWare())
	{
    auth.GET("/user/info", apis.UserInfo)
	  auth.Use(middlewares.RoleCasbinMiddleWare())
	  {
      // 用户管理
      auth.GET("/users", apis.UserList)
      auth.DELETE("/users", apis.DeleteUser)
      user := auth.Group("user")
      {
        user.POST("", apis.CreateUser)
        user.PUT("/:user_id", apis.UpdateUser)
        user.POST("/pwd/:user_id", apis.UpdatePwdUser)
        user.GET("/logs", apis.UserLogList)
        user.DELETE("/logs", apis.DeleteUserLog)
      }

      // 角色管理
      auth.GET("/roles", apis.RoleList)
      auth.DELETE("/roles", apis.DeleteRole)
      role := auth.Group("role")
      {
        role.POST("", apis.CreateRole)
        role.PUT("/:role_id", apis.UpdateRole)
      }

      // 接口管理
      auth.GET("/paths", apis.PathList)
      auth.DELETE("/paths", apis.DeletePath)
      path := auth.Group("path")
      {
        path.POST("", apis.CreatePath)
        path.GET("/:path_id", apis.GetPath)
        path.PUT("/:path_id", apis.UpdatePath)
      }

      // 角色管理
      auth.GET("/menus", apis.MenuList)
      auth.DELETE("/menus", apis.DeleteMenu)
      menu := auth.Group("menu")
      {
        menu.POST("", apis.CreateMenu)
        menu.GET("/:menu_id", apis.GetMenu)
        menu.PUT("/:menu_id", apis.UpdateMenu)
      }
    }

	}
	return r
}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
