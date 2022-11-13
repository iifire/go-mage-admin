package core

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type App struct {
	config Config
}

// Init 初始化工作
func (app *App) Init(options map[string]interface{}) *App {
	app.config = Config{}
	app.config.LoadBase()
	//加载基础配置

	//
	return app
}

// InitRequest 初始化Request请求
func (app *App) InitRequest() {
	return
}

func (app *App) GetController(g *gin.Engine) *Controller {
	controller := &Controller{g}
	return controller
}

// LoadTemplates 通用多模版加载
func (app *App) LoadTemplates(tplDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(tplDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	views, err := filepath.Glob(tplDir + "/views/*.html")
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

var AppGin *gin.Engine

// Run 运行主入口
func (app *App) Run(options map[string]interface{}) {
	g := gin.Default()
	AppGin = *&g

	//r := multitemplate.NewRenderer()
	//g.HTMLRender = r
	app.Init(options)
	//TODO... 全页缓存
	//模块化
	//app.InitModules();
	app.InitRequest()
	//
	//g.HTMLRender = app.LoadTemplates("/templates")
	app.GetController(g).Dispatch()
	g.Run(":8082")
	return
}
