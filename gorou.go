package main

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)
var wg sync.WaitGroup
var a = 0
var countMux sync.RWMutex
func gofunc(s string) {
	t := 0
	defer wg.Done()
	for i:= 0 ; i < 5 ; i++{
		fmt.Println("Info =", s,"times : ",t)
		// 延时1秒
		time.Sleep(time.Second)
		t++
		defer countMux.RUnlock()
		countMux.RLock()
		a++
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go gofunc("go coroutine function")

	go func (s string) {
		t := 0
		defer wg.Done()
		for i:= 0 ; i < 3 ; i++ {
			fmt.Println("Info =", s,"2 times : ",t)
			// 延时1秒
			time.Sleep(time.Second)
			t++
			defer countMux.RUnlock()
			countMux.RLock()
			fmt.Println(a)
		}
	}("go coroutine function")
	wg.Wait()
	fmt.Println("11111")
	// for {
	// 	fmt.Println("Info =", "main function","times : ",t)
	// 	// 延时1秒
	// 	time.Sleep(time.Second)
	// 	t++
	// }
}