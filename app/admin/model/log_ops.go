package model

import (
	"go-mage-admin/app/core/model"
	"go-mage-admin/app/mage"
	"log"
)

type LogOps struct {
	//匿名继承
	model.Collection
	LogOpsId   int    `gorm:"primary_key" json:"log_id"`
	LogUser    int    `json:"user_id"`
	Url        string `json:"url"`
	Ip         string `json:"ip"`
	Time       int    `json:"time"`
	Status     int    `json:"status"`
	DateCreate string `json:"date_create"`
}

// TableName 表明
func (*LogOps) TableName() string {
	return "admin_log_ops"
}

func (g *LogOps) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]LogOps, int64) {
	var logs []LogOps
	var total int64
	db := mage.AppDb["read"].Model(LogOps{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	log.Println("filters:", filters)
	db2 := db
	db.Debug().Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&logs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取LogOps失败")
	}
	return logs, total
}

// DelByIds 批量删除
func (g *LogOps) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["log_id"] = ids
	res := mage.AppDb["write"].Debug().Where(where).Delete(&LogOps{})
	if res.Error != nil {
		return false
	}
	return true
}
