package main

import "fmt"

func main() {
	//var a *int
	//*a = 100
	//fmt.Println(*a)
	//var b map[string]int
	//b["沙河娜扎"] = 100
	//fmt.Println(b)
	/*
		执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
		而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。
		Go语言中new和make是内建的两个函数，主要用来分配内存。
	*/

	var a1 = new(int)
	*a1 = 100
	fmt.Println(*a1)

}
