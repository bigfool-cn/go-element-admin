package models

import (
  orm "go-element-admin/db"
  "go-element-admin/utils"
)

type UserRole struct {
  UserRoleID  int64  `gorm:"column:user_role_id;primary_key;AUTO_INCREMENT" json:"user_role_id"`
  UserId      int64  `gorm:"column:user_id;" json:"user_id"`
  RoleId      int64  `gorm:"column:role_id;" json:"role_id"`
  CreateTime  string `gorm:"column:create_time" json:"create_time"`
}

// 创建用户角色
func (ur UserRole) Create() (id int64, err error) {
  ur.CreateTime = utils.GetCurrntTime()
  err = orm.Eloquent.Table("user_roles").Create(&ur).Error
  id = ur.UserRoleID
  return
}


// 删除用户角色
func (ur UserRole) Delete(userId int64,roleIds []int64)(err error)  {
  err = orm.Eloquent.Table("user_roles").Where("user_id = ?",userId).Where("role_id in (?)",roleIds).Delete(&ur).Error
  return
}

