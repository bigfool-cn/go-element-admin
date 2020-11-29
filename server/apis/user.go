package apis

import (
  "encoding/json"
  "github.com/gin-gonic/gin"
  "go-element-admin-api/models"
  "go-element-admin-api/utils"
  "log"
  "strconv"
)

// @Summary 用户登录
// @Tags 用户管理
// @accept json
// @Produce json
// @Param user body loginForm true "账号密码"
// @Success 200 {object} Res {"code":0,"data":{token:"token"},"msg":"登录成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var (
    userModel models.User
    loginForm loginForm
    result struct {
      Token string `json:"token"`
    }
  )
	err := c.BindJSON(&loginForm)
  if err != nil {
    c.JSON(400, Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }
  userModel.UserName = loginForm.UserName
	var password = utils.MD5(loginForm.Password)
	user, err := userModel.GetUser()
	if err != nil {
		c.JSON(400, Res{Code:400,Message:"用户名错误"})
    return
	} else if password != user.Password {
    c.JSON(400, Res{Code:400,Message:"密码错误"})
    return
  } else if user.Status != 1 {
    c.JSON(400, Res{Code:400,Message:"账号已被停用"})
    return
  } else {
    var roles []int64
    for _, role := range user.Roles  {
      roles = append(roles,role.RoleID)
    }
		token,err := utils.Jwt.GenerateToken(user.UserId,user.UserName,roles)
		if err != nil {
			c.JSON(400, Res{Code:400,Message:"生成token失败"})
			return
		}

    var userLog = &models.UserLog{
      UserID:user.UserId,
      IP: c.ClientIP(),
      UA: c.GetHeader("user-agent"),
    }
    log.Println(c.GetHeader("user-agent"))
    _, _ = userLog.Create()

		result.Token = token
		c.JSON(200, Res{Code:0,Message:"登录成功",Data:result})
	}

}

// @Summary 获取用户信息
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":info,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /user/info [get]
func UserInfo(c * gin.Context) {
	var userModel models.User
	userId, bol:= c.Get("_user_id")
  if !bol {
    c.JSON(401, Res{Code:401,Message:"未登录"})
    return
  }
  userRole, err := userModel.GetUserByUserId(userId.(int64))
	if err != nil {
		c.JSON(400, Res{Code:400,Message:"获取失败"})
		return
	} else {
		var (
      roles   []string
      menuIds []int64
      buttons []string
      menuModel models.Menu
    )
		for _, role := range userRole.Roles {
			roles = append(roles,role.RoleName)
			var (
        _menuIds []int64
        _buttons []models.Button
      )
			_ = json.Unmarshal([]byte(role.MenuIds),&_menuIds)
			menuIds = append(menuIds,_menuIds...)
			_ = json.Unmarshal([]byte(role.Buttons),&_buttons)
			for _, _button := range _buttons {
				buttons = append(buttons,_button.Btns...)
			}
		}

		_menuIds, _ := utils.RemoveDuplicateElement(menuIds)
		_buttons, _ := utils.RemoveDuplicateElement(buttons)

		menus, err := menuModel.GetTreeMenuByIds(_menuIds.([]int64))
		if err != nil {
			log.Println(err.Error())
			c.JSON(400, Res{Code:400,Message:"获取失败"})
			return
		}

		userInfo := &models.Info{
			Avatar: "https://s1.ax1x.com/2020/05/25/tp7UWF.gif",
			Buttons: _buttons.([]string),
			Menus: menus,
			Roles:roles,
			Name: c.MustGet("_user_name").(string),
			UserID: userId.(int64),
		}

		c.JSON(200,Res{Code:0,Message:"获取成功",Data:userInfo})
		return
	}
}

// @Summary 添加用户
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param user body userCreateForm true "账号信息"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"添加成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"添加失败"}
// @Router /user [post]
//]
func CreateUser(c *gin.Context)  {
  var (
    userForm userCreateForm
    userModel models.UserView
  )
  if err := c.BindJSON(&userForm); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  if usern, _ := userModel.GetUserByUserName(userForm.UserName); usern.UserId != 0 {
    c.JSON(400,Res{Code:400,Message:"用户名已存在"})
    return
  }

  for _, roleId := range userForm.RoleIds {
    role := models.Role{
      RoleID:     roleId,
    }
    userModel.Roles = append(userModel.Roles,role)
  }

  userModel.UserName = userForm.UserName
  userModel.Status   = userForm.Status
  userModel.Password = utils.MD5(userForm.Password)
  if err := userModel.Create(); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"添加失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"添加成功"})
}

