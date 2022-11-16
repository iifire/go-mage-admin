package session

import (
	helper2 "go-mage-admin/app/core/helper"
)

// Abstract 会话Session
type Abstract struct {
	formKey string
}

func (session *Abstract) GetFormKey() string {
	if session.formKey == "" {
		h := &helper2.String{}
		session.formKey = h.GetRandomString(16)
	}
	return session.formKey
}
