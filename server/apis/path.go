package apis

import (
  "github.com/gin-gonic/gin"
  "go-element-admin/models"
  "strconv"
)

// @Summary 获取接口
// @Tags 接口管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":path,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /path/:path_id [get]
func GetPath(c *gin.Context)  {
  var pathModel models.Path

  pathId, err := strconv.ParseInt(c.Param("path_id"),10,64)
  if err != nil {
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  pathModel.PathID = pathId

  path, err := pathModel.GetPath()
  if err != nil {
    lgr.Errorf("获取接口失败: %v",err.Error())
    c.JSON(400,Res{Code:400,Message:"获取失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"获取成功",Data:path})
}


// @Summary 添加接口
// @Tags 接口管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param path body pathForm true "接口数据"
// @Success 200 {object} Res {"code":0,"msg":"添加成功"}
// @Failure 400 {object} Res {"code":400,"msg":"msg"}
// @Router /path [post]
func CreatePath (c *gin.Context){
  var pathModel models.Path

  if err := c.BindJSON(&pathModel); err != nil {
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  if pathModel.Type == "J" {
    if pathModel.Path == "" {
      c.JSON(400,Res{Code:400,Message:"接口路径不能为空"})
      return
    }
    if pathModel.Method == "" {
      c.JSON(400,Res{Code:400,Message:"接口方法不能为空"})
      return
    }
  }

  if _,err := pathModel.Create(); err != nil {
    lgr.Errorf("添加接口失败: %v",err.Error())
    c.JSON(400,Res{Code:400,Message:"添加失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"添加成功"})
}

// @Summary 修改接口
// @Tags 接口管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param path body pathForm true " 接口数据"
// @Success 200 {object} Res {"code":0,"msg":"修改成功"}
// @Failure 400 {object} Res {"code":400,"msg":"msg"}
// @Router /path/:path_id [put]
func UpdatePath (c *gin.Context){
  var pathModel models.Path

  if err := c.BindJSON(&pathModel); err != nil {
    c.JSON(400,Res{Code:400,Message:"数据验证失败," + err.Error()})
    return
  }

  if pathModel.Type == "J" {
    if pathModel.Path == "" {
      c.JSON(400,Res{Code:400,Message:"接口路径不能为空"})
      return
    }
    if pathModel.Method == "" {
      c.JSON(400,Res{Code:400,Message:"接口方法不能为空"})
      return
    }
  }

  pathId, err := strconv.ParseInt(c.Param("path_id"),10,64)
  if err != nil {
    c.JSON(400,Res{Code:400,Message:"参数验证失败"})
    return
  }

  pathModel.PathID = pathId

  if err := pathModel.Update(); err != nil {
    lgr.Errorf("修改接口失败: %v",err.Error())
    c.JSON(400,Res{Code:400,Message:"修改失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"修改成功"})
}

// @Summary 删除接口
// @Tags 接口管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param path_id body []int64 true "接口ID数组"
// @Success 200 {object} Res {"code":0,"data":null,"msg":"删除成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"删除失败"}
// @Router /paths [delete]
//]
func DeletePath(c *gin.Context)  {
  var pathIds []int64
  if err := c.BindJSON(&pathIds); err != nil {
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  var (
    pathModel models.Path
  )
  if err := pathModel.Delete(pathIds); err != nil {
    lgr.Errorf("删除接口失败: %v",err.Error())
    c.JSON(400,Res{Code:400,Message:"删除失败"})
    return
  }

  c.JSON(200,Res{Code:0,Message:"删除成功"})
}

// @Summary 接口列表
// @Tags 接口管理
// @accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} Res {"code":0,"data":paths,"msg":"获取成功"}
// @Failure 400 {object} Res {"code":400,"data":null,"msg":"msg"}
// @Router /paths [get]
func PathList(c *gin.Context)  {
  var pathModel models.Path
  pathModel.Name = c.Query("name")
  pathModel.Path = c.Query("path")

  paths, err := pathModel.GetTreePaths()
  if err != nil {
    lgr.Errorf("获取接口列表失败: %v",err.Error())
    c.JSON(400,Res{Code:400,Message:"获取数据失败"})
    return
  }
  c.JSON(200,Res{Code:0,Message:"获取成功",Data:paths})
}
