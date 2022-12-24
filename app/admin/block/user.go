package block

import (
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/admin/model"
)

type UserGrid struct {
	Collection []model.User           `json:"collection"`
	Columns    ColumnsType            `json:"columns"`
	AllColumns ColumnsType            `json:"allColumns"`
	Code       string                 `json:"code"`
	Pk         string                 `json:"pk"`
	Pager      GridPager              `json:"pager"`
	Orders     [2]string              `json:"orders"`
	Filters    map[string]interface{} `json:"filters"`
	Buttons    []ButtonType           `json:"buttons"`
	MassAction []ButtonType           `json:"massAction"`
	Exports    []ButtonType           `json:"exports"`
	Renderers  map[string]interface{} `json:"renderers"`
}

func (g *UserGrid) PrepareCollection() {
	user := &model.User{}
	g.Pager = GetGridPager()
	g.Filters = GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = user.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
	g.Renderers = g.PrepareRenderers(g.Collection)
}

func (g *UserGrid) GetCollection() {
	g.Columns = make(ColumnsType, 0)
	g.AllColumns = make(ColumnsType, 0)
	g.Code = helper.GridCodeUser
	g.Pk = "user_id"
	g.PrepareColumns()
	g.PrepareCollection()
	g.PrepareExtra()
	g.Columns, g.AllColumns = PrepareGrid(g.Columns, g.Code)
}

func (g *UserGrid) PrepareColumns() {
	g.Columns = append(g.Columns, ColumnType{
		Header: "ID",
		Align:  "center",
		Width:  "120",
		Index:  "user_id",
		Show:   true,
		Sort:   true,
		Type:   "number",
		Filter: "more",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:       "用户名",
		Align:        "center",
		Width:        "",
		Index:        "username",
		Show:         true,
		Filter:       "header",
		Placeholder:  "请输入用户名...",
		Renderer:     true,
		RendererType: "after",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:      "手机号",
		Align:       "center",
		Width:       "200",
		Index:       "mobile",
		Show:        true,
		Filter:      "header",
		Placeholder: "请输入手机号...",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:       "邮箱",
		Align:        "center",
		Width:        "200",
		Index:        "email",
		Show:         true,
		Filter:       "more",
		Placeholder:  "请输入邮箱...",
		Renderer:     true,
		RendererType: "replace",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:      "状态",
		Align:       "center",
		Width:       "120",
		Index:       "is_active",
		Show:        true,
		Filter:      "more",
		Type:        "options",
		Options:     gin.H{"0": "禁用", "1": "启用"},
		Tag:         true,
		Placeholder: "帐号状态",
		Tags:        gin.H{"0": "info", "1": "success"}, //default success info danger warning
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:    "创建日期",
		Align:     "center",
		Width:     "140",
		Index:     "date_create",
		Show:      false,
		Filter:    "more",
		Type:      "datetime",
		OnlyDate:  true,
		Timestamp: true,
	})
	actions := make([]ButtonType, 0)
	actions = append(actions, ButtonType{
		Label: "编辑",
		Url:   "/admin/user/edit",
		Ajax:  true,
	})
	actions = append(actions, ButtonType{
		Label: "详情",
		Url:   "/admin/user/view",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:       "操作",
		Align:        "center",
		Width:        "140",
		Index:        "action",
		Type:         "action",
		Fixed:        "right", //left, right
		Actions:      actions,
		Renderer:     true,
		RendererType: "after",
		//使用renderer来渲染
	})
}

// PrepareExtra 渲染更多功能
func (g *UserGrid) PrepareExtra() {
	g.Buttons = append(g.Buttons, ButtonType{
		Label: "新增",
		Url:   "/admin/user/add",
		Ajax:  true,
		Class: "btn-blue",
		Icon:  "plus",
	})
	g.MassAction = append(g.MassAction, ButtonType{
		Label: "批量删除",
		Url:   "/admin/user/massDel",
		Ajax:  true,
		Class: "btn-red",
		Icon:  "delete",
	})
	g.Exports = append(g.Exports, ButtonType{
		Label: "XML",
		Url:   "/admin/user/exportXml",
	})
}

// PrepareRenderers 自定义渲染器
func (g *UserGrid) PrepareRenderers(collection []model.User) map[string]interface{} {
	renderers := make(map[string]interface{})
	actionRenderers := make([]interface{}, 0)
	emailRenderers := make([]interface{}, 0)
	usernameRenderers := make([]interface{}, 0)
	collectionCopy := make([]model.User, 0)
	for _, item := range collection {
		actionRds := make([]interface{}, 0)
		if item.IsActive == 0 {
			actionRds = append(actionRds, gin.H{
				"type": "button",
				"button": ButtonType{
					Label: "启用",
					Url:   "/admin/user/enable",
					Ajax:  true,
					Class: "",
					Icon:  "",
				},
			})
		} else {
			actionRds = append(actionRds, gin.H{
				"type": "button",
				"button": ButtonType{
					Label: "禁用",
					Url:   "/admin/user/enable",
					Ajax:  true,
					Class: "",
					Icon:  "",
				},
			})
		}
		if item.IsAdmin == 0 {
			actionRds = append(actionRds, gin.H{
				"type": "button",
				"button": ButtonType{
					Label: "设为管理员",
					Url:   "/admin/user/addAdmin",
					Ajax:  true,
					Class: "",
					Icon:  "",
				},
			})
		} else {
			actionRds = append(actionRds, gin.H{
				"type": "button",
				"button": ButtonType{
					Label: "取消管理员",
					Url:   "/admin/user/rmAdmin",
					Ajax:  true,
					Class: "",
					Icon:  "",
				},
			})
		}
		actionRenderers = append(actionRenderers, actionRds)

		usernameRds := make([]interface{}, 0)
		if item.IsAdmin == 1 {
			usernameRds = append(usernameRds, gin.H{
				"type": "text",
				"text": "<span class='ca'>[管理员]</span>",
			})
		}
		usernameRenderers = append(usernameRenderers, usernameRds)

		emailRds := make([]interface{}, 0)
		if item.Email != "" {
			emailRds = append(emailRds, gin.H{
				"type": "text",
				"text": "<a href='mailto:" + item.Email + "'>" + item.Email + "</a>",
			})
		}
		emailRenderers = append(emailRenderers, emailRds)

		//mobile 隐藏中间4位
		if item.Mobile != "" {
			item.Mobile = item.Mobile[:3] + "****" + item.Mobile[7:]
		}

		collectionCopy = append(collectionCopy, item)
	}
	renderers["action"] = actionRenderers
	renderers["email"] = emailRenderers
	renderers["username"] = usernameRenderers
	g.Collection = collectionCopy
	return renderers
}
