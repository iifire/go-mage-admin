package block

import (
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/admin/model"
)

type UserGrid struct {
	Collection    []model.User           `json:"collection"`
	Columns       ColumnsType            `json:"columns"`
	MoreColumns   ColumnsType            `json:"moreColumns"`
	HeaderFilters HeaderFiltersType      `json:"headerFilters"`
	MoreFilters   MoreFiltersType        `json:"moreFilters"`
	Code          string                 `json:"code"`
	Pager         GridPager              `json:"pager"`
	Orders        [2]string              `json:"orders"`
	Filters       map[string]interface{} `json:"filters"`
}

func (g *UserGrid) PrepareCollection() {
	user := &model.User{}
	g.Pager = GetGridPager()
	g.Collection, g.Pager.Total = user.GetCollection(g.Filters, g.Orders, g.Pager.Page, g.Pager.Size)
}

func (g *UserGrid) GetCollection() {
	g.Columns = make(ColumnsType, 0)
	g.MoreColumns = g.Columns
	g.Code = helper.GridCodeUser
	g.PrepareCollection()
	g.PrepareColumns()
	g.GetHeaderFilters()
	g.Columns, g.MoreFilters = PrepareGrid(g.Columns, g.Code)
}

func (g *UserGrid) PrepareColumns() {
	g.Columns = append(g.Columns, ColumnType{
		Header: "ID",
		Align:  "center",
		Width:  "120",
		Index:  "user_id",
		Show:   true,
		Sort:   true,
	})
	g.Columns = append(g.Columns, ColumnType{
		Header: "用户名",
		Align:  "center",
		Width:  "",
		Index:  "username",
		Show:   true,
	})
	g.Columns = append(g.Columns, ColumnType{
		Header: "邮箱",
		Align:  "center",
		Width:  "200",
		Index:  "email",
		Show:   true,
		Filter: true,
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:  "状态",
		Align:   "center",
		Width:   "120",
		Index:   "is_active",
		Show:    true,
		Filter:  true,
		Type:    "options",
		Options: gin.H{"0": "禁用", "1": "启用"},
		Tag:     true,
		Tags:    gin.H{"0": "info", "1": "success"}, //default success info danger warning
	})
	g.Columns = append(g.Columns, ColumnType{
		Header:    "创建日期",
		Align:     "center",
		Width:     "140",
		Index:     "date_create",
		Show:      false,
		Filter:    true,
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

// GetHeaderFilters 获取默认筛选项
func (g *UserGrid) GetHeaderFilters() {
	g.HeaderFilters = append(g.HeaderFilters, gin.H{
		"header":      "用户名",
		"type":        "text",
		"index":       "username",
		"placeholder": "请输入用户名...",
		"style":       "width:220px",
	})
	g.HeaderFilters = append(g.HeaderFilters, gin.H{
		"header":      "",
		"type":        "text",
		"index":       "email",
		"placeholder": "请输入邮箱地址..",
		"style":       "width:220px",
	})

}

/*
public function assignFilterValue($filterData) {
	$filter   = $this->getParam($this->getVarNameFilter(), null);

	if (is_string($filter)) {
		$data = $this->helper("adminhtml")->prepareFilterString($filter);

		foreach ($filterData as &$_filter) {
			if (array_key_exists($_filter[Index], $data)) {
				$_filter["value"] = $data[$_filter[Index]];
			}
		}
	}
	return $filterData;
}*/
