package controller

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	moduleAdmin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/core/controller"
	"go-mage-admin/app/core/helper"
	mageApp "go-mage-admin/app/mage"
	"log"
	"path/filepath"
)

// Admin 后台控制器基类
type Admin struct {
	controller.Abstract
	layout   string
	extraTpl []string
}

func (admin *Admin) GetSession() *model.Session {
	session := &model.Session{}
	return session
}

func (admin *Admin) GetUserId() int {
	sess := sessions.Default(mageApp.AppGinContext)
	uid := sess.Get("uid")
	return helper.GetInterfaceToInt(uid)
}

func (admin *Admin) SetLayout(layout string) {
	admin.layout = layout
	return
}

// SetExtraTpl 设置Tpl
func (admin *Admin) SetExtraTpl(extraTpl []string) {
	admin.extraTpl = extraTpl
	return
}

// AddExtraTpl 添加Tpl
func (admin *Admin) AddExtraTpl(tpl string) {
	admin.extraTpl = append(admin.extraTpl, tpl)
	return
}

// UseGridTpl 使用Grid列表Tpl
func (admin *Admin) UseGridTpl() {
	path := "templates/admin/views/widget/grid/*.html"
	admin.UseTplByPath(path)
	admin.AddExtraTpl("templates/admin/views/widget/form/fieldset.html")
	return
}

// UseFormTpl 使用Form表单的Tpl
func (admin *Admin) UseFormTpl() {
	path := "templates/admin/views/widget/form/*.html"
	admin.UseTplByPath(path)
	return
}

func (admin *Admin) UseTplByPath(path string) {
	views, err := filepath.Glob(path)
	if err != nil {
		panic(err.Error())
	}
	for _, v := range views {
		admin.AddExtraTpl(v)
	}
	return
}

// LoadWidgetLayout 加载通用Widget布局
func (admin *Admin) LoadWidgetLayout(name string) multitemplate.Renderer {
	tplArr := make([]string, 0)
	r := multitemplate.NewRenderer()
	pathLayout := "templates/admin/layouts/2columns.html"
	tplArr = append(tplArr, pathLayout)
	pathInclude := "templates/admin/includes/*.html"
	includes, err := filepath.Glob(pathInclude)
	if err != nil {
		panic(err.Error())
	}
	for _, v := range includes {
		tplArr = append(tplArr, v)
	}
	for _, v := range admin.extraTpl {
		tplArr = append(tplArr, v)
	}
	log.Println("tplArr=", tplArr)
	r.AddFromFiles(name, tplArr...)

	return r
}

func (admin *Admin) LoadLayout() multitemplate.Renderer {
	a := &moduleAdmin.Admin{}
	if admin.layout == "" {
		admin.layout = "2columns"
	}
	return a.LoadLayout(admin.layout, admin.extraTpl)
}
