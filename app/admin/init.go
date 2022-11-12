package admin

import (
	"go-mage-admin/app"
	"go-mage-admin/app/core/init"
)

// Init 模块初始化对象
type Init struct {
	init.Abstract
}

// Init 核心模块初始化
func (i *Init) Init(mage *app.Mage) {
	key := "_helper/" + "admin"
	mage.Registry[key] = &Helper{}
	return
}
