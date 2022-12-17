package model

import (
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type Collection struct {
}

// PrepareCollection 过滤条件拼接和排序 magento collection->addFieldToFilter()
func (c *Collection) PrepareCollection(db *gorm.DB, filters map[string]interface{}) *gorm.DB {
	for k, v := range filters {
		vt := reflect.TypeOf(v)
		switch vt.Kind() {
		case reflect.Array, reflect.Slice:
			if mv, ok := v.([]interface{}); ok && len(mv) == 2 {
				expr, ok0 := mv[0].(string)
				expv, ok1 := mv[1].(string)
				if ok0 && ok1 {
					//拼接SQL
					db = AssembleSql(db, k, expr, expv)
				}
			} else {
				panic("error.")
			}
		default:
		}
	}
	return db
}
func (c *Collection) LoadCollection(db *gorm.DB, orders [2]string, page int, size int) *gorm.DB {
	if orders[0] != "" && orders[1] != "" {
		db.Order(orders[0] + " " + orders[1])
	}
	if size > 0 {
		db.Offset((page - 1) * size).Limit(size)
	}
	return db
}

// AssembleSql 拼接SQL
//   - - array("between" => $fromValue-$toValue)
//   - - array("eq" => $value)
//   - - array("neq" => $value)
//   - - array("like" => $value)
//   - - array("in" => $v1-$v2-$v3)
//   - - array("nin" => $v1-$v2-$v3)
//   - - array("notnull" => true)
//   - - array("null" => true)
//   - - array("gt" => $value)
//   - - array("lt" => $value)
//   - - array("gteq" => $value)
//   - - array("lteq" => $value)
func AssembleSql(db *gorm.DB, field string, expr string, value string) *gorm.DB {
	if expr == "between" {
		arr := strings.Split(value, "-")
		if len(arr) == 2 {
			db.Where(field+" between ? and ?", arr[0], arr[1])
		}
	} else if expr == "eq" {
		db.Where(field+" = ?", value)
	} else if expr == "neq" {
		db.Where(field+" != ?", value)
	} else if expr == "like" {
		db.Where(field+" like '%?%'", value)
	} else if expr == "in" {
		arr := strings.Split(value, "-")
		db.Where(field+" in (?)", strings.Join(arr, ","))
	} else if expr == "nin" {
		arr := strings.Split(value, "-")
		db.Where(field+" not in (?)", strings.Join(arr, ","))
	} else if expr == "notnull" {
		db.Where(field + " is not null")
	} else if expr == "null" {
		db.Where(field + " is null")
	} else if expr == "gt" {
		db.Where(field+" > ?", value)
	} else if expr == "lt" {
		db.Where(field+" < ?", value)
	} else if expr == "gteq" {
		db.Where(field+" >= ?", value)
	} else if expr == "lteq" {
		db.Where(field+" <= ?", value)
	}
	return db
}