// @Summary 修改用户
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param user body userUpdateForm true "账号信息"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"修改成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"修改失败"}
// @Router /user/:user_id [put]
//]
func UpdateUser(c *gin.Context)  {
  var (
    userForm userUpdateForm
    userModel models.User
  )
  if err := c.BindJSON(&userForm); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  userId, err := strconv.ParseInt(c.Param("user_id"),10,64)
  if  err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  if user, err := userModel.GetUserByUserId(userId); user.UserId == 0 || err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"用户不存在"})
    return
  }

  if usern, err := userModel.GetUserByUserName(userForm.UserName);err == nil && userId != usern.UserId {
    c.JSON(400,Res{Code:400,Message:"用户名已存在"})
    return
  }

  for _, roleId := range userForm.RoleIds {
    role := models.Role{
      RoleID:     roleId,
    }
    userModel.Roles = append(userModel.Roles,role)
  }

  userModel.UserId = userId
  userModel.UserName = userForm.UserName
  userModel.Status = userForm.Status
  if err := userModel.Update(); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"修改失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"修改成功"})
}

// @Summary 修改用户密码
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param user body userUpdatePwdForm true "账号密码信息"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"修改成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"修改失败"}
// @Router /user/pwd/:user_id [put]
//]
func UpdatePwdUser(c *gin.Context)  {
  var (
    userPwdForm userUpdatePwdForm
    userModel models.User
  )
  if err := c.BindJSON(&userPwdForm); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  userId, err := strconv.ParseInt(c.Param("user_id"),10,64)
  if  err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }
  user, err := userModel.GetUserByUserId(userId)
  if err != nil || user.UserId == 0 {
    c.JSON(400,Res{Code:400,Message:"用户不存在"})
    return
  }

  if user.Password != utils.MD5(userPwdForm.OldPassword) {
    c.JSON(400,Res{Code:400,Message:"原密码不正确"})
    return
  }

  userModel.UserId = userId
  newPassword := utils.MD5(userPwdForm.Password)
  if err := userModel.UpdatePwd(newPassword); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"修改失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"修改成功"})
}

// @Summary 删除用户
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param user_id body []int64 true "用户ID数组"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"删除成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"删除失败"}
// @Router /users [delete]
//]
func DeleteUser(c *gin.Context)  {
  var userIds []int64
  if err := c.BindJSON(&userIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  var (
    userModel models.User
  )
  if err := userModel.Delete(userIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"删除成功"})
}

// @Summary 获取用户列表
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":users,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /users [get]
func UserList(c *gin.Context)  {
  var (
    userModel models.User
    user      [] models.User
    count     int64
    err       error
  )

  pageSize, _ := strconv.Atoi(c.DefaultQuery("limit","20"))
  pageIndex, _ := strconv.Atoi(c.DefaultQuery("page","1"))
  userName := c.DefaultQuery("user_name","")
  status, _ := strconv.Atoi(c.DefaultQuery("status","-1"))
  if user, count, err = userModel.GetUserPage(pageSize,pageIndex,userName,status); err != nil {
    c.JSON(400,Res{Code:400,Message:"获取失败"})
    return
  }

  type users struct {
    Users  []models.User   `json:"users"`
    Total int64            `json:"total"`
  }

  c.JSON(200,Res{Code:0,Message:"获取成功",Data:&users{Users:user,Total:count}})
}

// @Summary 获取登录日志列表
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":logs,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /user/logs [get]
func UserLogList(c *gin.Context)  {
  var (
    userLogModel models.UserLog
    userLog [] models.UserLog
    count  int64
    err    error
  )

  pageSize, _ := strconv.Atoi(c.DefaultQuery("limit","20"))
  pageIndex, _ := strconv.Atoi(c.DefaultQuery("page","1"))
  date := c.QueryArray("date[]")
  if userLog, count, err = userLogModel.GetUserLogPage(pageSize,pageIndex,date); err != nil {
    c.JSON(400,Res{Code:400,Message:"获取失败"})
    return
  }
  type logs struct {
    Logs  []models.UserLog `json:"logs"`
    Total int64            `json:"total"`
  }
  c.JSON(200,Res{Code:0,Message:"获取成功",Data:&logs{Logs:userLog,Total:count}})
}

// @Summary 删除登录日志
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param user_log_id body []int64 true "日志ID数组"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"删除成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"删除失败"}
// @Router /user/logs [delete]
//]
func DeleteUserLog(c *gin.Context)  {
  var userLogIds []int64
  if err := c.BindJSON(&userLogIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  var (
   userLogModel models.UserLog
  )
  if err := userLogModel.Delete(userLogIds); err != nil {
   c.JSON(400,Res{Code:400,Message:"删除失败"})
   return
  }

  c.JSON(200,Res{Code:0,Message:"删除成功"})
}


