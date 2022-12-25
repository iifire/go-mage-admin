package block

import (
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/core/model"
)

type CoreSmsLogGrid struct {
	Collection []model.CoreSmsLog     `json:"collection"`
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

func (g *CoreSmsLogGrid) PrepareCollection() {
	logSms := &model.CoreSmsLog{}
	g.Pager = GetGridPager()
	g.Orders = GetGridOrders(g.Columns)
	g.Filters = GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = logSms.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
	g.Renderers = g.PrepareRenderers(g.Collection)
}

func (g *CoreSmsLogGrid) GetCollection() {
	g.Pk = "log_id"
	g.Columns = make(ColumnsType, 0)
	g.AllColumns = make(ColumnsType, 0)
	g.Code = helper.GridCodeLogSms
	g.PrepareColumns()
	g.PrepareCollection()
	g.PrepareExtra()
	g.Columns, g.AllColumns = PrepareGrid(g.Columns, g.Code)
}

func (g *CoreSmsLogGrid) PrepareColumns() {
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
		Header:       "手机号",
		Align:        "center",
		Width:        "",
		Index:        "mobile",
		Show:         true,
		Filter:       "header",
		Placeholder:  "请输入手机号...",
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
		Options:     gin.H{"0": "发送失败", "1": "发送成功"},
		Tag:         true,
		Placeholder: "发送状态",
		Tags:        gin.H{"0": "success", "1": "warning"}, //default success info danger warning
	})
	g.Columns = append(g.Columns, ColumnType{
		Header: "验证码",
		Align:  "center",
		Width:  "80",
		Index:  "code",
		Show:   false,
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
		Header: "其他信息",
		Align:  "center",
		Width:  "",
		Index:  "extra",
		Show:   true,
	})

	g.Columns = append(g.Columns, ColumnType{
		Header:    "发送时间",
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
		Url:   "/admin/logSms/del",
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
func (g *CoreSmsLogGrid) PrepareExtra() {
	g.MassAction = append(g.MassAction, ButtonType{
		Label: "批量删除",
		Url:   "/admin/logSms/massDel",
		Type:  "action",
		Class: "btn-red",
		Icon:  "delete",
	})
	g.Exports = append(g.Exports, ButtonType{
		Label: "XML",
		Url:   "/admin/logSms/exportXml",
	})
}

// PrepareRenderers 自定义渲染器
func (g *CoreSmsLogGrid) PrepareRenderers(collection []model.CoreSmsLog) map[string]interface{} {
	renderers := make(map[string]interface{})
	g.Collection = collection
	return renderers
}
