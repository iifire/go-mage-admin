package core

import (
	r "go-mage-admin/app/core/init"
)

// init 模块初始化方法 bootstrap里调用
func init() {

}

// Init 模块初始化对象
type Init struct {
	r.Abstract
}

// Init 核心模块初始化
func (i *Init) Init() {
	return
}
