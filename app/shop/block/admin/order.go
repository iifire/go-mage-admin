package admin

import (
	"go-mage-admin/app/admin/block"
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/shop/model"
)

type OrderGrid struct {
	Collection []model.Order          `json:"collection"`
	Columns    block.ColumnsType      `json:"columns"`
	AllColumns block.ColumnsType      `json:"allColumns"`
	Code       string                 `json:"code"`
	Pk         string                 `json:"pk"`
	Pager      block.GridPager        `json:"pager"`
	Orders     [2]string              `json:"orders"`
	Filters    map[string]interface{} `json:"filters"`
	Buttons    []block.ButtonType     `json:"buttons"`
	MassAction []block.ButtonType     `json:"massAction"`
	Exports    []block.ButtonType     `json:"exports"`
	Renderers  map[string]interface{} `json:"renderers"`
}

func (g *OrderGrid) PrepareCollection() {
	fans := &model.Order{}
	g.Pager = block.GetGridPager()
	g.Orders = block.GetGridOrders(g.Columns)
	g.Filters = block.GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = fans.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
	g.Renderers = g.PrepareRenderers(g.Collection)
}

func (g *OrderGrid) GetCollection() {
	g.Pk = "order_id"
	g.Columns = make(block.ColumnsType, 0)
	g.AllColumns = make(block.ColumnsType, 0)
	g.Code = helper.GridCodeRole
	g.PrepareColumns()
	g.PrepareCollection()
	g.PrepareExtra()
	g.Columns, g.AllColumns = block.PrepareGrid(g.Columns, g.Code)
}

func (g *OrderGrid) PrepareColumns() {
	g.Columns = append(g.Columns, block.ColumnType{
		Header: "编号",
		Align:  "center",
		Width:  "110",
		Index:  "increment_id",
		Show:   true,
		Sort:   true,
		Filter: "header",
	})
	g.Columns = append(g.Columns, block.ColumnType{
		Header: "金额",
		Align:  "center",
		Width:  "80",
		Index:  "grand_total",
		Show:   true,
	})
	g.Columns = append(g.Columns, block.ColumnType{
		Header:   "下单日期",
		Align:    "center",
		Width:    "120",
		Index:    "date_create",
		Show:     true,
		OnlyDate: true,
		Type:     "datetime",
	})

	actions := make([]block.ButtonType, 0)

	actions = append(actions, block.ButtonType{
		Label: "详情",
		Url:   "/admin/shopOrder/view",
		Type:  "view",
	})
	g.Columns = append(g.Columns, block.ColumnType{
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
func (g *OrderGrid) PrepareExtra() {
	//nothing
}

// PrepareRenderers 自定义渲染器
func (g *OrderGrid) PrepareRenderers(collection []model.Order) map[string]interface{} {
	renderers := make(map[string]interface{})
	g.Collection = collection
	return renderers
}
