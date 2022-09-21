package main

import "fmt"

//数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。 基本语法：
// 定义一个长度为3元素类型为int的数组a
// var 数组变量名 [元素数量]T
// var a [3]int

//数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1，访问越界（下标在合法范围之外），则触发访问越界，会panic。

func main() {
	//var a [3]int
	//var b [4]int
	//a = b //不可以这样做，因为此时a和b是不同的类型

	//数组的初始化
	// 方法一
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]

	//方法二
	//一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：
	var testArray1 [3]int
	var numArray1 = [...]int{1, 2}
	var cityArray1 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray1)                          //[0 0 0]
	fmt.Println(numArray1)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray1)   //type of numArray:[2]int
	fmt.Println(cityArray1)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray1) //type of cityArray:[3]string

	//方法三
	//可以使用指定索引值的方式来初始化数组，例如:
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int


	//多维数组
	//多维数组只有第一层可以使用...来让编译器推导数组长度。例如：
	//支持的写法
	arr := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//不支持多维数组的内层使用...
	//brr := [3][...]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}
	fmt.Println(arr)



	// 数组是值类型
	arr1 := [3]int{10, 20, 30}
	modifyArray(arr1) //在modify中修改的是a的副本x
	fmt.Println(arr1) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}

//数组是值类型
//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}


//注意：
//
//	数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
//	[n]*T表示指针数组，*[n]T表示数组指针 。
