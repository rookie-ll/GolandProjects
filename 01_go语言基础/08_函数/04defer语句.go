package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	//fmt.Println(f1()) //5
	//fmt.Println(f2()) //6
	//fmt.Println(f3()) //5
	//fmt.Println(f4()) //5

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
	//提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	//A 1 2 3
	//B 10 2 12
	//BB 10 12 22
	//AA 1 3 4

}
