package models

import (
  "github.com/jinzhu/gorm"
  orm "go-element-admin-api/db"
  "go-element-admin-api/utils"
  "log"
)

type Menu struct {
	MenuID      int64  `gorm:"column:menu_id;primary_key;AUTO_INCREMENT" json:"menu_id"`
	ParentID    int64  `gorm:"column:parent_id" json:"parent_id" binding:""`
	Title       string `gorm:"column:title" json:"title" binding:"required"`
	Sort        int64  `gorm:"column:sort" json:"sort" binding:""`
	Type        string `gorm:"column:type" json:"type" binding:"required"`
	Icon        string `gorm:"column:icon" json:"icon" binding:"required"`
	Name        string `gorm:"column:name" json:"name" binding:"required"`
	Component   string `gorm:"column:component" json:"component" binding:"required"`
	Path        string `gorm:"column:path" json:"path" binding:"required"`
	Redirect    string `gorm:"column:redirect" json:"redirect" binding:""`
	Permission  string `gorm:"column:permission" json:"permission" binding:""`
	Hidden      int    `gorm:"column:hidden" json:"hidden" binding:""`
	UpdateTime  string `gorm:"column:update_time;default:NULL" json:"update_time"`
	CreateTime  string `gorm:"column:create_time" json:"create_time"`
}

type TreeMenu struct {
	Menu
	Children []TreeMenu `json:"children"`
}

// 获取菜单
func (m Menu) GetMenu() (menu Menu, err error)  {
	if err = orm.Eloquent.Table("menus").Where(&m).First(&menu).Error; err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

// 创建菜单
func (m Menu) Create() (menuId int64, err error)  {
  m.CreateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("menus").Create(&m).Error; err != nil {
    log.Println(err.Error())
  }
  menuId = m.MenuID
  return
}

// 修改菜单
func (m Menu) Update() (err error) {
  m.UpdateTime = utils.GetCurrntTime()
  if err = orm.Eloquent.Table("menus").Omit("create_time").Save(&m).Error; err != nil {
    log.Println(err.Error())
  }
  return
}


// 删除菜单
func (m Menu) Delete(menuIds []int64)(err error)  {
  if err = orm.Eloquent.Table("menus").Where("menu_id in (?)",menuIds).Delete(&m).Error; err != nil {
    log.Println(err.Error())
  }
  return
}


// 根据菜单ID切片获取菜单
func (m Menu) GetMenuByIDs(menuIds []int64) (menus []Menu, err error)  {
	if err = orm.Eloquent.Table("menus").Where("menu_id in (?)",menuIds).Find(&menus).Error; err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
	return
}

func (m Menu) GetTreeMenus() (menus []TreeMenu, err error) {
  if err = orm.Eloquent.Table("menus").Where(&m).Order("parent_id asc").Order("sort desc").Find(&menus).Error; err != nil {
    log.Println(err.Error())
    if err == gorm.ErrRecordNotFound  {
      err = nil
    }
  }
  if len(menus) > 0 {
    menus = makeTreeMenu(menus,menus[0].ParentID)
  }
  return
}

func (m Menu) GetTreeMenuByIds(menuIds []int64) (menus []TreeMenu, err error) {
	if err = orm.Eloquent.Table("menus").Where("menu_id in (?)",menuIds).Order("sort desc").Find(&menus).Error; err != nil {
		log.Println(err.Error())
		if err == gorm.ErrRecordNotFound  {
			err = nil
		}
	}
  menus = makeTreeMenu(menus,0)
	return
}

// 生成树形结构数据
func makeTreeMenu(treeMenus []TreeMenu, parentId int64) []TreeMenu {
	var treeMenu []TreeMenu
	for _, menu := range treeMenus {
		if menu.ParentID == parentId {
      menu.Children = makeTreeMenu(treeMenus,menu.MenuID)
			treeMenu = append(treeMenu, menu)
		}
	}

	return treeMenu
}
