package block

import (
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/admin/model"
)

type LogOpsGrid struct {
	Collection []model.LogOps         `json:"collection"`
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

func (g *LogOpsGrid) PrepareCollection() {
	logOps := &model.LogOps{}
	g.Pager = GetGridPager()
	g.Orders = GetGridOrders(g.Columns)
	g.Filters = GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = logOps.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
	g.Renderers = g.PrepareRenderers(g.Collection)
}

func (g *LogOpsGrid) GetCollection() {
	g.Pk = "log_id"
	g.Columns = make(ColumnsType, 0)
	g.AllColumns = make(ColumnsType, 0)
	g.Code = helper.GridCodeLogOps
	g.PrepareColumns()
	g.PrepareCollection()
	g.PrepareExtra()
	g.Columns, g.AllColumns = PrepareGrid(g.Columns, g.Code)
}

func (g *LogOpsGrid) PrepareColumns() {
	g.Columns = append(g.Columns, ColumnType{
		Header: "编号",
		Align:  "center",
		Width:  "120",
		Index:  "log_id",
		Show:   true,
		Sort:   true,
		Type:   "number",
		Filter: "more",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:       "Request info",
		Align:        "left",
		Width:        "",
		Index:        "url",
		Show:         true,
		Filter:       "more",
		Placeholder:  "请输入URL...",
		Renderer:     true,
		RendererType: "after",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:       "操作人员",
		Align:        "center",
		Width:        "80",
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
		Options:     gin.H{"0": "正常", "-1": "异常"},
		Tag:         true,
		Placeholder: "操作状态",
		Tags:        gin.H{"0": "success", "-1": "warning"}, //default success info danger warning
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
		Header:    "操作日期",
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
		Label: "详细",
		Url:   "/admin/logOps/view",
		Type:  "view",
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
func (g *LogOpsGrid) PrepareExtra() {
	g.MassAction = append(g.MassAction, ButtonType{
		Label: "批量删除",
		Url:   "/admin/logOps/massDel",
		Type:  "action",
		Class: "btn-red",
		Icon:  "delete",
	})
	g.Exports = append(g.Exports, ButtonType{
		Label: "XML",
		Url:   "/admin/logOps/exportXml",
	})
}

// PrepareRenderers 自定义渲染器
func (g *LogOpsGrid) PrepareRenderers(collection []model.LogOps) map[string]interface{} {
	renderers := make(map[string]interface{})
	g.Collection = collection
	return renderers
}
