package apis

import (
  "encoding/json"
  "github.com/gin-gonic/gin"
  "go-element-admin-api/models"
  "log"
  "strconv"
)

// @Summary 添加角色
// @Tags 角色管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param role body roleForm true "角色数据"
// @Success 200 {object} Res {"code":0,"msg":"添加成功"}
// @Failure 400 {object} Res {"code":400,"msg":"msg"}
// @Router /role [post]
func CreateRole (c *gin.Context){
  var role roleForm
  if err := c.BindJSON(&role); err != nil {
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  pathIds, _ := json.Marshal(role.PathIds)
  menuIds, _ := json.Marshal(role.MenuIds)
  buttons, _ := json.Marshal(role.Buttons)
  var roleModel = models.Role{
    RoleName:   role.RoleName,
    Remark:     role.Remark,
    Status:     role.Status,
    PathIds:    string(pathIds),
    MenuIds:    string(menuIds),
    Buttons:    string(buttons),
  }
  if _,err := roleModel.Create(); err != nil {
    c.JSON(400,Res{Code:400,Message:"添加失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"添加成功"})
}

// @Summary 修改角色
// @Tags 角色管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param role body roleForm true "角色数据"
// @Success 200 {object} Res {"code":0,"msg":"修改成功"}
// @Failure 400 {object} Res {"code":400,"msg":"msg"}
// @Router /role/:role_id [put]
func UpdateRole (c *gin.Context){
  var role roleForm
  if err := c.BindJSON(&role); err != nil {
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  roleId, err := strconv.ParseInt(c.Param("role_id"),10,64)
  if err != nil {
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  pathIds, _ := json.Marshal(role.PathIds)
  menuIds, _ := json.Marshal(role.MenuIds)
  buttons, _ := json.Marshal(role.Buttons)
  var roleModel = models.Role{
    RoleID:     roleId,
    RoleName:   role.RoleName,
    Remark:     role.Remark,
    Status:     role.Status,
    PathIds:    string(pathIds),
    MenuIds:    string(menuIds),
    Buttons:    string(buttons),
  }
  if err := roleModel.Update(); err != nil {
    c.JSON(400,Res{Code:400,Message:"修改失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"修改成功"})
}

// @Summary 删除角色
// @Tags 角色管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param role_id body []int64 true "日志ID数组"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"删除成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"删除失败"}
// @Router /roles [delete]
//]
func DeleteRole(c *gin.Context)  {
  var roleIds []int64
  if err := c.BindJSON(&roleIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  var (
    roleModel models.Role
  )
  if err := roleModel.Delete(roleIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"删除成功"})
}

// @Summary 获取角色列表
// @Tags 角色管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":roles,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /roles [get]
func RoleList(c *gin.Context)  {
  var (
    roleModel models.Role
    role      []models.RoleView
    count     int64
    err       error
  )

  pageSize, _ := strconv.Atoi(c.DefaultQuery("limit","20"))
  pageIndex, _ := strconv.Atoi(c.DefaultQuery("page","1"))
  roleName := c.DefaultQuery("role_name","")
  status, _ := strconv.Atoi(c.DefaultQuery("status","-1"))
  if role, count, err = roleModel.GetRolePage(pageSize,pageIndex,roleName,status); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"获取失败"})
    return
  }

  type roles struct {
    Roles  []models.RoleView   `json:"roles"`
    Total int64            `json:"total"`
  }

  c.JSON(200,Res{Code:0,Message:"获取成功",Data:&roles{Roles:role,Total:count}})
}
