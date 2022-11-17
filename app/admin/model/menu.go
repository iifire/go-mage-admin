package model

import (
	"go-mage-admin/app/core"
)

type Menu struct {
	MenuId     int    `gorm:"primary_key" json:"menu_id"`
	Parent     int    `json:"parent"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Level      int    `json:"level"`
	Icon       string `json:"icon"`
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
	result := core.AppDb["read"].Find(&menus).Where("level = ?", "0")
	if result.Error != nil {
		panic("获取菜单失败")
	}
	return menus
}
