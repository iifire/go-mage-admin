package controller

import (
	"github.com/gin-contrib/multitemplate"
	moduleAdmin "go-mage-admin/app/admin"
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/controller"
)

// Admin 后台控制器基类
type Admin struct {
	controller.Abstract
}

func (admin *Admin) GetSession() *core.Session {
	session := &core.Session{}
	return session
}
func (admin *Admin) LoadLayout() multitemplate.Renderer {
	a := &moduleAdmin.Admin{}
	return a.LoadLayout()
}
