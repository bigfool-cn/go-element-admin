package models

import (
  orm "go-element-admin-api/db"
  "go-element-admin-api/utils"
  "log"
)

type UserRole struct {
  UserId      int64  `gorm:"column:user_id;" json:"user_id"`
  RoleId      int64  `gorm:"column:role_id;" json:"role_id"`
  CreateTime  string `gorm:"column:create_time" json:"create_time"`
}

// 创建用户角色
func (ur UserRole) Create() (err error) {
  ur.CreateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("user_roles").Create(&ur).Error; err != nil {
    log.Println(err.Error())
  }
  return
}


// 删除用户角色
func (ur UserRole) Delete(userId int64,roleIds []int64)(err error)  {
  if err = orm.Eloquent.Table("user_roles").Where("user_id = ?",userId).Where("role_id in (?)",roleIds).Delete(&ur).Error; err != nil {
    log.Println(err.Error())
  }
  return
}

