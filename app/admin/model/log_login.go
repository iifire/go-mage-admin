package model

import (
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/model"
)

type LogLogin struct {
	model.Collection
	LogId      int    `gorm:"primary_key" json:"log_id"`
	LogUser    string `json:"user_id"`
	Ip         string `json:"ip"`
	Status     int    `json:"status"`
	DateCreate string `json:"date_create"`
}

// TableName 表名
func (*LogLogin) TableName() string {
	return "admin_log_login"
}

func (g *LogLogin) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]LogLogin, int64) {
	var logs []LogLogin
	var total int64
	db := core.AppDb["read"].Model(LogLogin{})

	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&logs)
	if res.Error != nil {
		panic("获取LogLogin失败")
	}
	return logs, total
}

// DelByIds 批量删除
func (g *LogLogin) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["log_id"] = ids
	res := core.AppDb["write"].Where(where).Delete(&LogLogin{})
	if res.Error != nil {
		return false
	}
	return true
}
