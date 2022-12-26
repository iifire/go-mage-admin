package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/core/helper"
	mageApp "go-mage-admin/app/mage"
	"net/http"
	"reflect"
	"strings"
)

// Route 路由结构体
type Route struct {
	path       string         //url路径
	httpMethod string         //http方法 get post
	Method     reflect.Value  //方法路由
	Args       []reflect.Type //参数类型
}

// Routes 路由集合
var Routes []Route

func InitRouter(g *gin.Engine) {
	Bind(g)
	return
}

// Register 注册控制器
func Register(controller interface{}, moduleName string) bool {
	ctrlName := reflect.TypeOf(controller).String()
	//fmt.Println("ctrlName=", ctrlName)
	ctrlNameCopy := ctrlName
	if strings.Contains(ctrlName, ".") {
		ctrlNameCopy = ctrlName[strings.Index(ctrlName, ".")+1:]
	}
	v := reflect.ValueOf(controller)
	strHelper := &helper.String{}
	//遍历方法
	for i := 0; i < v.NumMethod(); i++ {
		action := v.Type().Method(i).Name
		httpMethod := ""
		if strings.HasSuffix(action, "ANY") {
			httpMethod = "ANY"
		} else if strings.HasSuffix(action, "REQ") {
			httpMethod = "REQ"
		} else if strings.HasSuffix(action, "GET") {
			httpMethod = "GET"
		} else if strings.HasSuffix(action, "POST") {
			httpMethod = "POST"
		} else if strings.HasSuffix(action, "PUT") {
			httpMethod = "PUT"
		} else if strings.HasSuffix(action, "PATCH") {
			httpMethod = "PATCH"
		} else if strings.HasSuffix(action, "DELETE") {
			httpMethod = "DELETE"
		}
		if httpMethod == "" {
			continue
		}
		method := v.Method(i)
		actName := strings.Replace(action, httpMethod, "", -1)
		path := "/" + moduleName + "/" + strHelper.DeCapitalize(ctrlNameCopy) + "/" + strHelper.DeCapitalize(actName)
		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}
		route := Route{path: path, Method: method, Args: params, httpMethod: httpMethod}
		Routes = append(Routes, route)
	}
	//fmt.Println("Routes=", Routes)
	return true
}

// Bind 绑定路由 m是方法GET POST等
func Bind(e *gin.Engine) {
	//路由分组
	for _, route := range Routes {
		if route.httpMethod == "ANY" {
			e.Any(route.path, match(route.path, route))
		} else if route.httpMethod == "REQ" {
			e.POST(route.path, match(route.path, route))
			e.GET(route.path, match(route.path, route))
		} else if route.httpMethod == "GET" {
			e.GET(route.path, match(route.path, route))
		} else if route.httpMethod == "POST" {
			e.POST(route.path, match(route.path, route))
		} else if route.httpMethod == "PUT" {
			e.PUT(route.path, match(route.path, route))
		} else if route.httpMethod == "PATCH" {
			e.PATCH(route.path, match(route.path, route))
		} else if route.httpMethod == "DELETE" {
			e.DELETE(route.path, match(route.path, route))
		}
	}
}

// match根据path匹配对应的方法
func match(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		//fmt.Println("fields,len(fields)=", fields, len(fields))
		if len(fields) < 3 {
			return
		}
		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context
			//注册Session
			session := sessions.Default(c)
			var count int
			v := session.Get("count")
			if v == nil {
				count = 0
			} else {
				count = v.(int)
				count++
			}

			session.Set("count", count)
			session.Save()
			/******** 临时手动注册用户 start *******/
			session.Set("uid", 1)
			/******** 临时手动注册用户 end *******/
			//判断后台请求 未登录则跳转到登录页面
			uid := session.Get("uid")
			if strings.HasPrefix(path, "/admin") && path != "/admin/login/index" {
				//判断是否已登录
				if uid == nil {
					c.Redirect(http.StatusFound, "/admin/login/index")
				} else {
					//更新cookie失效时间

				}
			} else if path == "/admin/login/index" {
				if uid != nil {
					c.Redirect(http.StatusFound, "/admin/dashboard/index")
				}
			}
			mageApp.AppGinContext = c
			route.Method.Call(arguments)
		}
	}
}
