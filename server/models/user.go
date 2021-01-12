package models

import (
  "github.com/jinzhu/gorm"
  orm "go-element-admin/db"
  "go-element-admin/utils"
)

type User struct {
	UserId       int64  `gorm:"column:user_id;primary_key;AUTO_INCREMENT" json:"user_id"`
	UserName     string `gorm:"column:user_name" json:"user_name"`

	Status       int    `gorm:"column:status" json:"status"`
	UpdateTime   string `gorm:"column:update_time;default:NULL" json:"update_time"`
	CreateTime   string `gorm:"column:create_time" json:"create_time"`

  Roles []Role `gorm:"association_autoupdate:false;many2many:user_roles;foreignkey:user_id;association_foreignkey:role_id;association_jointable_foreignkey:role_id;jointable_foreignkey:user_id;" json:"roles"`

}

type UserView struct {
  User
  Password     string `gorm:"column:password" json:"password"`
}

type Info struct {
	Avatar   string      `json:"avatar"`
	Buttons  []string    `json:"buttons"`
	Menus    []TreeMenu  `json:"menus"`
	Roles    []string    `json:"roles"`
	Name     string      `json:"name"`
	UserID   int64       `json:"user_id"`
}

func (u UserView) Create() (err error) {

  tran := orm.Eloquent.Begin()

  u.CreateTime = utils.GetCurrntTime()

  if err = orm.Eloquent.Table("users").Create(&u).Error; err != nil {
    tran.Rollback()
    return
  }

  err = tran.Commit().Error
  return
}

func (u User) Update() (err error) {
  var (
    oldRoleIds []int64
    newRoleIds []int64
    userRoleModel UserRole
    user User
  )

  u.UpdateTime = utils.GetCurrntTime()
  tran := orm.Eloquent.Begin()

  if user, err = u.GetUserRoleByUserId(u.UserId); err != nil {
    tran.Rollback()
    return
  }

  for _, us := range user.Roles {
    oldRoleIds = append(oldRoleIds,us.RoleID)
  }
  for _, us := range u.Roles {
    newRoleIds = append(newRoleIds,us.RoleID)
  }

  // 删除取消的用户角色
  delRoleIds, err := utils.Difference(oldRoleIds,newRoleIds)
  if err != nil {
    tran.Rollback()
    return err
  }
  if len(delRoleIds.([]int64)) > 0 {
    if err := userRoleModel.Delete(u.UserId,delRoleIds.([]int64)); err != nil {
      tran.Rollback()
      return err
    }
  }

  // 添加新的用户角色
  addRoleIds, err := utils.Difference(newRoleIds,oldRoleIds)
  if err != nil {
    tran.Rollback()
    return err
  }
  if len(addRoleIds.([]int64)) > 0 {
    for _, roleId := range addRoleIds.([] int64) {
      userRoleModel = UserRole{
        UserId:     u.UserId,
        RoleId:     roleId,
      }
      if userRoleId, err := userRoleModel.Create(); userRoleId <=0 && err != nil {
        tran.Rollback()
        return err
      }
    }
  }

  if err = orm.Eloquent.Table("users").Omit("create_time").Save(&u).Error; err != nil {
    tran.Rollback()
    return
  }

  err = tran.Commit().Error
  return
}

func (u User) UpdatePwd(password string) (err error) {
  err = orm.Eloquent.Table("users").Where("user_id = ?",u.UserId).Updates(map[string]interface{}{"password":password}).Error
  return
}

// 删除用户
func (u User) Delete(userIds []int64) (err error)  {
  tran := orm.Eloquent.Begin()
  if err = orm.Eloquent.Table("users").Where("user_id in (?)",userIds).Delete(&u).Error; err != nil {
    tran.Rollback()
    return
  }
  if err = orm.Eloquent.Table("user_logs").Where("user_id in (?)",userIds).Delete(&u).Error; err != nil {
    tran.Rollback()
    return
  }
  if err = orm.Eloquent.Table("user_roles").Where("user_id in (?)",userIds).Delete(&u).Error; err != nil {
    tran.Rollback()
    return
  }
  err = tran.Commit().Error
  return
}

//获取用户
func (u User) GetUser() (user UserView, err error)  {
	if err = orm.Eloquent.Table("users").Preload("Roles").Where(&u).Take(&user).Error; err != nil{
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

// 根据user_id 获取用户
func (u User) GetUserByUserId(userId int64) (user UserView, err error)  {
	if err = orm.Eloquent.Table("users").Preload("Roles").Take(&user,userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

// 根据user_name获取用户
func (u User) GetUserByUserName(useName string) (user UserView, err error)  {
	if err = orm.Eloquent.Table("users").Preload("Roles").Where("user_name = ?",useName).Take(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

func (u User) GetUserRoleByUserId(userId int64) (user User, err error)  {
	if err = orm.Eloquent.Model(user).Preload("Roles").Take(&user,userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

// 获取用户列表
func (u User) GetUserPage(pageSize int, pageIndex int, userName string, status int) (users []User,count int64,err error)  {
  table := orm.Eloquent.Model(users).Preload("Roles")
  if userName != "" {
    table = table.Where("users.user_name = ?",userName)
  }
  if status != -1 {
    table = table.Where("users.status = ?",status)
  }
  if err = table.Offset((pageIndex -1) * pageSize).Limit(pageSize).Order("users.create_time desc").Find(&users).Error; err != nil {
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  table.Count(&count)
  return
}
