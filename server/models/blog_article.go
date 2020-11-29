package models

import (
  "github.com/jinzhu/gorm"
  orm "go-element-admin-api/db"
  "go-element-admin-api/utils"
  "log"
)

type BlogArticle struct {
  ArticleID       int64   `gorm:"column:article_id;primary_key;AUTO_INCREMENT" json:"article_id"`
  ArticleAuthor   string  `gorm:"column:article_author;default:'bigfool';" json:"article_author" binding:""`
  ArticleTitle    string  `gorm:"column:article_title;" json:"article_title" binding:"required"`
  ArticleFace     string  `gorm:"column:article_face;" json:"article_face" binding:""`
  ArticleDesc     string  `gorm:"column:article_desc;" json:"article_desc" binding:""`
  ArticleStatus   int     `gorm:"column:article_status;default:1" json:"article_status" binding:""`
  ArticleContent  string  `gorm:"column:article_content;" json:"article_content" binding:"required"`
  ArticleRead     int64   `gorm:"column:article_read;" json:"article_read"`
  UpdateTime      string  `gorm:"column:update_time;default:NULL" json:"update_time"`
  CreateTime      string  `gorm:"column:create_time" json:"create_time"`

  Tags            []BlogTag `gorm:"association_autoupdate:false;many2many:blog_article_tags;foreignkey:article_id;association_foreignkey:tag_id;association_jointable_foreignkey:tag_id;jointable_foreignkey:article_id;" json:"tags"`

}

// 获取文章
func (a BlogArticle) GetBlogArticle() (article BlogArticle, err error)  {
  if err = orm.Eloquent.Table("blog_articles").Preload("Tags").Where(&a).First(&article).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  return
}

// 根据id 获取文章
func (a BlogArticle) GetBlogArticleById(articleId int64) (article BlogArticle, err error)  {
  if err = orm.Eloquent.Table("blog_articles").Preload("Tags").Take(&article,articleId).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  return
}

// 创建接口
func (a BlogArticle) Create() (err error)  {
  a.CreateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("blog_articles").Create(&a).Error; err != nil {
    log.Println(err.Error())
  }
  return
}

// 修改文章
func (a BlogArticle) Update() (err error) {
  var (
    oldTagIds []int64
    newTagIds []int64
    blogArticleTag BlogArticleTag
    article BlogArticle
  )

  a.UpdateTime = utils.GetCurrntTime()
  tran := orm.Eloquent.Begin()

  if article, err = a.GetBlogArticleById(a.ArticleID); err != nil {
    log.Println(err.Error())
    tran.Rollback()
    return
  }

  for _, t := range article.Tags {
    oldTagIds = append(oldTagIds,t.TagID)
  }
  for _, t := range a.Tags {
    newTagIds = append(newTagIds,t.TagID)
  }

  // 删除取消的标签
  delTagIds, err := utils.Difference(oldTagIds,newTagIds)
  if err != nil {
    log.Println(err.Error())
    tran.Rollback()
    return err
  }
  if len(delTagIds.([]int64)) > 0 {
    if err := blogArticleTag.Delete(a.ArticleID,delTagIds.([]int64)); err != nil {
      log.Println(err.Error())
      tran.Rollback()
      return err
    }
  }

  // 添加新的标签
  addTagIds, err := utils.Difference(newTagIds,oldTagIds)
  if err != nil {
    log.Println(err.Error())
    tran.Rollback()
    return err
  }
  if len(addTagIds.([]int64)) > 0 {
    for _, tagId := range addTagIds.([] int64) {
      blogArticleTag = BlogArticleTag{
        ArticleId:     a.ArticleID,
        TagId:         tagId,
      }
      if err := blogArticleTag.Create(); err != nil {
        log.Println(err.Error())
        tran.Rollback()
        return err
      }
    }
  }

  if err = orm.Eloquent.Table("blog_articles").Omit("create_time").Save(&a).Error; err != nil {
    log.Println(err.Error())
    tran.Rollback()
    return
  }

  err = tran.Commit().Error
  return
}


// 删除文章
func (a BlogArticle) Delete(articleIds []int64)(err error)  {
  tran := orm.Eloquent.Begin()
  if err = orm.Eloquent.Table("blog_articles").Where("article_id in (?)",articleIds).Delete(&a).Error; err != nil {
    tran.Rollback()
    log.Println(err.Error())
    return
  }

  if err = orm.Eloquent.Table("blog_article_tags").Where("article_id in (?)",articleIds).Delete(&a).Error; err != nil {
    tran.Rollback()
    log.Println(err.Error())
    return
  }
  err = tran.Commit().Error
  return
}

// 获取列表
func (a BlogArticle) GetBlogArticlePage(pageSize int, pageIndex int, articleTitle string, articleStatus int, tagId int64) (blogArticles []BlogArticle,count int64,err error)  {
  table := orm.Eloquent.Table("blog_articles")
  if articleTitle != "" {
    table = table.Where("blog_articles.article_title = ?",articleTitle)
  }
  if articleStatus != -1 {
    table = table.Where("blog_articles.article_status = ?",articleStatus)
  }

  if tagId != -1 {
    table = table.Where("article_id IN (SELECT article_id FROM blog_article_tags WHERE tag_id = ?)",tagId)
  }
  table = table.Preload("Tags")

  table = table.Select([]string{"article_id","article_author","article_title","article_desc","article_status","article_read","update_time","create_time"})

  if err = table.Offset((pageIndex -1) * pageSize).Limit(pageSize).Order("blog_articles.create_time desc").Find(&blogArticles).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  table.Count(&count)
  return
}

