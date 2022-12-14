package block

// GridInterface 列表对象接口
type GridInterface interface {
	PrepareCollection()
	PrepareColumns()
	GetHeaderFilters()
}

type ColumnsType = []map[string]interface{}
