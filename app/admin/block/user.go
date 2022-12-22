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
	Pager      GridPager              `json:"pager"`
	Orders     [2]string              `json:"orders"`
	Filters    map[string]interface{} `json:"filters"`
	Buttons    []ButtonType           `json:"buttons"`
	MassAction []ButtonType           `json:"massAction"`
	Exports    []ButtonType           `json:"exports"`
}

func (g *UserGrid) PrepareCollection() {
	user := &model.User{}
	g.Pager = GetGridPager()
	g.Filters = GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = user.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
}

func (g *UserGrid) GetCollection() {
	g.Columns = make(ColumnsType, 0)
	g.AllColumns = make(ColumnsType, 0)
	g.Code = helper.GridCodeUser
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
		Header:      "用户名",
		Align:       "center",
		Width:       "",
		Index:       "username",
		Show:        true,
		Filter:      "header",
		Placeholder: "请输入用户名...",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:      "邮箱",
		Align:       "center",
		Width:       "200",
		Index:       "email",
		Show:        true,
		Filter:      "header",
		Placeholder: "请输入邮箱...",
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
		Format:    "Y-m-d H:i:s",
		Timestamp: true,
	})
	g.Columns = append(g.Columns, ColumnType{
		Header: "操作",
		Align:  "center",
		Width:  "140",
		Index:  "action",
		Type:   "action",
		Fixed:  "right", //left, right
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
