package model

import (
	"go-mage-admin/app/core"
	"log"
)

type CoreSmsLog struct {
	Collection
	LogId      int    `gorm:"primary_key" json:"log_id"`
	Mobile     string `json:"mobile"`
	Type       int    `json:"type"`
	Code       string `json:"code"`
	Ip         string `json:"ip"`
	Extra      string `json:"extra"`
	Status     int    `json:"status"`
	DateCreate string `json:"date_create"`
}

func (*CoreSmsLog) TableName() string {
	return "core_sms_log"
}

func (g *CoreSmsLog) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]CoreSmsLog, int64) {
	var logs []CoreSmsLog
	var total int64
	db := core.AppDb["read"].Model(CoreSmsLog{})

	//多条件过滤
	collection := Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&logs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取CoreSmsLog失败")
	}
	return logs, total
}

// DelByIds 批量删除
func (g *CoreSmsLog) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["log_id"] = ids
	res := core.AppDb["write"].Where(where).Delete(&CoreSmsLog{})
	if res.Error != nil {
		return false
	}
	return true
}
