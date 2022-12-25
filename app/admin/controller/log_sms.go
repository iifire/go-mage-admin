package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/block"
	"go-mage-admin/app/core/route"
	mageApp "go-mage-admin/app/mage"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&logSms{}, _admin.Name)
}

// logSms 后台首页
type logSms struct {
	// 继承控制器基类
}

func (g *logSms) IndexGET(c *gin.Context) {
	admin := &Admin{}
	tplName := "IndexGET"
	admin.UseGridTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"menu":    "system",
		"urlCur":  "admin/logSms/index",
		"urlGrid": "admin/logSms/grid",
		"msg":     "ok",
	})
}
func (g *logSms) GridREQ(c *gin.Context) {
	grid := block.CoreSmsLogGrid{}
	grid.GetCollection()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": grid,
	})
}
