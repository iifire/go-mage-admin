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
	route.Register(&logOps{}, _admin.Name)
}

// logOps 后台首页
type logOps struct {
	// 继承控制器基类
}

func (g *logOps) IndexGET(c *gin.Context) {
	admin := &Admin{}
	tplName := "IndexGET"
	admin.UseGridTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"menu":    "system",
		"urlCur":  "/admin/logOps/index",
		"urlGrid": "/admin/logOps/grid",
		"msg":     "ok",
	})
}
func (g *logOps) GridREQ(c *gin.Context) {
	grid := block.LogOpsGrid{}
	grid.GetCollection()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": grid,
	})
}
