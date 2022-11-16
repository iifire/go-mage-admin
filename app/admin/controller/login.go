package controller

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_admin "go-mage-admin/app/admin"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/core"
	"go-mage-admin/app/core/route"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

// init bootstrap初始化时调用 自动注册路由
func init() {
	route.Register(&Login{}, _admin.Name)
}

// Login 后台登录页面
type Login struct {
	// 继承控制器基类
}

func (login *Login) IndexGET(c *gin.Context) {
	// 输出页面 调用对象来处理
	admin := &Admin{}
	admin.SetLayout("1column")

	core.AppGin.HTMLRender = admin.LoadLayout()
	s := sessions.Default(c)
	//s.Set("sid", s.ID())
	log.Println("session user = ", s.Get("user"))

	c.HTML(http.StatusOK, "login.html", gin.H{
		"code": 1,
		"msg":  "ok",
	})

}
func (login *Login) SavePOST(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	remember := c.PostForm("remember")

	sess := &model.Session{}
	flag, msg := sess.Login(username, password)
	if flag {
		s := sessions.Default(c)
		gob.Register(model.SessionUserInfo{})
		sessionUserInfo := &model.SessionUserInfo{
			Id:       sess.User.UserId,
			Username: sess.User.Username,
		}

		s.Set("user", sessionUserInfo)

		if remember == "" {
			//TODO... 会话级Session
		}
		s.Save()
		log.Println(s.Get("user"))
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "登录成功！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  msg,
		})
	}
}

// RegGET 模拟注册用户
func (login *Login) RegGET(c *gin.Context) {
	u := model.User{}
	u.Username = "mage"
	u.Password = "mage"
	u.IsActive = 1
	u.DateCreate = 1668469470
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	u.Password = string(hash)
	core.AppDb["Write"].Create(u)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
	})
}
