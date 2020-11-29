package models

import (
  orm "go-element-admin-api/db"
  "go-element-admin-api/utils"
  "log"
)

type BlogArticleTag struct {
  ArticleId  int64  `gorm:"column:article_id;" json:"article_id"`
  TagId      int64  `gorm:"column:tag_id;" json:"tag_id"`
  CreateTime string `gorm:"column:create_time" json:"create_time"`
}

// 创建文章标签
func (at BlogArticleTag) Create() (err error) {
  at.CreateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("blog_article_tags").Create(&at).Error; err != nil {
    log.Println(err.Error())
  }
  return
}

// 删除文章标签
func (at BlogArticleTag) Delete(articleId int64,tagIds []int64)(err error)  {
  if err = orm.Eloquent.Table("blog_article_tags").Where("article_id = ?",articleId).Where("tag_id in (?)",tagIds).Delete(&at).Error; err != nil {
    log.Println(err.Error())
  }
  return
}
