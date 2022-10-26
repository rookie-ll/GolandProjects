package main

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) GetName() {
	fmt.Println(a.Name)
}

type Cat struct {
	Age int
	*Animal
}

func (c *Cat) Miao() {
	fmt.Println("喵喵", c.Name)
}

func main() {
	a := &Cat{Age: 12, Animal: &Animal{Name: "菜菜"}}
	a.GetName()
	a.Miao()
}
