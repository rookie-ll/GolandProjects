package main

import "fmt"

//请问下面代码的执行结果是什么？

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	fmt.Printf("%v,%p, %T\n", stus, stus, stus)
	for _, stu := range stus {
		fmt.Println(stu)
		fmt.Printf("%p\n", &stu)
		m[stu.name] = &stu
	}
	fmt.Println(m)
	//for k, v := range m {
	//	fmt.Println(k, "=>", v.name)
	//}
}
