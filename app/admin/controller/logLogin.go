package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/route"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&logLogin{}, _admin.Name)
}

// logLogin 后台首页
type logLogin struct {
	// 继承控制器基类
}

func (u *logLogin) IndexGET(c *gin.Context) {
	// 输出页面 调用对象来处理
	admin := &Admin{}
	//admin.SetLayout("2columns")
	core.AppGin.HTMLRender = admin.LoadLayout()
	c.HTML(http.StatusOK, "logLogin.html", gin.H{
		"code": 1,
		"msg":  "ok",
	})

}
