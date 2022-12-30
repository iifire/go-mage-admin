package admin

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/controller"
	"go-mage-admin/app/core/route"
	mageApp "go-mage-admin/app/mage"
	_modBlock "go-mage-admin/app/wxmp/block/admin"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&WxmpFans{}, _admin.Name)
}

// WxmpFans 自定义菜单
type WxmpFans struct {
}

// IndexGET 菜单列表页
func (m *WxmpFans) IndexGET(c *gin.Context) {
	admin := &controller.Admin{}
	tplName := "IndexGET"
	admin.UseGridTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"menu":    "wxmp",
		"urlCur":  "/admin/wxmpFans/index",
		"urlGrid": "/admin/wxmpFans/grid",
		"msg":     "ok",
	})
}

// GridREQ Grid接口
func (m *WxmpFans) GridREQ(c *gin.Context) {
	grid := _modBlock.FansGrid{}
	grid.GetCollection()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": grid,
	})
}
