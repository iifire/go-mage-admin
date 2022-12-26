package admin

import (
	"go-mage-admin/app/admin/block"
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/wxmp/model"
)

type FansGrid struct {
	Collection []model.Fans           `json:"collection"`
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

func (g *FansGrid) PrepareCollection() {
	fans := &model.Fans{}
	g.Pager = block.GetGridPager()
	g.Orders = block.GetGridOrders(g.Columns)
	g.Filters = block.GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = fans.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
	g.Renderers = g.PrepareRenderers(g.Collection)
}

func (g *FansGrid) GetCollection() {
	g.Pk = "fans_id"
	g.Columns = make(block.ColumnsType, 0)
	g.AllColumns = make(block.ColumnsType, 0)
	g.Code = helper.GridCodeRole
	g.PrepareColumns()
	g.PrepareCollection()
	g.PrepareExtra()
	g.Columns, g.AllColumns = block.PrepareGrid(g.Columns, g.Code)
}

func (g *FansGrid) PrepareColumns() {
	g.Columns = append(g.Columns, block.ColumnType{
		Header: "ID",
		Align:  "center",
		Width:  "80",
		Index:  "fans_id",
		Show:   true,
		Sort:   true,
	})
	g.Columns = append(g.Columns, block.ColumnType{
		Header:      "昵称",
		Align:       "center",
		Width:       "",
		Index:       "nickname",
		Show:        true,
		Filter:      "header",
		Placeholder: "请输入粉丝昵称...",
	})
	g.Columns = append(g.Columns, block.ColumnType{
		Header: "OpenID",
		Align:  "center",
		Width:  "300",
		Index:  "openid",
		Show:   true,
	})

	actions := make([]block.ButtonType, 0)

	actions = append(actions, block.ButtonType{
		Label: "发送消息",
		Url:   "/admin/wxmpFans/sendMsg",
		Type:  "form",
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
func (g *FansGrid) PrepareExtra() {
	//nothing
}

// PrepareRenderers 自定义渲染器
func (g *FansGrid) PrepareRenderers(collection []model.Fans) map[string]interface{} {
	renderers := make(map[string]interface{})
	g.Collection = collection
	return renderers
}
