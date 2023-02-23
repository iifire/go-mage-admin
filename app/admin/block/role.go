package block

import (
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/admin/model"
)

type RoleGrid struct {
	Collection []model.Role           `json:"collection"`
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

func (g *RoleGrid) PrepareCollection() {
	role := &model.Role{}
	g.Pager = GetGridPager()
	g.Orders = GetGridOrders(g.Columns)
	g.Filters = GetGridFilters(g.Columns)
	g.Collection, g.Pager.Total = role.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
	g.Renderers = g.PrepareRenderers(g.Collection)
}

func (g *RoleGrid) GetCollection() {
	g.Pk = "role_id"
	g.Columns = make(ColumnsType, 0)
	g.AllColumns = make(ColumnsType, 0)
	g.Code = helper.GridCodeRole
	g.PrepareColumns()
	g.PrepareCollection()
	g.PrepareExtra()
	g.Columns, g.AllColumns = PrepareGrid(g.Columns, g.Code)
}

func (g *RoleGrid) PrepareColumns() {
	g.Columns = append(g.Columns, ColumnType{
		Header: "ID",
		Align:  "center",
		Width:  "120",
		Index:  "role_id",
		Show:   true,
		Sort:   true,
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:      "角色名称",
		Align:       "center",
		Width:       "",
		Index:       "role_name",
		Show:        true,
		Filter:      "header",
		Placeholder: "请输入角色名称...",
	})

	g.Columns = append(g.Columns, ColumnType{
		Header: "排序",
		Align:  "center",
		Width:  "60",
		Index:  "position",
		Show:   true,
	})

	g.Columns = append(g.Columns, ColumnType{
		Header: "创建时间",
		Align:  "center",
		Width:  "140",
		Index:  "date_create",
		Show:   false,
		Type:   "datetime",
	})
	g.Columns = append(g.Columns, ColumnType{
		Header: "创建人员",
		Align:  "center",
		Width:  "80",
		Index:  "date_create",
		Show:   false,
	})
	actions := make([]ButtonType, 0)

	actions = append(actions, ButtonType{
		Label: "修改",
		Url:   "/admin/role/edit",
		Type:  "edit",
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
func (g *RoleGrid) PrepareExtra() {
	g.Buttons = append(g.Buttons, ButtonType{
		Label:    "新增",
		Method:   "GET",
		Url:      "/admin/role/edit",
		Type:     "form",
		FormType: "dialog",
		FormUrl:  "/admin/role/save",
		ElType:   "primary",
		Icon:     "plus",
	})
	g.MassAction = append(g.MassAction, ButtonType{
		Label:  "批量删除",
		Url:    "/admin/role/massDel",
		Type:   "action",
		ElType: "primary",
		Icon:   "delete",
	})
}

// PrepareRenderers 自定义渲染器
func (g *RoleGrid) PrepareRenderers(collection []model.Role) map[string]interface{} {
	renderers := make(map[string]interface{})
	g.Collection = collection
	return renderers
}
