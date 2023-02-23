package block

import "go-mage-admin/app/admin/model"

type RoleForm struct {
	Title     string         `json:"title"`
	Fieldsets []FieldsetType `json:"fieldsets"`

	role *model.Role
}

// GetForm 加载表单
func (f *RoleForm) GetForm(role *model.Role) {
	f.role = role
	f.PrepareForm()
	f.PrepareExtra()
}

// PrepareForm 装配表单
func (f *RoleForm) PrepareForm() {
	if f.role.RoleId > 0 {
		f.Title = "编辑角色"
	} else {
		f.Title = "新增角色"
	}
	//fieldset tr td
	fd := FieldsetType{
		Label:  "",
		Class:  "",
		Style:  "",
		Layout: []string{"15%", "85%"},
		//Items: []TrType
	}
	tr := TrType{
		Class: "",
		//Items: []TdType
	}
	tr.Items = append(tr.Items, TdType{
		Name:        "role_name",
		Label:       "角色名称",
		Type:        "text",
		Placeholder: "请填写角色名称",
		Required:    true,
	})
	fd.Items = append(fd.Items, tr)

	tr3 := TrType{
		Class: "",
	}
	vs := make([]map[string]interface{}, 0)
	vs = append(vs, map[string]interface{}{
		"label": "全部权限",
		"value": "1",
	})
	vs = append(vs, map[string]interface{}{
		"label": "自定义",
		"value": "0",
	})
	tr3.Items = append(tr3.Items, TdType{
		Name:   "allow_all",
		Label:  "权限设置",
		Type:   "select",
		Values: vs,
	})
	fd.Items = append(fd.Items, tr3)
	tr2 := TrType{
		Class: "",
		//Items: []TdType
	}
	tr2.Items = append(tr2.Items, TdType{
		Name:        "memo",
		Label:       "备注",
		Type:        "textarea",
		Placeholder: "请填写备注",
	})
	fd.Items = append(fd.Items, tr2)
	f.Fieldsets = append(f.Fieldsets, fd)
}

// PrepareExtra 渲染更多功能
func (f *RoleForm) PrepareExtra() {

}
