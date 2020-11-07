package middlewares

import (
  "element-admin-api/utils/gcasbin"
  "github.com/gin-gonic/gin"
  "strings"
)

// role casbin middleware 权限认证中间件
func RoleCasbinMiddleWare() gin.HandlerFunc {
  return func(c *gin.Context) {
    // 请求的path
    p := c.Request.URL.Path
    // 请求的方法
    m := c.Request.Method
    rolesStr, bol := c.Get("_user_roles")
    if !bol {
     c.JSON(403, gin.H{
       "code":    403,
       "message": "没有操作权限",
       "data":    "",
     })
     c.Abort()
     return
    }

    roles := strings.Split(rolesStr.(string),",")

    var abort = true
    for _,role := range roles  {
     res, _ := gcasbin.Enforcer.Enforce(role,p,m)
     if res == true {
       abort = false
       break
     }
    }
    if abort == true {
     c.JSON(403, gin.H{
       "code":    403,
       "message": "没有操作权限",
       "data":    "",
     })
     c.Abort()
     return
    }
    c.Next()
  }
}
