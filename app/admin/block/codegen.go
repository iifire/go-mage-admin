package block

import (
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/model"
)

type CodegenGrid struct {
	Collection []model.Codegen        `json:"collection"`
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

func (g *CodegenGrid) PrepareCollection() {
	user := &model.Codegen{}
	g.Pager = GetGridPager()
	g.Orders = GetGridOrders(g.Columns)
	g.Filters = GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = user.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
	g.Renderers = g.PrepareRenderers(g.Collection)
}

func (g *CodegenGrid) GetCollection() {
	g.Columns = make(ColumnsType, 0)
	g.AllColumns = make(ColumnsType, 0)
	g.Pk = "codegen_id"
	g.PrepareColumns()
	g.PrepareCollection()
	g.PrepareExtra()
	g.Columns, g.AllColumns = PrepareGrid(g.Columns, g.Code)
}

func (g *CodegenGrid) PrepareColumns() {
	g.Columns = append(g.Columns, ColumnType{
		Header: "ID",
		Align:  "center",
		Width:  "120",
		Index:  "codegen_id",
		Sort:   true,
		Type:   "number",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:  "类型",
		Align:   "center",
		Index:   "type",
		Filter:  "header",
		Type:    "options",
		Options: gin.H{"0": "模型", "1": "列表", "2": "表单", "3": "模块"},
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:    "创建日期",
		Align:     "center",
		Width:     "180",
		Index:     "date_create",
		Type:      "datetime",
		OnlyDate:  true,
		Timestamp: true,
	})
	actions := make([]ButtonType, 0)
	actions = append(actions, ButtonType{
		Label: "详情",
		Url:   "/admin/user/view",
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
func (g *CodegenGrid) PrepareExtra() {
	g.Buttons = append(g.Buttons, ButtonType{
		Label:  "新增模块",
		Url:    "/admin/codegen/addModule",
		Type:   "form",
		ElType: "info",
		Icon:   "module",
		Sort:   1,
	})
	g.Buttons = append(g.Buttons, ButtonType{
		Label:  "新增模型",
		Url:    "/admin/codegen/addModel",
		Type:   "form",
		ElType: "primary",
		Icon:   "model",
		Sort:   2,
	})
	g.Buttons = append(g.Buttons, ButtonType{
		Label:  "新增列表",
		Url:    "/admin/codegen/addList",
		Type:   "form",
		ElType: "success",
		Icon:   "list",
		Sort:   3,
	})
	g.Buttons = append(g.Buttons, ButtonType{
		Label:  "新增表单",
		Url:    "/admin/codegen/addForm",
		Type:   "form",
		ElType: "warning",
		Icon:   "sdesc",
		Sort:   4,
	})
}

// PrepareRenderers 自定义渲染器
func (g *CodegenGrid) PrepareRenderers(collection []model.Codegen) map[string]interface{} {
	renderers := make(map[string]interface{})
	collectionCopy := make([]model.Codegen, 0)
	for _, item := range collection {
		collectionCopy = append(collectionCopy, item)
	}
	g.Collection = collectionCopy
	return renderers
}
