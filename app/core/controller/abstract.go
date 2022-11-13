package controller

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
)

// Abstract 控制器相关常用功能
type Abstract struct {
	gin.Engine
}

func (a *Abstract) LoadLayout(moduleName string, isAlone bool, layoutName string) multitemplate.Renderer {
	return a.loadTemplates(moduleName, isAlone, layoutName)
}

func (a *Abstract) loadTemplates(moduleName string, isAlone bool, layoutName string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	tplPrefix := "./templates/"
	pathLayout := "templates/layouts/" + layoutName + ".html"
	pathInclude := tplPrefix + "includes/*.html"
	if isAlone {
		pathLayout = "templates/" + moduleName + "/layouts/" + layoutName + ".html"
		pathInclude = tplPrefix + moduleName + "/includes/*.html"
	}

	includes, err := filepath.Glob(pathInclude)
	if err != nil {
		panic(err.Error())
	}
	viewsPath := tplPrefix + "views/*.html"
	if isAlone {
		viewsPath = tplPrefix + moduleName + "/views/*.html"
	}

	views, err := filepath.Glob(viewsPath)
	if err != nil {
		panic(err.Error())
	}
	log.Println(views)
	// Generate our templates map from our layouts/ and includes/ directories
	//pathLayout = strings.ReplaceAll()
	for _, v := range views {
		includeCopy := make([]string, 0)
		includeCopy = append(includeCopy, pathLayout)
		//copy(includeCopy, includes)
		files := append(includeCopy, v)
		files = append(files, includes...)

		log.Println("files:", files)
		r.AddFromFiles(filepath.Base(v), files...)
	}
	log.Println("r:", r)
	return r
}
