package models

import (
  "github.com/jinzhu/gorm"
  orm "go-element-admin-api/db"
  "go-element-admin-api/utils"
  "log"
)

type BlogTag struct {
  TagID       int64   `gorm:"column:tag_id;primary_key;AUTO_INCREMENT" json:"tag_id"`
  TagTitle    string  `gorm:"column:tag_title;unique" json:"tag_title" binding:"required"`
  TagStatus   int     `gorm:"column:tag_status;" json:"tag_status" binding:""`
  UpdateTime  string  `gorm:"column:update_time;default:NULL" json:"update_time"`
  CreateTime  string  `gorm:"column:create_time" json:"create_time"`
}

// 获取接口
func (t BlogTag) GetBlogTag() (tag BlogTag, err error)  {
  if err = orm.Eloquent.Table("blog_tags").Where(&t).First(&tag).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  return
}

// 创建接口
func (t BlogTag) Create() (tagId int64, err error)  {
  t.CreateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("blog_tags").Create(&t).Error; err != nil {
    log.Println(err.Error())
  }
  tagId = t.TagID
  return
}

// 修改接口
func (t BlogTag) Update() (err error) {
  t.UpdateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("blog_tags").Omit("create_time").Save(&t).Error; err != nil {
    log.Println(err.Error())
  }
  return
}


// 删除接口
func (t BlogTag) Delete(tagIds []int64)(err error)  {
  tran := orm.Eloquent.Begin()

  if err = orm.Eloquent.Table("blog_tags").Where("tag_id in (?)",tagIds).Delete(&t).Error; err != nil {
    tran.Rollback()
    log.Println(err.Error())
    return
  }

  if err = orm.Eloquent.Table("blog_article_tags").Where("tag_id in (?)",tagIds).Delete(&t).Error; err != nil {
    tran.Rollback()
    log.Println(err.Error())
    return
  }

  err = tran.Commit().Error
  return
}

// 获取列表
func (t BlogTag) GetBlogTagPage(pageSize int, pageIndex int, tagTitle string, tagStatus int) (blogTags []BlogTag,count int64,err error)  {
  table := orm.Eloquent.Table("blog_tags")
  if tagTitle != "" {
    table = table.Where("blog_tags.tag_title = ?",tagTitle)
  }
  if tagStatus != -1 {
    table = table.Where("blog_tags.tag_status = ?",tagStatus)
  }
  if err = table.Offset((pageIndex -1) * pageSize).Limit(pageSize).Order("blog_tags.create_time desc").Find(&blogTags).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  table.Count(&count)
  return
}
