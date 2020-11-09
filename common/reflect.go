package common

import (
	"reflect"
)

//反射创建新对象。
func ReflectNew(t reflect.Type) (object interface{}) {
	//指针类型获取真正type需要调用Elem
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// 调用反射创建对象
	return reflect.New(t)
}
