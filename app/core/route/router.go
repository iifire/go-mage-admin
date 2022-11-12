package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/core/helper"
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
		method := v.Method(i)
		action := v.Type().Method(i).Name
		path := "/" + moduleName + "/" + strHelper.DeCapitalize(ctrlNameCopy) + "/" + strHelper.DeCapitalize(action)
		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		httpMethod := "GET"
		if strings.Contains(action, "_post") {
			httpMethod = "post"
		}

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
		if route.httpMethod == "GET" {
			e.GET(route.path, match(route.path, route))
		}
		if route.httpMethod == "POST" {
			e.POST(route.path, match(route.path, route))
		}
	}
}

// match根据path匹配对应的方法
func match(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		fmt.Println("fields,len(fields)=", fields, len(fields))
		if len(fields) < 3 {
			return
		}
		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context
			//reflect.ValueOf(method).Method(0).Call(arguments)
			route.Method.Call(arguments)
		}
	}
}
