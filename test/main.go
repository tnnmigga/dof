package main

import (
	"fmt"
	"reflect"
)

func setStructFieldValues[T any]() T {
	input := new(T)
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)
	v = v.Elem()
	t = t.Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		if field.Kind() == reflect.String && field.CanSet() {
			field.SetString(fieldType.Name)
		}
	}

	return *input
}

// 定义一个测试结构体
type MyStruct struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// 初始化结构体
	// 设置字段值
	myStruct := setStructFieldValues[MyStruct]()

	// 打印结果
	fmt.Printf("Updated struct: %+v\n", myStruct)
}
