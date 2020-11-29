package apis

import (
  "github.com/gin-gonic/gin"
  "go-element-admin-api/models"
  "log"
  "strconv"
)

// @Summary 获取标签
// @Tags 标签管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":tag,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /tag/:tag_id [get]
func GetBlogTag(c *gin.Context)  {
  var blogTagModel models.BlogTag

  tagId, err := strconv.ParseInt(c.Param("tag_id"),10,64)
  if err != nil {
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  blogTagModel.TagID = tagId

  tag, err := blogTagModel.GetBlogTag()
  if err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"获取失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"获取成功",Data:tag})
}


// @Summary 添加标签
// @Tags 标签管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param path body blogTagForm true "标签数据"
// @Success 200 {object} Res {"code":0,"msg":"添加成功"}
// @Failure 400 {object} Res {"code":400,"msg":"msg"}
// @Router /tag [post]
func CreateBlogTag (c *gin.Context){
  var blogTagModel models.BlogTag

  if err := c.BindJSON(&blogTagModel); err != nil {
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  var _blogTagModel models.BlogTag
  _blogTagModel.TagTitle = blogTagModel.TagTitle
  if tag, _ := _blogTagModel.GetBlogTag(); tag.TagID > 0 {
    c.JSON(400,Res{Code:400,Message:"该标签已存在"})
    return
  }

  if _,err := blogTagModel.Create(); err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"添加失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"添加成功"})
}

// @Summary 修改标签
// @Tags 标签管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param path body blogTagForm true " 标签数据"
// @Success 200 {object} Res {"code":0,"msg":"修改成功"}
// @Failure 400 {object} Res {"code":400,"msg":"msg"}
// @Router /tag/:tag_id [put]
func UpdateBlogTag (c *gin.Context){
  var blogTagModel models.BlogTag

  if err := c.BindJSON(&blogTagModel); err != nil {
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }


  tagId, err := strconv.ParseInt(c.Param("tag_id"),10,64)
  if err != nil {
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  var _blogTagModel models.BlogTag
  _blogTagModel.TagTitle = blogTagModel.TagTitle
  if tag, _ := _blogTagModel.GetBlogTag(); tag.TagID > 0 && tag.TagID != tagId {
    c.JSON(400,Res{Code:400,Message:"该标签已存在"})
    return
  }

  blogTagModel.TagID = tagId

  if err := blogTagModel.Update(); err != nil {
    c.JSON(400,Res{Code:400,Message:"修改失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"修改成功"})
}

// @Summary 删除标签
// @Tags 标签管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param tag_id body []int64 true "标签ID数组"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"删除成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"删除失败"}
// @Router /tags [delete]
//]
func DeleteBlogTag(c *gin.Context)  {
  var blogTagIds []int64
  if err := c.BindJSON(&blogTagIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  var (
    blogTagModel models.BlogTag
  )
  if err := blogTagModel.Delete(blogTagIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"删除成功"})
}

// @Summary 标签列表
// @Tags 标签管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":tags,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /tags [get]
func BlogTagList(c *gin.Context)  {
  var blogTagModel models.BlogTag

  pageIndex, _ := strconv.Atoi(c.DefaultQuery("page","1"))
  pageSize, _ := strconv.Atoi(c.DefaultQuery("limit","20"))
  tagTitle := c.Query("tag_title")
  tagStatus, _ := strconv.Atoi(c.DefaultQuery("tag_status","1"))

  tags, total, err := blogTagModel.GetBlogTagPage(pageSize,pageIndex,tagTitle,tagStatus)
  if err != nil {
    log.Println(err.Error())
    c.JSON(400,Res{Code:400,Message:"获取数据失败"})
    return
  }
  type blogTags struct {
    BlogTags  []models.BlogTag `json:"tags"`
    Total int64               `json:"total"`
  }
  c.JSON(200,Res{Code:0,Message:"获取成功",Data:&blogTags{BlogTags:tags,Total:total}})
}
