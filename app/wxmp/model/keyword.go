package model

import (
	"go-mage-admin/app/core/model"
	"go-mage-admin/app/mage"
	"log"
)

type Keyword struct {
	KeywordId  int    `gorm:"primary_key" json:"keyword_id"`
	Cid        int    `json:"cid"`
	Rule       int    `json:"rule"`
	Content    string `json:"content"`
	Type       int    `json:"type"`
	Position   int    `json:"position"`
	DateCreate string `json:"date_create"`
}

// TableName 表名
func (*Keyword) TableName() string {
	return "wxmp_keyword"
}

func (r *Keyword) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]Keyword, int64) {
	var rs []Keyword
	var total int64
	db := mage.AppDb["read"].Model(Keyword{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&rs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取Keyword失败")
	}
	return rs, total
}

// DelByIds 批量删除
func (r *Keyword) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["keyword_id"] = ids
	res := mage.AppDb["write"].Where(where).Delete(&Keyword{})
	if res.Error != nil {
		return false
	}
	return true
}
