package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/core/route"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Page{}, _admin.Name)
}

// Page 后台页面公共部分
type Page struct {
	// 继承控制器基类
}

// HeaderGET 获取Header
func (p *Page) HeaderGET(c *gin.Context) {
	// 输出JSON
	menu := &model.Menu{}

	topMenu := menu.GetTopMenus()
	//TODO... 权限校验
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": topMenu,
	})

}
