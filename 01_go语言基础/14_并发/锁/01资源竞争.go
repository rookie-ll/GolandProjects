package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var lock sync.Mutex
var rwlock sync.RWMutex
var x = 0

func add() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		//lock.Lock()
		rwlock.RLock()
		x++
		rwlock.RUnlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
