package controller

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

// Abstract 控制器相关常用功能
type Abstract struct {
	gin.Engine
}

func (a *Abstract) LoadLayout(moduleName string, isAlone bool) multitemplate.Renderer {
	return a.loadTemplates(moduleName, isAlone)
}

func (a *Abstract) loadTemplates(moduleName string, isAlone bool) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	pathLayout := "/templates/layouts/*.html"
	if isAlone {
		pathLayout = "/templates/" + moduleName + "/layouts/*.html"
	}
	layouts, err := filepath.Glob(pathLayout)
	if err != nil {
		panic(err.Error())
	}
	viewLayout := "/templates/views/*.html"
	if isAlone {
		viewLayout = "/templates/" + moduleName + "/views/*.html"
	}
	views, err := filepath.Glob(viewLayout)
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, v := range views {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, v)
		r.AddFromFiles(filepath.Base(v), files...)
	}
	return r
}
