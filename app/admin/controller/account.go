package controller

import (
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/block"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/core/route"
	mageApp "go-mage-admin/app/mage"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Account{}, _admin.Name)
}

// Account 个人账户
type Account struct {
}

// IndexGET 账户信息
func (a *Account) IndexGET(c *gin.Context) {
	admin := &Admin{}
	// 启用Grid模版来渲染
	tplName := "IndexGET"
	admin.UseFormTpl()
	mageApp.AppGin.HTMLRender = admin.LoadWidgetLayout(tplName)
	c.HTML(http.StatusOK, tplName, gin.H{
		"code":    1,
		"msg":     "ok",
		"urlForm": "/admin/account/edit",
		"urlSave": "/admin/account/save",
	})
}

func (a *Account) EditPOST(c *gin.Context) {
	form := block.AccountForm{}
	form.GetForm()
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
		"data": form,
	})
}

// SavePOST 更新账户信息
func (a *Account) SavePOST(c *gin.Context) {
	admin := &Admin{}
	//username := c.PostForm("username")
	password := c.PostForm("password")
	mobile := c.PostForm("mobile")
	email := c.PostForm("email")
	u := &model.User{}
	u = u.LoadById(admin.GetUserId())

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	u.Password = string(hash)
	u.Mobile = mobile
	u.Email = email
	mageApp.AppDb["Write"].Save(u)
}
