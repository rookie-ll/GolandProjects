package main

import "fmt"

//因为数组的长度是固定的并且数组长度属于类型的一部分，所以数组有很多的局限性。 例如：

func arraySum(x [3]int) int {
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//这个求和函数只能接受[3]int类型，其他的都不支持。 再比如，
//
//a := [3]int{1, 2, 3}
//数组a中已经有三个元素了，我们不能再继续往数组a中添加新元素了。

//切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
//
//切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。

func main() {
	var a []string
	//var b []string
	a = append(a, "老李")
	fmt.Println(a == nil) //false
	// fmt.Println(a == b)   //切片是引用类型，不支持直接比较，只能和nil比较
	fmt.Printf("值：%v, 类型：%T, 长度：%v, 容量：%v\n", a, a, len(a), cap(a))

	arr := [5]int{1, 2, 3, 4, 5}
	s := arr[1:3] // s:= arr[low:high]
	//长度=high-low
	//arr[low : high : max]
	//容量 = max - low
	t := arr[2:2:3]

	fmt.Printf("值：%v, 类型：%T, 长度：%v, 容量：%v\n", t, t, len(t), cap(t))
	fmt.Printf("值：%v, 类型：%T, 长度：%v, 容量：%v\n", s, s, len(s), cap(s))
}
