package model

import (
	"encoding/json"
	"go-mage-admin/app/core"
	"gorm.io/gorm"
)

type ConfigGrid struct {
	Id         int    `gorm:"primary_key" json:"id"`
	LinkUser   int    `json:"user"`
	LinkCode   string `json:"code"`
	Config     string `json:"config"`
	DateCreate int    `json:"date_create"`
}

// TableName 表名
func (*ConfigGrid) TableName() string {
	return "admin_config_grid"
}

// getConfig 获取一级菜单
func (*ConfigGrid) getConfig(uid int, code string) []string {
	var cfg ConfigGrid
	result := core.AppDb["read"].First(&cfg, "link_user = ? and link_code = ?", uid, code)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic("获取Grid Column配置失败")
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	if cfg.Config != "" {
		var s []string
		json.Unmarshal([]byte(cfg.Config), &s)
		return s
	} else {
		return nil
	}

}
