package model

import (
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/model"
	"log"
)

type Attach struct {
	AttachId   int    `gorm:"primary_key" json:"attach_id"`
	Cid        int    `json:"cid"`
	Filename   string `json:"filename"`
	Attachment string `json:"attachment"`
	MediaId    string `json:"media_id"`
	Type       string `json:"type"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Model      string `json:"model"`
	DateCreate string `json:"date_create"`
}

// TableName 表名
func (*Attach) TableName() string {
	return "wxmp_attach"
}

func (r *Attach) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]Attach, int64) {
	var rs []Attach
	var total int64
	db := core.AppDb["read"].Model(Attach{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&rs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取Attach失败")
	}
	return rs, total
}

// DelByIds 批量删除
func (r *Attach) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["attach_id"] = ids
	res := core.AppDb["write"].Where(where).Delete(&Attach{})
	if res.Error != nil {
		return false
	}
	return true
}
