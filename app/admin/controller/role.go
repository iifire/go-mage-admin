package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/block"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/core/helper"
	"go-mage-admin/app/core/route"
	mageApp "go-mage-admin/app/mage"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Role{}, _admin.Name)
}

// Role 后台首页
type Role struct {
	// 继承控制器基类
}

func (g *Role) IndexGET(c *gin.Context) {
	admin := &Admin{}
	tplName := "IndexGET"
	admin.UseGridTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"menu":    "system",
		"urlCur":  "/admin/role/index",
		"urlGrid": "/admin/role/grid",
		"msg":     "ok",
	})
}
func (g *Role) GridREQ(c *gin.Context) {
	grid := block.RoleGrid{}
	grid.GetCollection()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": grid,
	})
}

// EditGET 新增/编辑
func (g *Role) EditGET(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	r := &model.Role{}
	if id != "" {
		r = r.LoadById(helper.GetStringToInt(id))
	}

	form := block.RoleForm{}
	form.GetForm(r)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": form,
	})
}

// SavePOST 更新账户信息
func (g *Role) SavePOST(c *gin.Context) {
	r := &model.Role{}
	id := c.PostForm("id")
	if id != "" {
		r = r.LoadById(helper.GetStringToInt(id))
	}
	params, _ := helper.GetPostFormParams(c)
	helper.StructByReflect(params, r)
	mageApp.AppDb["write"].Save(r)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "角色信息更新成功！",
	})
}
