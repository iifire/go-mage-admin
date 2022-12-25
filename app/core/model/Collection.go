package model

import (
	"gorm.io/gorm"
	"log"
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
			log.Println("v:", v)
			mv, ok := v.([2]string)
			log.Println("ok:", ok)
			log.Println("len(mv):", len(mv))
			if ok && len(mv) == 2 {
				expr := mv[0]
				expv := mv[1]
				if expr != "" {
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
func AssembleSql(db *gorm.DB, field string, expr string, value string) *gorm.DB {
	if expr == "between" {
		arr := strings.Split(value, "_")
		if len(arr) == 2 {
			db.Where(field+" between ? and ?", arr[0], arr[1])
		}
	} else if expr == "eq" {
		db.Where(field+" = ?", value)
	} else if expr == "neq" {
		db.Where(field+" != ?", value)
	} else if expr == "like" {
		db.Where(field+" like ?", "%"+value+"%")
	} else if expr == "in" {
		where := map[string]interface{}{}
		where[field] = strings.Split(value, ",")
		db.Where(where)
	} else if expr == "nin" {
		//TODO...
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
