package block

import (
	"github.com/gin-gonic/gin"
	"go-mage-admin/app/admin/helper"
	"go-mage-admin/app/admin/model"
)

var gridCode = helper.GridCodeUser

type UserGrid struct {
	Collection []model.User `json:"collection"`
	Columns    ColumnsType  `json:"columns"`
}

func (g *UserGrid) PrepareCollection() {
	g.Columns = make(ColumnsType, 0)
	var users []model.User
	user := &model.User{}
	users = user.GetCollection()
	g.Collection = users
}

func (g *UserGrid) GetCollection() []model.User {
	g.PrepareCollection()
	g.PrepareColumns()
	return g.Collection
}

func (g *UserGrid) PrepareColumns() {
	g.Columns = append(g.Columns, map[string]interface{}{
		"header":         "ID",
		"align":          "center",
		"width":          "120",
		"index":          "user_id",
		"default_column": true,
		"options":        "",
	})
	g.Columns = append(g.Columns, map[string]interface{}{
		"header":         "用户名",
		"align":          "center",
		"width":          "",
		"index":          "username",
		"default_column": true,
	})
	g.Columns = append(g.Columns, map[string]interface{}{
		"header":         "邮箱",
		"align":          "center",
		"width":          "200",
		"index":          "email",
		"default_column": true,
	})
	g.Columns = append(g.Columns, map[string]interface{}{
		"header":         "状态",
		"align":          "center",
		"width":          "120",
		"index":          "is_active",
		"default_column": true,
		"type":           "options",
		"options":        gin.H{"0": "禁用", "1": "启用"},
	})
	g.Columns = append(g.Columns, map[string]interface{}{
		"header":         "创建日期",
		"align":          "center",
		"width":          "140",
		"index":          "date_create",
		"default_column": true,
		"type":           "datetime",
		"format":         "Y-m-d H:i:s",
	})
}

// GetHeaderFilters 获取默认筛选项
func (g *UserGrid) GetHeaderFilters() {
	/*
			$filters = array();
			$filters[] = array(
				"header"    => "",
				"type"		=> "text",
				"index"		=> "contract_query",
				"placeholder" => "输入合同编号、名称...",
				"style"	=> "width:220",
			);
			public function assignFilterValue($filterData) {
		    	$filter   = $this->getParam($this->getVarNameFilter(), null);

		        if (is_string($filter)) {
		            $data = $this->helper("adminhtml")->prepareFilterString($filter);

		            foreach ($filterData as &$_filter) {
			    		if (array_key_exists($_filter["index"], $data)) {
			    			$_filter["value"] = $data[$_filter["index"]];
			    		}
			    	}
		        }
			    return $filterData;
		    }
	*/
}
