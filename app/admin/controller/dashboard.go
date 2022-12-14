package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/core/route"
	mageApp "go-mage-admin/app/mage"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Dashboard{}, _admin.Name)
}

// Dashboard 后台首页
type Dashboard struct {
	// 继承控制器基类
}

func (d *Dashboard) IndexGET(c *gin.Context) {
	// 输出页面 调用对象来处理
	admin := &Admin{}
	mageApp.AppGin.HTMLRender = admin.LoadLayout()
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"code": 1,
		"msg":  "ok",
	})

}
