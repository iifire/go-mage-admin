package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/block"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/core/route"
	mageApp "go-mage-admin/app/mage"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Codegen{}, _admin.Name)
}

// Codegen 代码生成
type Codegen struct {
	// 继承控制器基类

}

// IndexGET 代码列表
func (u *Codegen) IndexGET(c *gin.Context) {
	admin := &Admin{}
	tplName := "IndexGET"
	admin.UseGridTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"menu":    "dev",
		"urlCur":  "/admin/codegen/index",
		"urlGrid": "/admin/codegen/grid",
		"msg":     "ok",
	})
}

func (u *Codegen) GridREQ(c *gin.Context) {
	grid := block.CodegenGrid{}
	grid.GetCollection()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": grid,
	})
}

// EditGET 代码新增/编辑
func (u *Codegen) EditGET(c *gin.Context) {
	admin := &Admin{}
	// 启用Grid模版来渲染
	admin.UseFormTpl()
	mageApp.AppGin.HTMLRender = admin.LoadLayout()
	c.HTML(http.StatusOK, "EditGET", gin.H{
		"code": 1,
		"msg":  "ok",
	})
}

// MassDelPOST 批量删除
func (u *Codegen) MassDelPOST(c *gin.Context) {
	ids, _ := c.GetPostFormArray("ids[]")
	user := &model.Codegen{}
	if len(ids) > 0 {
		flag := user.DelByIds(ids)
		if flag {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "删除成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "删除失败！",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "未选中记录！",
		})
	}

}
