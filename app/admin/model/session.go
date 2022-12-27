package model

import (
	"github.com/gin-contrib/sessions"
	"go-mage-admin/app/core/model/session"
	"go-mage-admin/app/mage"
)

// Session Admin会话Session
type Session struct {
	session.Abstract
	User *User
}

// Login 登录
func (s *Session) Login(username string, password string) (bool, string) {
	user := &User{}
	flag, msg, u := user.Authenticate(username, password)
	s.User = u
	return flag, msg
}

func (s *Session) GetUser() *User {
	session := sessions.Default(mage.AppGinContext)
	u := new(User)
	v := session.Get("uid")
	var id int
	if v != nil {
		id = v.(int)
	}
	return u.LoadById(id)
}

type SessionUserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
