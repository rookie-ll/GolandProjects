package main

import (
	"fmt"
	"unsafe"
)

//type 类型名 struct {
//	字段名 字段类型
//	字段名 字段类型
//	…
//}
type person struct {
	name string
	city string
	age  int8
}

func main() {
	// 结构体实例化
	var per person
	per.name = "老王"

	// 匿名结构体

	var user struct {
		Name string
		Age  int
	}
	user.Name = "老李"

	// 指针类型结构体

	var pp = new(person)
	pp.name = "老刘"

	//取结构体的地址实例化
	//用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	var pp2 = &person{}
	pp2.name = "老赵"
	(*pp2).age = 12

	//p3.name = "老赵"其实在底层是(*p3).name = "七米"，这是Go语言帮我们实现的语法糖。

	// 空结构体
	var v struct{}
	fmt.Println(unsafe.Sizeof(v))

}

// 构造函数

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
