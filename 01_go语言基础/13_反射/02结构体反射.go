package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (t Student) Sleep(a string) {
	fmt.Println("睡觉", a)
}

func (t Student) Study() {
	fmt.Println("学习")
}

func reflectStruct(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.Name(), t.Kind(), t.NumField(), t.NumMethod())
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Name"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method:%s\n", methodType)

		fmt.Printf("method name:%s\n", t.Method(i).Name)

	}
	// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
	var args = []reflect.Value{reflect.ValueOf("睡觉睡觉")}
	v.MethodByName("Sleep").Call(args)
}
func main() {
	st := Student{
		Name: "老刘",
		Age:  12,
	}
	//reflectStruct(st)
	printMethod(st)
}
