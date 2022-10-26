package main

import (
	"fmt"
)

func funcA(a int) {
	fmt.Println(a)
}

func funcB(a *int) {
	fmt.Println(*a)
}
func main() {
	a := 10
	b := &a
	fmt.Println(b)
	funcA(a)
	funcB(&a)
}
