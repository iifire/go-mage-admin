package model

import (
	"go-mage-admin/app/core/model"
	"go-mage-admin/app/mage"
	"log"
)

type Rule struct {
	RuleId     int    `gorm:"primary_key" json:"rule_id"`
	Cid        int    `json:"cid"`
	Name       string `json:"name"`
	Position   int    `json:"position"`
	IsActive   int    `json:"is_active"`
	Type       string `json:"type"`
	DateCreate string `json:"date_create"`
}

// TableName 表名
func (*Rule) TableName() string {
	return "wxmp_rule"
}

func (r *Rule) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]Rule, int64) {
	var rs []Rule
	var total int64
	db := mage.AppDb["read"].Model(Rule{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&rs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取Rule失败")
	}
	return rs, total
}

// DelByIds 批量删除
func (r *Rule) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["rule_id"] = ids
	res := mage.AppDb["write"].Where(where).Delete(&Rule{})
	if res.Error != nil {
		return false
	}
	return true
}
