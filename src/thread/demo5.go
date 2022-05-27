package main

import (
	"sync"
	"time"
)

/**
经典多线程 计数器
*/
var counter int

func main() {
	counter = 1
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			//defer 保证unlock 在函数体结束执行  finally
			defer mu.Unlock()
			counter++
		}()
	}
	time.Sleep(3 * time.Second)
	println(counter)
}
