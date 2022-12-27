package model

import (
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/model"
	"log"
)

type News struct {
	NewsId             int    `gorm:"primary_key" json:"news_id"`
	Cid                int    `json:"cid"`
	Type               int    `json:"type"`
	Attach             int    `json:"attach"`
	Title              string `json:"title"`
	Author             string `json:"author"`
	Digest             string `json:"digest"`
	Content            string `json:"content"`
	ThumbUrl           string `json:"thumb_url"`
	ThumbMediaId       string `json:"thumb_media_id"`
	ShowCoverPic       int    `json:"show_cover_pic"`
	ContentSourceUrl   string `json:"content_source_url"`
	Url                string `json:"url"`
	Position           int    `json:"position"`
	NeedOpenComment    int    `json:"need_open_comment"`
	OnlyFansCanComment int    `json:"only_fans_can_comment"`
	DateCreate         string `json:"date_create"`
}

// TableName 表名
func (*News) TableName() string {
	return "wxmp_news"
}

func (r *News) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]News, int64) {
	var rs []News
	var total int64
	db := core.AppDb["read"].Model(News{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&rs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取News失败")
	}
	return rs, total
}

// DelByIds 批量删除
func (r *News) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["news_id"] = ids
	res := core.AppDb["write"].Where(where).Delete(&News{})
	if res.Error != nil {
		return false
	}
	return true
}
