package block

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/mage"
	"strconv"
)

// GridInterface 列表对象接口
type GridInterface interface {
	PrepareCollection()
	PrepareColumns()
	GetHeaderFilters()
}

type GridPager struct {
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

// ColumnType 列类型
type ColumnType struct {
	Header    string `json:"header"`
	Align     string `json:"align"`
	Width     string `json:"width"`
	Index     string `json:"index"`
	Type      string `json:"type"`
	Show      bool   `json:"show"`
	Filter    bool   `json:"filter"`
	Sort      bool   `json:"sort"`
	Tag       bool   `json:"tag"`
	Timestamp bool   `json:"timestamp"`
	Options   gin.H  `json:"options"`
	Tags      gin.H  `json:"tags"`
	Format    string `json:"format"`
	Fixed     string `json:"fixed"`
}

type ColumnsType = []ColumnType
type HeaderFiltersType = []map[string]interface{}
type MoreFiltersType = []ColumnType

func PrepareGrid(columns ColumnsType, gridCode string) (ColumnsType, MoreFiltersType) {
	var moreFilters MoreFiltersType
	var columnsCopy ColumnsType
	session := sessions.Default(mage.AppGinContext)
	uid, _ := strconv.Atoi(fmt.Sprintf("%v", session.Get("uid")))
	user := model.User{}
	columnsCfg := user.GetGridColumnsConfig(uid, gridCode)
	columnsCfgSet := ConvertStrSlice2Map(columnsCfg)
	for _, column := range columns {
		if columnsCfg != nil && len(columnsCfg) > 0 {
			if InMap(columnsCfgSet, column.Index) {
				columnsCopy = append(columnsCopy, column)
			}
		} else {
			if column.Show {
				columnsCopy = append(columnsCopy, column)
			}
		}
		if column.Filter {
			moreFilters = append(moreFilters, column)
		}
	}
	return columnsCopy, moreFilters
}

// GetGridPager 获取分页对象
func GetGridPager() GridPager {
	pager := GridPager{}
	page, _ := strconv.Atoi(mage.AppGinContext.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(mage.AppGinContext.DefaultPostForm("size", "10"))
	pager.Page = page
	pager.Size = size
	return pager
}
func ConvertStrSlice2Map(sl []string) map[string]struct{} {
	set := make(map[string]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InMap(m map[string]struct{}, s string) bool {
	_, ok := m[s]
	return ok
}
