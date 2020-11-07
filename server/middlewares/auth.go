package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
  "strings"
  "time"
	"element-admin-api/apis"
	"element-admin-api/utils"
)

// 验证jwt令牌
func JwtMiddleWare() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		log.Println("token",token)
		if token == "" {
			c.JSON(http.StatusUnauthorized,apis.Res{Code:401,Message:"token不可用"})
			c.Abort()
			return
		} else {
			claims, err := utils.Jwt.ParseToken(token)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusUnauthorized,apis.Res{Code:401,Message:"token不可用"})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusUnauthorized,apis.Res{Code:401,Message:"token已过期"})
				c.Abort()
				return
			}
			c.Set("_user_id",claims.UserId)
      c.Set("_user_name",claims.UserName)
      c.Set("_user_roles",strings.Join(utils.SliceInt64ToString(claims.Roles),","))
		}
		c.Next()
	}
}
