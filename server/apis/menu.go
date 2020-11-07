package apis

import (
  "element-admin-api/models"
  "github.com/gin-gonic/gin"
  "log"
  "strconv"
)

// @Summary 获取菜单
// @Tags 菜单管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":menu,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /menu/:menu_id [get]
func GetMenu(c *gin.Context)  {
  var menuModel models.Menu

  menuId, err := strconv.ParseInt(c.Param("menu_id"),10,64)
  if err != nil {
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  menuModel.MenuID = menuId

  menu, err := menuModel.GetMenu()
  if err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"获取失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"获取成功",Data:menu})
}


// @Summary 添加菜单
// @Tags 菜单管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param menu body menuForm true "菜单数据"
// @Success 200 {object} Res {"code":0,"msg":"添加成功"}
// @Failure 400 {object} Res {"code":400,"msg":"msg"}
// @Router /menu [post]
func CreateMenu (c *gin.Context){
  var menuModel models.Menu

  if err := c.BindJSON(&menuModel); err != nil {
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  if _,err := menuModel.Create(); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"添加失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"添加成功"})
}

// @Summary 修改菜单
// @Tags 菜单管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param menu body menuForm true " 菜单数据"
// @Success 200 {object} Res {"code":0,"msg":"修改成功"}
// @Failure 400 {object} Res {"code":400,"msg":"msg"}
// @Router /menu/:menu_id [put]
func UpdateMenu (c *gin.Context){
  var menuModel models.Menu

  if err := c.BindJSON(&menuModel); err != nil {
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  menuId, err := strconv.ParseInt(c.Param("menu_id"),10,64)
  if err != nil {
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  menuModel.MenuID = menuId

  if err := menuModel.Update(); err != nil {
    c.JSON(400,Res{Code:400,Message:"修改失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"修改成功"})
}

// @Summary 删除菜单
// @Tags 菜单管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param menu_id body []int64 true "菜单ID数组"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"删除成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"删除失败"}
// @Router /menus [delete]
//]
func DeleteMenu(c *gin.Context)  {
  var menuIds []int64
  if err := c.BindJSON(&menuIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  var (
    menuModel models.Menu
  )
  if err := menuModel.Delete(menuIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"删除成功"})
}

// @Summary 菜单列表
// @Tags 菜单管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":menus,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /menus [get]
func MenuList(c *gin.Context)  {
  var menuModel models.Menu
  menuModel.Title = c.Query("title")
  menuModel.Hidden,_ = strconv.Atoi(c.Query("hidden"))

  menus, err := menuModel.GetTreeMenus()
  if err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"获取数据失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"获取成功",Data:menus})
}
