package core

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

// AppGin 全局变量 controller action需要渲染模版时要给AppGin.HTMLRender赋值
var AppGin *gin.Engine
var AppDb = make(map[string]*gorm.DB)
var AppConfig *Config
var APPSid string

type App struct {
	config Config
}

// Init 初始化工作
func (app *App) Init(options map[string]interface{}) *App {
	//加载基础配置
	app.config = Config{}
	app.config.LoadBase(options)
	log.Println("app.config.Mysql=", app.config.Mysql)
	AppConfig = &app.config
	app.InitDb()
	app.InitSession()
	return app
}

// InitDb 初始化Db
func (app *App) InitDb() {
	r := AppConfig.Mysql.Read
	w := AppConfig.Mysql.Write
	rDsn := r.Username + ":" + r.Password + "@tcp(" + r.Host + ":" + strconv.Itoa(r.Port) + ")/" + r.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("rDsn=", rDsn)
	Rdb, err := gorm.Open(mysql.Open(rDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	AppDb["read"] = *&Rdb
	wDsn := w.Username + ":" + w.Password + "@tcp(" + w.Host + ":" + strconv.Itoa(w.Port) + ")/" + w.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	Wdb, err := gorm.Open(mysql.Open(wDsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	AppDb["write"] = *&Wdb
	return
}

// InitSession 初始化Session
func (app *App) InitSession() {
	// 区分前台后台Session
	urlPath := AppGin.BasePath()
	APPSid = "mage"
	if strings.HasPrefix(urlPath, "/admin") {
		APPSid = "admin"
	}
	var store cookie.Store
	if AppConfig.Session.Storage == "" || AppConfig.Session.Storage == "Cookie" {
		store = cookie.NewStore([]byte(AppConfig.Session.Key))
		AppGin.Use(sessions.Sessions(APPSid, store))
	} else if AppConfig.Session.Storage == "redis" {
		redisCfg := AppConfig.SessionRedis

		storeRedis, e := redis.NewStoreWithDB(redisCfg.Maxsize, "tcp", redisCfg.Host+":"+strconv.Itoa(redisCfg.Port), redisCfg.Password, strconv.Itoa(redisCfg.Db), []byte(AppConfig.Session.Key))
		storeRedis.Options(sessions.Options{
			Secure:   true,
			SameSite: 4,
			Path:     "/",
			MaxAge:   AppConfig.Session.Lifetime,
		})
		if e != nil {
			log.Panicln(e)
			panic("Redis链接错误")
		}
		AppGin.Use(sessions.Sessions(APPSid, storeRedis))
	} else {
		panic("未正确设置Session Storage[Cookie/Redis]")
	}
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

// Run 运行主入口
func (app *App) Run(options map[string]interface{}) {
	g := gin.Default()
	AppGin = *&g
	app.Init(options)
	//TODO... 全页缓存

	//加载静态资源
	g.Static("/static", "./static")
	//模块化
	//app.InitModules();
	app.InitRequest()
	//
	//g.HTMLRender = app.LoadTemplates("/templates")
	app.GetController(g).Dispatch()
	g.Run(":8082")
	return
}
