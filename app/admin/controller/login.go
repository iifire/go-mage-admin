package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/core/route"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Login{}, _admin.Name)
}

// Login 后台登录页面
type Login struct {
	// 继承控制器基类
	//controller.Abstract
}

func (login *Login) Index(c *gin.Context) {
	// 输出页面 调用对象来处理
	admin := &Admin{}
	admin.HTMLRender = admin.LoadLayout()

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
	})

}
func (login *Login) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
	})
}

func (login *Login) Act(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
	})
}
