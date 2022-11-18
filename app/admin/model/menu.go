package model

import (
	"go-mage-admin/app/core"
	"log"
)

type Menu struct {
	MenuId     int    `gorm:"primary_key" json:"menu_id"`
	Parent     int    `json:"parent"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Level      int    `json:"level"`
	Icon       string `json:"icon"`
	Action     string `json:"action"`
	Position   int    `json:"position"`
	DateCreate int    `json:"date_create"`
	DateUpdate int    `json:"date_update"`
}

// TableName 表名
func (*Menu) TableName() string {
	return "admin_menu"
}

// GetTopMenus 获取一级菜单
func (menu *Menu) GetTopMenus() []Menu {
	var menus []Menu
	result := core.AppDb["read"].Where("parent = ?", "0").Find(&menus)
	if result.Error != nil {
		panic("获取菜单失败")
	}
	return menus
}

// GetSubMenus 获取一级菜单的子菜单
func (menu *Menu) GetSubMenus(pid string) []Menu {
	var menus []Menu
	log.Println("pid=", pid)
	result := core.AppDb["read"].Where("parent = ?", pid).Find(&menus)
	if result.Error != nil {
		panic("获取菜单失败")
	}
	return menus
}
