package apis

import "go-element-admin-api/models"

type Res struct {
  Code    int `json:"code"`
  Data    interface{} `json:"data"`
  Message string `json:"message"`
}

// 登录表单
type loginForm struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// 登录账号信息
type info struct {
  Avatar   string      `json:"avatar"`
  Buttons  []string    `json:"buttons"`
  Menus    []models.TreeMenu  `json:"menus"`
  Roles    []string    `json:"roles"`
  UserID   int64       `json:"user_id"`
}

// 角色表单
type roleForm struct {
  models.RoleView
}

// 账号信息表单--添加
type userCreateForm struct {
  userUpdateForm
  passwordToRe
}

// 账号信息表单--修改
type userUpdateForm struct {
  UserName  string `json:"user_name" binding:"required"`
  Status    int    `json:"status" binding:"numeric"`
  RoleIds   []int64 `json:"role_ids" binding:"required"`
}

// 账号密码表单--修改
type userUpdatePwdForm struct {
  OldPassword string `json:"old_password" binding:"required"`
  passwordToRe
}

type passwordToRe struct {
  Password    string `json:"password" binding:"required"`
  Repassword  string `json:"repassword" binding:"required,eqfield=Password"`
}

type menuForm struct {
  models.Menu
}

type pathForm struct {
  models.Path
}

type blogTagForm struct {
  models.BlogTag
}

type blogArticleForm struct {
  models.BlogArticle
  TagIds          []int64 `json:"tag_ids" binding:"required"`
}


