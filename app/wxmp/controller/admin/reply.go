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
	route.Register(&WxmpReply{}, _admin.Name)
}

// WxmpReply 自定义菜单
type WxmpReply struct {
}

// IndexGET 菜单列表页
func (m *WxmpReply) IndexGET(c *gin.Context) {
	admin := &controller.Admin{}
	tplName := "IndexGET"
	admin.UseGridTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"menu":    "wxmp",
		"urlCur":  "/admin/wxmpReply/index",
		"urlGrid": "/admin/wxmpReply/grid",
		"msg":     "ok",
	})
}

// GridREQ Grid接口
func (m *WxmpReply) GridREQ(c *gin.Context) {
	grid := _modBlock.ReplyGrid{}
	grid.GetCollection()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": grid,
	})
}
