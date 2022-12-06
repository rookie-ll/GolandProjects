package main

import (
	"fmt"
	"sync"
)

var gw sync.WaitGroup

// 1. 生成0 - 100 的数字发送到ch1
func f1(ch chan int) {
	for i := 0; i < 101; i++ {
		ch <- i
	}
	close(ch)
}

// 2. 从ch1中取出数据计算他的平方，放到ch2中
func f2(ch1, ch2 chan int) {
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func recv(c chan int) {
	defer gw.Done()
	ret := <-c
	fmt.Println(ret)
}

func main() {
	//gw.Add(1)
	//c := make(chan int)
	//go recv(c)
	//c <- 10
	//close(c)
	//gw.Wait()

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)
	go f1(ch1)
	go f2(ch1, ch2)
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
