package main

import "fmt"

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
	//注意：
	//recover()必须搭配defer使用。
	//defer一定要在可能引发panic的语句之前定义。
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
