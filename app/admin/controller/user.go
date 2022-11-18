package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/route"
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
	userModel := model.User{}
	users := userModel.GetCollection()
	admin := &Admin{}
	admin.UseGridTpl()
	tplName := "IndexGET"
	core.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code": 1,
		"msg":  "ok",
		"grid": gin.H{
			"collection": users,
		},
	})
}

// EditGET 用户新增/编辑
func (u *User) EditGET(c *gin.Context) {
	admin := &Admin{}
	// 启用Grid模版来渲染
	admin.UseFormTpl()
	core.AppGin.HTMLRender = admin.LoadLayout()
	c.HTML(http.StatusOK, "EditGET", gin.H{
		"code": 1,
		"msg":  "ok",
	})
}
