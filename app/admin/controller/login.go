package controller

import (
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/controller"
	"go-mage-admin/app/core/route"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Login{})
}

// Login 后台登录页面
type Login struct {
	// 继承控制器基类
	controller.Abstract
	Admin
}

func (login *Login) Index(c *core.Context) {
	// 输出页面 调用对象来处理
	admin := &Admin{}
	login.HTMLRender = admin.LoadLayout()

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
	})

}
func (login *Login) Save(c *core.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
	})
}
