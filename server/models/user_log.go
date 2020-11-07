package models

import (
  orm "element-admin-api/db"
  "element-admin-api/utils"
  "github.com/jinzhu/gorm"
  "log"
)

type UserLog struct {
  UserLogID  int64  `gorm:"column:user_log_id;primary_key;AUTO_INCREMENT" json:"user_log_id"`
  UserID     int64  `json:"user_id"`
  IP         string `json:"ip"`
  UA         string `json:"ua"`
  CreateTime string `json:"create_time"`
  User       User   `gorm:"foreignKey:user_id;association_foreignkey:user_id" json:"user"`
}

// 创建日志
func (ul UserLog) Create() (id int64, err error) {
  ul.CreateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("user_logs").Create(&ul).Error; err != nil {
   log.Println(err.Error())
  }
  id = ul.UserLogID
  return
}

// 删除日志
func (ul UserLog) Delete(userLogIds []int64)(err error)  {
  if err = orm.Eloquent.Table("user_logs").Where("user_log_id in (?)",userLogIds).Delete(&ul).Error; err != nil {
    log.Println(err.Error())
  }
  return
}

// 获取日志列表
func (ul UserLog) GetUserLogPage(pageSize int, pageIndex int, date []string) (userLogs []UserLog,count int64,err error)  {
  table := orm.Eloquent.Model(userLogs).Preload("User")
  if len(date) > 0 && date[0] != "" {
    table = table.Where("create_time >= ?",date[0])
  }
  if len(date) > 0 && date[1] != "" {
    table = table.Where("create_time <= ?",date[1])
  }
  if err = table.Offset((pageIndex -1) * pageSize).Limit(pageSize).Order("user_logs.create_time desc").Find(&userLogs).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  table.Count(&count)
  return
}
