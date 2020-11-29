package models

import (
  "encoding/json"
  "fmt"
  "github.com/jinzhu/gorm"
  orm "go-element-admin-api/db"
  "go-element-admin-api/utils"
  "go-element-admin-api/utils/gcasbin"
  "log"
  "strings"
)

type Role struct {
	RoleID       int64   `gorm:"column:role_id;primary_key;AUTO_INCREMENT" json:"role_id"`
	RoleName     string  `gorm:"column:role_name;primary_key" json:"role_name"`
	Remark       string  `gorm:"column:remark" json:"remark"`
	Status       int     `gorm:"column:status" json:"status"`
  PathIds      string  `gorm:"column:path_ids" json:"path_ids"`
	MenuIds      string  `gorm:"column:menu_ids" json:"menu_ids"`
	Buttons      string  `gorm:"column:buttons" json:"buttons"`
	UpdateTime   string  `gorm:"column:update_time;default:NULL" json:"update_time"`
	CreateTime   string  `gorm:"column:create_time" json:"create_time"`
}

type Button struct {
	Btns   []string  `json:"btns"`
	MenuID int64     `json:"menu_id"`
}

type RoleView struct {
  RoleID       int64    `json:"role_id" binding:""`
  Buttons      []Button `json:"buttons" binding:""`
  MenuIds      []int64  `json:"menu_ids"  binding:""`
  PathIds      []int64  `json:"path_ids"  binding:""`
  Remark       string   `json:"remark" binding:""`
  RoleName     string   `json:"role_name" binding:"required"`
  Status       int      `json:"status"  binding:""`
  UpdateTime   string  `json:"update_time"`
  CreateTime   string  `json:"create_time"`
}

//获取用户
func (r Role) GetRole() (RoleView, error)  {
  var (
    role Role
    err error
  )
  if err = orm.Eloquent.Table("roles").Where(&r).Take(&role).Error; err != nil{
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  var roleView = roleToRoleView(role)
  return roleView,err
}

// 新增角色
func (r Role) Create() (roleId int64, err error) {
  r.CreateTime = utils.GetCurrntTime()
  tran := orm.Eloquent.Begin()
  if err = orm.Eloquent.Table("roles").Create(&r).Error; err != nil {
    tran.Rollback()
    log.Println(err.Error())
  }
  if len(r.PathIds) > 0 {
    var (
      pathModel Path
      pathIds   []int64
    )
    if err := json.Unmarshal([] byte(r.PathIds),&pathIds);err != nil {
      log.Println(err.Error())
      tran.Rollback()
    }
    paths, err := pathModel.GetPathByIDs(pathIds)
    if err == nil && len(paths) > 0 {
      var stringBuild strings.Builder
      stringBuild.WriteString("INSERT INTO casbin_rule (`p_type`,`v0`,`v1`,`v2`) VALUES")
      for idx, path := range paths  {
        if path.Type == "J" {
          if len(paths)-1 == idx {
            //最后一条数据 以分号结尾
            stringBuild.WriteString(fmt.Sprintf("('p','%d','%s','%s');", r.RoleID,path.Path,path.Method))
          } else {
            stringBuild.WriteString(fmt.Sprintf("('p','%d','%s','%s'),", r.RoleID,path.Path,path.Method))
          }
        }
      }
      sql := stringBuild.String()
      orm.Eloquent.Exec(sql)
    }
  }

  err = tran.Commit().Error
  gcasbin.Enforcer.LoadPolicy()
  roleId = r.RoleID
  return
}

// 修改角色
func (r Role) Update() (err error) {
  r.UpdateTime = utils.GetCurrntTime()
  tran := orm.Eloquent.Begin()
  if err = orm.Eloquent.Table("roles").Omit("create_time").Save(&r).Error; err != nil {
    log.Println(err.Error())
    tran.Rollback()
    return
  }
  if len(r.PathIds) > 0 {
    var (
      pathModel Path
      pathIds   []int64
    )
    if err := json.Unmarshal([] byte(r.PathIds),&pathIds);err != nil {
      log.Println(err.Error())
      tran.Rollback()
    }
    paths, err := pathModel.GetPathByIDs(pathIds)
    if err == nil && len(paths) > 0 {
      var stringBuild strings.Builder
      stringBuild.WriteString("INSERT INTO casbin_rule (`p_type`,`v0`,`v1`,`v2`) VALUES")
      for idx, path := range paths  {
        if path.Type == "J" {
          if len(paths)-1 == idx {
            //最后一条数据 以分号结尾
            stringBuild.WriteString(fmt.Sprintf("('p','%d','%s','%s');", r.RoleID,path.Path,path.Method))
          } else {
            stringBuild.WriteString(fmt.Sprintf("('p','%d','%s','%s'),", r.RoleID,path.Path,path.Method))
          }
        }
      }
      orm.Eloquent.Exec("DELETE FROM casbin_rule WHERE v0 = ?",r.RoleID)
      sql := stringBuild.String()
      orm.Eloquent.Exec(sql)
    }
  }
  err = tran.Commit().Error
  gcasbin.Enforcer.LoadPolicy()
  return
}

// 删除角色
func (r Role) Delete(roleIds []int64)(err error)  {
  tran := orm.Eloquent.Begin()
  if err = orm.Eloquent.Table("roles").Where("role_id in (?)",roleIds).Delete(&r).Error; err != nil {
    log.Println(err.Error())
    tran.Rollback()
    return
  }
  if err = orm.Eloquent.Table("user_roles").Where("role_id in (?)",roleIds).Delete(&r).Error; err != nil {
    log.Println(err.Error())
    tran.Rollback()
    return
  }
  orm.Eloquent.Exec("DELETE FROM casbin_rule WHERE v0 IN (?)",roleIds)
  err = tran.Commit().Error
  gcasbin.Enforcer.LoadPolicy()
  return
}

// 获取用户列表
func (r Role) GetRolePage(pageSize int, pageIndex int, roleName string, status int) ([]RoleView, int64, error) {
  var (
    roles []Role
    count int64
    err error
    rolesView []RoleView
  )

  table := orm.Eloquent.Table("roles").Model(roles)

  if roleName != "" {
    table = table.Where("role_name = ?",roleName)
  }
  if status != -1 {
    table = table.Where("status = ?",status)
  }
  if err = table.Offset((pageIndex -1) * pageSize).Limit(pageSize).Order("create_time desc").Find(&roles).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }

  table.Count(&count)

  for _, role := range roles {
    var roleView  = roleToRoleView(role)
    rolesView = append(rolesView,roleView)
  }
  return rolesView,count,err
}

func roleToRoleView(role Role) RoleView {
  var (
    pathIds   []int64
    menuIds   []int64
    buttons   []Button
  )
  _ = json.Unmarshal([]byte(role.PathIds),&pathIds)
  _ = json.Unmarshal([]byte(role.MenuIds),&menuIds)
  _ = json.Unmarshal([]byte(role.Buttons),&buttons)
  var roleView = RoleView{
    RoleID:     role.RoleID,
    PathIds:    pathIds,
    MenuIds:    menuIds,
    Buttons:    buttons,
    Remark:     role.Remark,
    RoleName:   role.RoleName,
    Status:     role.Status,
    UpdateTime: role.UpdateTime,
    CreateTime: role.CreateTime,
  }
  return roleView
}

