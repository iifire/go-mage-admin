package admin

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/controller"
	"go-mage-admin/app/core/route"
	mageApp "go-mage-admin/app/mage"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&WxmpMenu{}, _admin.Name)
}

// WxmpMenu 自定义菜单
type WxmpMenu struct {
}

// IndexGET 菜单列表页
func (m *WxmpMenu) IndexGET(c *gin.Context) {
	admin := &controller.Admin{}
	mageApp.AppGin.HTMLRender = admin.LoadLayout()
	c.HTML(http.StatusOK, "wxmp_menu.html", gin.H{
		"code":   1,
		"msg":    "ok",
		"menu":   "system",
		"urlCur": "admin/wxmpMenu/index",
	})
}
