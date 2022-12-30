package model

import (
	"go-mage-admin/app/core/model"
	"go-mage-admin/app/mage"
	"log"
)

type Fans struct {
	RoleId   int    `gorm:"primary_key" json:"fans_id"`
	Nickname string `json:"nickname"`
	Openid   string `json:"openid"`
}

// TableName 表名
func (*Fans) TableName() string {
	return "wxmp_fans"
}

func (r *Fans) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]Fans, int64) {
	var rs []Fans
	var total int64
	db := mage.AppDb["read"].Model(Fans{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&rs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取Fans失败")
	}
	return rs, total
}

// DelByIds 批量删除
func (r *Fans) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["fans_id"] = ids
	res := mage.AppDb["write"].Where(where).Delete(&Fans{})
	if res.Error != nil {
		return false
	}
	return true
}
