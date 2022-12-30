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
	route.Register(&WxmpMaterial{}, _admin.Name)
}

// WxmpMaterial 自定义菜单
type WxmpMaterial struct {
}

// IndexGET 菜单列表页
func (m *WxmpMaterial) IndexGET(c *gin.Context) {
	admin := &controller.Admin{}
	extraTpl := make([]string, 0)
	extraTpl = append(extraTpl, "wxmp/material.html")
	admin.SetExtraTpl(extraTpl)
	mageApp.AppGin.HTMLRender = admin.LoadLayout()
	c.HTML(http.StatusOK, "wxmp_material.html", gin.H{
		"code":   1,
		"msg":    "ok",
		"menu":   "wxmp",
		"urlCur": "/admin/wxmpMaterial/index",
	})
}

// EditGET 菜单列表页
func (m *WxmpMaterial) EditGET(c *gin.Context) {
	admin := &controller.Admin{}
	extraTpl := make([]string, 0)
	extraTpl = append(extraTpl, "wxmp/material/edit.html")
	admin.SetExtraTpl(extraTpl)
	mageApp.AppGin.HTMLRender = admin.LoadLayout()
	c.HTML(http.StatusOK, "wxmp_material_edit.html", gin.H{
		"code":   1,
		"msg":    "ok",
		"menu":   "wxmp",
		"urlCur": "/admin/wxmpMaterial/index",
	})
}
