package model

import "go-mage-admin/app/core/model/session"

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

type SessionUserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
