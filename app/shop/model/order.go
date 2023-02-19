package model

import (
	"go-mage-admin/app/core/model"
	"go-mage-admin/app/mage"
	"log"
)

type Order struct {
	OrderId     int    `gorm:"primary_key" json:"order_id"`
	IncrementId string `json:"increment_id"`
	GrandTotal  string `json:"grand_total"`
	Total       string `json:"total"`
	DateCreate  string `json:"date_create"`
	Ip          string `json:"ip"`
}

// TableName 表名
func (*Order) TableName() string {
	return "shop_order"
}

func (r *Order) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]Order, int64) {
	var rs []Order
	var total int64
	db := mage.AppDb["read"].Model(Order{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&rs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取Order失败")
	}
	return rs, total
}

// DelByIds 批量删除
func (r *Order) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["order_id"] = ids
	res := mage.AppDb["write"].Where(where).Delete(&Order{})
	if res.Error != nil {
		return false
	}
	return true
}
