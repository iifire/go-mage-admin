package model

import (
	"go-mage-admin/app/mage"
	"gorm.io/gorm"
	"strings"
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
func (cfg *ConfigGrid) Reset(uid int, code string) bool {
	return cfg.Update(uid, code, "")
}
func (*ConfigGrid) Update(uid int, code string, configData string) bool {
	var cfg ConfigGrid
	result := mage.AppDb["read"].Debug().Model(&cfg).Where("link_user = ? and link_code = ?", uid, code).Update("config", configData)
	if result.Error != nil {
		return false
	} else {
		return true
	}
}

// getConfig 获取一级菜单
func (*ConfigGrid) getConfig(uid int, code string) []string {
	var cfg ConfigGrid
	result := mage.AppDb["read"].First(&cfg, "link_user = ? and link_code = ?", uid, code)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		panic("获取Grid Column配置失败")
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	if cfg.Config != "" {
		return strings.Split(cfg.Config, ",")
	} else {
		return nil
	}

}
