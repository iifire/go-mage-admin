package block

type AccountForm struct {
	Title     string         `json:"title"`
	Fieldsets []FieldsetType `json:"fieldsets"`
}

// GetForm 加载表单
func (f *AccountForm) GetForm() {
	f.PrepareForm()
	f.PrepareExtra()
}

// PrepareForm 装配表单
func (f *AccountForm) PrepareForm() {
	f.Title = "账户信息"
	//fieldset tr td
	fd := FieldsetType{
		Label:  "",
		Class:  "",
		Style:  "",
		Layout: []string{"10%", "90%"},
		//Items: []TrType
	}
	tr := TrType{
		Class: "",
		//Items: []TdType
	}
	tr.Items = append(tr.Items, TdType{
		Name:        "mobile",
		Label:       "手机号",
		Type:        "text",
		Placeholder: "请填写手机号",
	})
	fd.Items = append(fd.Items, tr)
	tr2 := TrType{
		Class: "",
		//Items: []TdType
	}
	tr2.Items = append(tr2.Items, TdType{
		Name:        "email",
		Label:       "邮箱地址",
		Type:        "text",
		Placeholder: "请填写邮箱地址",
	})
	fd.Items = append(fd.Items, tr2)
	tr3 := TrType{
		Class: "",
	}
	tr3.Items = append(tr3.Items, TdType{
		Name:        "password",
		Label:       "密码",
		Type:        "password",
		Placeholder: "请填写密码",
	})
	fd.Items = append(fd.Items, tr3)
	tr4 := TrType{
		Class: "",
	}
	tr4.Items = append(tr4.Items, TdType{
		Name:        "password_copy",
		Label:       "重复密码",
		Type:        "password",
		Placeholder: "请再次输入密码",
	})
	fd.Items = append(fd.Items, tr4)
	f.Fieldsets = append(f.Fieldsets, fd)
}

// PrepareExtra 渲染更多功能
func (f *AccountForm) PrepareExtra() {

}
