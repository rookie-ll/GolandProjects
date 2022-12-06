package main

import (
	"fmt"
	"reflect"
)

// 获取类型
func reflectTyp(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t)
}

type myInt int64

//  获取原始类型
func reflectTypKind(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t)
	fmt.Println(t.Kind(), t.Name(), t.Align())
}

// 获取值
func reflectVal(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is int32, value is %.2f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is int64, value is %.2f\n", v.Float())

	}
}

// 修改值
func reflectValEdit(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Elem().Kind()
	switch k {
	case reflect.Int64:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.445)

	}
}

func main() {
	//var a int = 31
	//reflectTyp(a)
	//var b myInt = 31
	//reflectTypKind(b)
	//var a int64 = 1
	//var b int64 = 100
	//reflectVal(a)
	//reflectValEdit(&a)
	//fmt.Println(a)

	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())

	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())

}
