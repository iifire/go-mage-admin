package block

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/model"
	"go-mage-admin/app/mage"
	"strconv"
	"strings"
	"time"
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
	Header       string       `json:"header"`
	Align        string       `json:"align"`
	Width        string       `json:"width"`
	Index        string       `json:"index"`
	Type         string       `json:"type"`
	Show         bool         `json:"show"`
	Sort         bool         `json:"sort"`
	Tag          bool         `json:"tag"`
	Timestamp    bool         `json:"timestamp"`
	OnlyDate     bool         `json:"onlyDate"`
	Options      gin.H        `json:"options"`
	Tags         gin.H        `json:"tags"`
	Actions      []ButtonType `json:"actions"`
	Renderer     bool         `json:"renderer"`
	RendererType string       `json:"rendererType"` //replace / before / after

	Fixed string `json:"fixed"`
	//Value     string `json:"value"`
	// 筛选属性
	Filter      string `json:"filter"` //header / more
	Multiple    bool   `json:"multiple"`
	Placeholder string `json:"placeholder"`
	Style       string `json:"style"`
}

// ButtonType 按钮类型 页面跳转
type ButtonType struct {
	Label      string `json:"label"`
	Url        string `json:"url"`
	Method     string `json:"method"`
	Class      string `json:"class"`
	Sort       int    `json:"sort"`
	Type       string `json:"type"`     //link|tab[历史菜单] / popup / action / form / view
	ElType     string `json:"elType"`   //primary / success / warning / danger / info / text
	FormType   string `json:"formType"` //drawer[默认] / dialog
	FormUrl    string `json:"formUrl"`
	Icon       string `json:"icon"`
	Confirm    bool   `json:"confirm"`
	ConfirmTxt string `json:"confirmTxt"`
}

// ColumnsType 列切片
type ColumnsType = []ColumnType

// PrepareGrid 预处理Grid
func PrepareGrid(columns ColumnsType, gridCode string) (ColumnsType, ColumnsType) {
	var columnsCopy ColumnsType
	var allColumnsCopy ColumnsType
	session := sessions.Default(mage.AppGinContext)
	uid, _ := strconv.Atoi(fmt.Sprintf("%v", session.Get("uid")))
	user := model.User{}
	columnsCfg := user.GetGridColumnsConfig(uid, gridCode)
	columnsCfgSet := ConvertStrSlice2Map(columnsCfg)
	for _, column := range columns {
		if columnsCfg != nil && len(columnsCfg) > 0 {
			if InMap(columnsCfgSet, column.Index) {
				columnsCopy = append(columnsCopy, column)
				if column.Type != "action" {
					column.Show = true
					allColumnsCopy = append(allColumnsCopy, column)
				}
			} else {
				if column.Type == "action" {
					columnsCopy = append(columnsCopy, column)
				}
				if column.Type != "action" {
					column.Show = false
					allColumnsCopy = append(allColumnsCopy, column)
				}
			}
		} else {
			if gridCode == "" || column.Show || column.Type == "action" {
				columnsCopy = append(columnsCopy, column)
			}
			if column.Type != "action" {
				allColumnsCopy = append(allColumnsCopy, column)
			}
		}
	}
	return columnsCopy, allColumnsCopy
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

// GetGridOrders 获取排序
func GetGridOrders(columns ColumnsType) [2]string {
	orders := [2]string{}
	index := mage.AppGinContext.DefaultPostForm("sort", "")
	order := mage.AppGinContext.DefaultPostForm("order", "")
	if index != "" && (order == "asc" || order == "desc") {
		for _, v := range columns {
			if v.Sort && v.Index == index {
				orders[0] = v.Index
				orders[1] = order
			}
		}
	}
	return orders
}

// GetGridFilters 获取过滤
// TODO... 防SQL注入
func GetGridFilters(columns ColumnsType) map[string]interface{} {
	filters := make(map[string]interface{})
	fs, _ := mage.AppGinContext.GetPostFormArray("filters[]")
	if len(fs) > 1 {
		i := 0
		for ; i < len(fs); i++ {
			for _, v := range columns {
				if v.Index == fs[i] {
					if v.Type == "" || v.Type == "text" {
						filters[fs[i]] = [2]string{"like", fs[i+1]}
					} else if v.Type == "options" {
						if strings.Contains(fs[i+1], ",") {
							//多选
							filters[fs[i]] = [2]string{"in", fs[i+1]}
						} else if fs[i+1] != "" {
							filters[fs[i]] = [2]string{"eq", fs[i+1]}
						}
					} else if v.Type == "datetime" {
						if strings.Contains(fs[i+1], "_") {
							timeTmp := strings.Split(fs[i+1], "_")
							timeStartStr := timeTmp[0]
							timeEndStr := timeTmp[1]
							if v.Timestamp {
								tsStart := ToTimeStamp(timeStartStr)
								tsEnd := ToTimeStamp(timeEndStr)
								if timeStartStr != "" && timeEndStr != "" {
									if tsStart != "" && tsEnd != "" {
										filters[fs[i]] = [2]string{"between", tsStart + "_" + tsEnd}
									}
								} else if timeStartStr != "" && tsStart != "" {
									filters[fs[i]] = [2]string{"gteq", tsStart}
								} else if timeEndStr != "" && tsEnd != "" {
									filters[fs[i]] = [2]string{"lteq", tsEnd}
								}
							} else {
								//DB保存日期格式
								if timeStartStr != "" && timeEndStr != "" {
									filters[fs[i]] = [2]string{"between", timeStartStr + "_" + timeEndStr}
								} else if timeStartStr != "" {
									filters[fs[i]] = [2]string{"gteq", timeStartStr}
								} else if timeEndStr != "" {
									filters[fs[i]] = [2]string{"lteq", timeEndStr}
								}
							}
						}
					} else if v.Type == "number" && strings.Contains(fs[i+1], "_") {
						numTmp := strings.Split(fs[i+1], "_")
						numStart := numTmp[0]
						numEnd := numTmp[1]
						if numStart != "" && numEnd != "" {
							filters[fs[i]] = [2]string{"between", numStart + "_" + numEnd}
						} else if numStart != "" {
							filters[fs[i]] = [2]string{"gteq", numStart + "_" + numEnd}
						} else if numEnd != "" {
							filters[fs[i]] = [2]string{"lteq", numStart + "_ " + numEnd}
						}
					}
					break
				}
			}
			i++
		}
	}
	return filters
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

// ToTimeStamp 字符串转时间戳
func ToTimeStamp(str string) string {
	if strings.Contains(str, "-") {
		arr := strings.Split(str, "-")
		if len(arr) == 3 {
			y, err0 := strconv.Atoi(arr[0])
			m, err1 := strconv.Atoi(arr[1])
			d, err2 := strconv.Atoi(arr[2])
			if err0 == nil && err1 == nil && err2 == nil {
				ts := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local).Unix()
				return strconv.FormatInt(ts, 10)
			}
		}
	}
	return ""
}
