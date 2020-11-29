package apis

import (
  "github.com/gin-gonic/gin"
  "go-element-admin-api/models"
  "log"
  "strconv"
)

// @Summary 获取文章
// @Tags 标签管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":article,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /article/:article_id [get]
func GetBlogArticle(c *gin.Context)  {
  var blogArticleModel models.BlogArticle

  articleId, err := strconv.ParseInt(c.Param("article_id"),10,64)
  if err != nil {
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  blogArticleModel.ArticleID = articleId

  article, err := blogArticleModel.GetBlogArticle()
  if err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"获取失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"获取成功",Data:article})
}

// @Summary 添加文章
// @Tags 文章管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param article body blogArticleForm true "账号信息"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"添加成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"添加失败"}
// @Router /article [post]
//]
func CreateBlogArticle(c *gin.Context)  {
  var (
    blogArticleForm blogArticleForm
    blogArticleModel models.BlogArticle
  )
  if err := c.BindJSON(&blogArticleForm); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  blogArticleModel.ArticleTitle = blogArticleForm.ArticleTitle
  blogArticleModel.ArticleStatus = blogArticleForm.ArticleStatus
  blogArticleModel.ArticleDesc = blogArticleForm.ArticleDesc
  blogArticleModel.ArticleContent = blogArticleForm.ArticleContent

  for _, tagId := range blogArticleForm.TagIds {
    tag := models.BlogTag{
      TagID:     tagId,
    }
    blogArticleModel.Tags = append(blogArticleModel.Tags,tag)
  }

  if err := blogArticleModel.Create(); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"添加失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"添加成功"})
}

// @Summary 修改文章
// @Tags 文章管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param article body blogArticleForm true "账号信息"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"修改成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"修改失败"}
// @Router /article/:article_id [put]
//]
func UpdateBlogArticle(c *gin.Context)  {
  var (
    blogArticleForm blogArticleForm
    blogArticleModel models.BlogArticle
  )
  if err := c.BindJSON(&blogArticleForm); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  articleId, err := strconv.ParseInt(c.Param("article_id"),10,64)
  if  err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  if article, err := blogArticleModel.GetBlogArticleById(articleId); article.ArticleID == 0 || err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"文章不存在"})
    return
  }

  blogArticleModel.ArticleID = articleId
  blogArticleModel.ArticleTitle = blogArticleForm.ArticleTitle
  blogArticleModel.ArticleStatus = blogArticleForm.ArticleStatus
  blogArticleModel.ArticleDesc = blogArticleForm.ArticleDesc
  blogArticleModel.ArticleContent = blogArticleForm.ArticleContent

  for _, tagId := range blogArticleForm.TagIds {
    tag := models.BlogTag{
      TagID:     tagId,
    }
    blogArticleModel.Tags = append(blogArticleModel.Tags,tag)
  }

  if err := blogArticleModel.Update(); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"修改失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"修改成功"})
}

// @Summary 删除用户
// @Tags 用户管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param article_id body []int64 true "文章ID数组"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"删除成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"删除失败"}
// @Router /articles [delete]
//]
func DeleteBlogArticle(c *gin.Context)  {
  var articleIds []int64
  if err := c.BindJSON(&articleIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  var (
    blogArticleModel models.BlogArticle
  )
  if err := blogArticleModel.Delete(articleIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"删除成功"})
}

// @Summary 文章列表
// @Tags 文章管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":articles,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /tags [get]
func BlogArticleList(c *gin.Context)  {
  var blogArticleModel models.BlogArticle

  pageIndex, _ := strconv.Atoi(c.DefaultQuery("page","1"))
  pageSize, _ := strconv.Atoi(c.DefaultQuery("limit","20"))
  articleTitle := c.Query("article_title")
  articleStatus, _ := strconv.Atoi(c.DefaultQuery("article_status","-1"))
  tagId, _ := strconv.ParseInt(c.DefaultQuery("tag_id","-1"),10,64)

  articles, total, err := blogArticleModel.GetBlogArticlePage(pageSize,pageIndex,articleTitle,articleStatus,tagId)
  if err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"获取数据失败"})
    return
  }
  type blogArticles struct {
    BlogArticles  []models.BlogArticle `json:"articles"`
    Total          int64 `json:"total"`
  }
  c.JSON(200,Res{Code:0,Message:"获取成功",Data:&blogArticles{BlogArticles:articles,Total:total}})
}

