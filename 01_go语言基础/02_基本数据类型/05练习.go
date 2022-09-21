package main

import "fmt"

//1. 编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
//2. 编写代码统计出字符串"hello沙河小王子"中汉字的数量。

func topic01() {
	var i int
	fmt.Printf("值：%v, 类型：%T\n", i, i)
	var f float32
	fmt.Printf("值：%v, 类型：%T\n", f, f)
	var b bool
	fmt.Printf("值：%v, 类型：%T\n", b, b)
	var s string
	fmt.Printf("值：%v, 类型：%T\n", s, s)
}

func topic02() {
	var s = "hello内蒙古小王子"
	runeBytes := []rune(s)
	fmt.Println("runner数组", runeBytes)
	var sum int
	for _, i := range runeBytes {
		if i > 2<<7 {
			sum++
		}
	}
	fmt.Println(sum)
}
func main() {
	//topic01(
	topic02()
}
