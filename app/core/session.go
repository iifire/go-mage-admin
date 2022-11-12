package core

// Session 会话Session
type Session struct {
	formKey string
}

func (session *Session) GetFormKey() string {
	if session.formKey == "" {
		helper := &Helper{}
		session.formKey = helper.GetRandomString(16)
	}
	return session.formKey
}
