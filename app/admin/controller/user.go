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
	route.Register(&User{}, _admin.Name)
}

// User 后台首页
type User struct {
	// 继承控制器基类

}

// IndexGET 用户列表
func (u *User) IndexGET(c *gin.Context) {
	admin := &Admin{}
	tplName := "IndexGET"
	admin.UseGridTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"menu":    "system",
		"urlCur":  "/admin/user/index",
		"urlGrid": "/admin/user/grid",
		"msg":     "ok",
	})
}

func (u *User) GridREQ(c *gin.Context) {
	grid := block.UserGrid{}
	grid.GetCollection()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": grid,
	})
}

// EditGET 用户新增/编辑
func (u *User) EditGET(c *gin.Context) {
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
func (u *User) MassDelPOST(c *gin.Context) {
	ids, _ := c.GetPostFormArray("ids[]")
	user := &model.User{}
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
