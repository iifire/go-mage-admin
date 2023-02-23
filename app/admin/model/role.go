package model

import (
	"go-mage-admin/app/core/model"
	"go-mage-admin/app/mage"
	"log"
)

type Role struct {
	RoleId     int    `gorm:"primary_key" json:"role_id"`
	RoleName   string `json:"role_name"`
	AllowAll   int    `json:"allow_all"`
	Memo       string `json:"memo"`
	DateUpdate string `json:"date_update"`
	DateCreate string `json:"date_create"`
}

// TableName 表名
func (*Role) TableName() string {
	return "admin_role"
}

func (r *Role) GetCollection(filters map[string]interface{}, orders [2]string, page int, size int) ([]Role, int64) {
	var rs []Role
	var total int64
	db := mage.AppDb["read"].Model(Role{})

	//多条件过滤
	collection := model.Collection{}
	db = collection.PrepareCollection(db, filters)
	db2 := db
	db.Count(&total)
	res := collection.LoadCollection(db2, orders, page, size).Find(&rs)
	if res.Error != nil {
		log.Println(res.Error)
		panic("获取Role失败")
	}
	return rs, total
}

// DelByIds 批量删除
func (r *Role) DelByIds(ids []string) bool {
	where := map[string]interface{}{}
	where["role_id"] = ids
	res := mage.AppDb["write"].Where(where).Delete(&Role{})
	if res.Error != nil {
		return false
	}
	return true
}

func (r *Role) LoadById(id int) *Role {
	rule := new(Role)
	if id > 0 {
		mage.AppDb["read"].First(&rule, "rule_id = ?", id)
	}
	return rule
}
