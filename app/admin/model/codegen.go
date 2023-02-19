package model

import (
	"go-mage-admin/app/core/model"
	"go-mage-admin/app/mage"
	"log"
)

type Codegen struct {
	model.Collection
	CodegenId  int    `gorm:"primary_key" json:"codegen_id"`
	Type       string `json:"type"`
	Creator    string `json:"creator"`
	DateCreate int    `json:"date_create"`
}

func (*Codegen) TableName() string {
	return "admin_codegen"
}

func (codegen *Codegen) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]Codegen, int64) {
	var codegens []Codegen
	var total int64
	db := mage.AppDb["read"].Model(Codegen{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Debug().Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&codegens)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取Codegen失败")
	}
	return codegens, total
}

// GetGridColumnsConfig 获取用户自定义的Grid
func (codegen *Codegen) GetGridColumnsConfig(uid int, gridCode string) []string {
	configGrid := new(ConfigGrid).getConfig(uid, gridCode)
	return configGrid
}

// DelByIds 批量删除
func (codegen *Codegen) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["codegen_id"] = ids
	res := mage.AppDb["write"].Where(where).Delete(&Codegen{})
	if res.Error != nil {
		return false
	}
	return true
}

func (codegen *Codegen) LoadById(id int) *Codegen {
	u := new(Codegen)
	if id > 0 {
		mage.AppDb["read"].First(&u, "codegen_id = ?", id)
	}
	return u
}
