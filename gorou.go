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

func gorou() {
	runtime.GOMAXPROCS(1)
	wg.Add(1)
	chanName := make(chan string,5)
	// go gofunc("go coroutine function")

	go func () {
		defer wg.Done()
		for i:= 0 ; i < 3 ; i++ {
			chanName <- "go coroutine function"
			fmt.Println("send once.....")
		}
		// close(chanName)
		time.Sleep(time.Second)
	}()
	// wg.Wait()
	for {
		time.Sleep(time.Second)
		time.Sleep(time.Second)
		select {
		case msg,ok := <- chanName :
			if ok {
				fmt.Println("chan Msg =", msg," ok = ",ok)
			} else {
				fmt.Println("chan is closed"," ok = ",ok)
				break
			}
		case chanName <- "main send" :
			fmt.Println("main send")
		case <-time.After(3 * time.Second):
			fmt.Println("TimeOut")
			wg.Wait()
			fmt.Println("finished")
			return
		default :
			fmt.Println("default")
		}
	}	
}
