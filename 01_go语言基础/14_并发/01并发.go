package main

import (
	"fmt"
	"sync"
	"time"
)

var g sync.WaitGroup

func fun1() {
	defer g.Done()
	time.Sleep(time.Second)
	fmt.Println("我是任务1")
}
func fun2() {
	defer g.Done()
	time.Sleep(time.Second)
	fmt.Println("我是任务2")
}
func fun3() {
	defer g.Done()
	time.Sleep(time.Second)
	fmt.Println("我是任务3")
}

func main() {
	g.Add(3)
	go fun1()
	go fun2()
	go fun3()

	g.Wait()

}
