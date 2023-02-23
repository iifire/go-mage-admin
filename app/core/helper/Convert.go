package helper

import "reflect"

// Convert 类型转换功能
type Convert struct {
}

// StructByReflect 遍历struct并且自动进行赋值
func StructByReflect(data map[string]interface{}, inStructPtr interface{}) {
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}
	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		// 得到tag中的字段名
		key := t.Tag.Get("key")
		if v, ok := data[key]; ok {
			// 检查是否需要类型转换
			dataType := reflect.TypeOf(v)
			structType := f.Type()
			if structType == dataType {
				f.Set(reflect.ValueOf(v))
			} else {
				if dataType.ConvertibleTo(structType) {
					// 转换类型
					f.Set(reflect.ValueOf(v).Convert(structType))
				} else {
					panic(t.Name + " type mismatch")
				}
			}
		} else {
			panic(t.Name + " not found")
		}
	}
}
