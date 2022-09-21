package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

		强制类型转换的基本语法如下：
			T(表达式)
		其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.
	*/
}

/*
比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，
而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
*/
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
