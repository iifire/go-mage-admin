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
	route.Register(&Dep{}, _admin.Name)
}

// Dep 后台首页
type Dep struct {
	// 继承控制器基类
}

func (d *Dep) IndexGET(c *gin.Context) {
	// 输出页面 调用对象来处理
	admin := &Admin{}
	extraTpl := make([]string, 0)
	extraTpl = append(extraTpl, "dep/list.html")
	admin.SetExtraTpl(extraTpl)
	mageApp.AppGin.HTMLRender = admin.LoadLayout()
	c.HTML(http.StatusOK, "dep.html", gin.H{
		"code":   1,
		"msg":    "ok",
		"menu":   "system",
		"urlCur": "/admin/dep/index",
	})

}
