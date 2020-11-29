package models

import (
  "github.com/jinzhu/gorm"
  orm "go-element-admin-api/db"
  "go-element-admin-api/utils"
  "log"
)

type Path struct {
  PathID       int64   `gorm:"column:path_id;primary_key;AUTO_INCREMENT" json:"path_id"`
  ParentID     int64   `gorm:"column:parent_id;" json:"parent_id" binding:""`
  Type         string  `gorm:"column:type;" json:"type" binding:"required"`
  Name         string  `gorm:"column:name;" json:"name" binding:"required"`
  Path         string  `gorm:"column:path;" json:"path" binding:""`
  Method       string  `gorm:"column:method;" json:"method" binding:""`
  UpdateTime   string  `gorm:"column:update_time;default:NULL" json:"update_time"`
  CreateTime   string  `gorm:"column:create_time" json:"create_time"`
}

type TreePath struct {
  Path
  Children []TreePath `json:"children"`
}

// 获取接口
func (p Path) GetPath() (path Path, err error)  {
  if err = orm.Eloquent.Table("paths").Where(&p).First(&path).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  return
}

// 创建接口
func (p Path) Create() (pathId int64, err error)  {
  p.CreateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("paths").Create(&p).Error; err != nil {
    log.Println(err.Error())
  }
  pathId = p.PathID
  return
}

// 修改接口
func (p Path) Update() (err error) {
  p.UpdateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("paths").Omit("create_time").Save(&p).Error; err != nil {
    log.Println(err.Error())
  }
  return
}


// 删除接口
func (p Path) Delete(pathIds []int64)(err error)  {
  if err = orm.Eloquent.Table("paths").Where("path_id in (?)",pathIds).Delete(&p).Error; err != nil {
    log.Println(err.Error())
  }
  return
}


// 根据接口ID切片获取接口
func (p Path) GetPathByIDs(pathIds []int64) (paths []Path, err error)  {
  if err = orm.Eloquent.Table("paths").Where("path_id in (?)",pathIds).Find(&paths).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  return
}

func (p Path) GetTreePaths() (paths []TreePath, err error) {
  if err = orm.Eloquent.Table("paths").Where(&p).Order("parent_id asc").Find(&paths).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  if len(paths) > 0 {
    paths = makeTreePath(paths,paths[0].ParentID)
  }
  return
}

func (p Path) GetTreePathByIds(pathIds []int64) (paths []TreePath, err error) {
  if err = orm.Eloquent.Table("paths").Where("path_id in (?)",pathIds).Find(&paths).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  paths = makeTreePath(paths,0)
  return
}

// 生成树形结构数据
func makeTreePath(treePaths []TreePath, parentId int64) []TreePath {
  var treePath []TreePath
  for _, path := range treePaths {
    if path.ParentID == parentId {
      path.Children = makeTreePath(treePaths,path.PathID)
      treePath = append(treePath, path)
    }
  }

  return treePath
}
