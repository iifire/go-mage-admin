package app

// 由cmd(调Cobra库)动态生成 切勿修改(再次生成后会被覆盖) 导入当前生效的模块并调用模块下的helper|init等go文件的init方法
import (
	_ "go-mage-admin/app/admin"
	_ "go-mage-admin/app/admin/controller"
	_ "go-mage-admin/app/core"
	_ "go-mage-admin/app/customer"
	_ "go-mage-admin/app/customer/controller"
)
