package admin

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/controller"
	"go-mage-admin/app/core/route"
	mageApp "go-mage-admin/app/mage"
	_modBlock "go-mage-admin/app/shop/block/admin"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&ShopOrder{}, _admin.Name)
}

// ShopOrder 自定义菜单
type ShopOrder struct {
}

// IndexGET 订单列表页
func (m *ShopOrder) IndexGET(c *gin.Context) {
	admin := &controller.Admin{}
	tplName := "IndexGET"
	admin.UseGridTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"menu":    "shop",
		"urlCur":  "/admin/shopOrder/index",
		"urlGrid": "/admin/shopOrder/grid",
		"msg":     "ok",
	})
}

// GridREQ Grid接口
func (m *ShopOrder) GridREQ(c *gin.Context) {
	grid := _modBlock.OrderGrid{}
	grid.GetCollection()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": grid,
	})
}
