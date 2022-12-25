package block

import (
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/admin/model"
)

type LogLoginGrid struct {
	Collection []model.LogLogin       `json:"collection"`
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

func (g *LogLoginGrid) PrepareCollection() {
	logLogin := &model.LogLogin{}
	g.Pager = GetGridPager()
	g.Orders = GetGridOrders(g.Columns)
	g.Filters = GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = logLogin.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
	g.Renderers = g.PrepareRenderers(g.Collection)
}

func (g *LogLoginGrid) GetCollection() {
	g.Pk = "log_id"
	g.Columns = make(ColumnsType, 0)
	g.AllColumns = make(ColumnsType, 0)
	g.Code = helper.GridCodeLogLogin
	g.PrepareColumns()
	g.PrepareCollection()
	g.PrepareExtra()
	g.Columns, g.AllColumns = PrepareGrid(g.Columns, g.Code)
}

func (g *LogLoginGrid) PrepareColumns() {
	g.Columns = append(g.Columns, ColumnType{
		Header: "ID",
		Align:  "center",
		Width:  "120",
		Index:  "log_id",
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
		Header:      "状态",
		Align:       "center",
		Width:       "120",
		Index:       "status",
		Show:        true,
		Filter:      "header",
		Type:        "options",
		Options:     gin.H{"0": "登录失败", "1": "登录成功"},
		Tag:         true,
		Placeholder: "系统登陆日志状态",
		Tags:        gin.H{"0": "warning", "1": "success"}, //default success info danger warning
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:      "IP地址",
		Align:       "center",
		Width:       "200",
		Index:       "mobile",
		Show:        true,
		Filter:      "header",
		Placeholder: "请输入ip地址...",
	})

	g.Columns = append(g.Columns, ColumnType{
		Header:    "登录时间",
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
		Label: "删除",
		Url:   "/admin/logLogin/del",
		Type:  "action",
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
func (g *LogLoginGrid) PrepareExtra() {
	g.MassAction = append(g.MassAction, ButtonType{
		Label: "批量删除",
		Url:   "/admin/logLogin/massDel",
		Type:  "action",
		Class: "btn-red",
		Icon:  "delete",
	})
	g.Exports = append(g.Exports, ButtonType{
		Label: "XML",
		Url:   "/admin/logLogin/exportXml",
	})
}

// PrepareRenderers 自定义渲染器
func (g *LogLoginGrid) PrepareRenderers(collection []model.LogLogin) map[string]interface{} {
	renderers := make(map[string]interface{})
	g.Collection = collection
	return renderers
}
