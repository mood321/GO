package main

import (
	"sync"
	"time"
)

/**
锁粒度 影响数据中间状态
*/

var banlce int
var alice int
var bob int
var total int

func main() {
	banlce = 1000
	bob = 1000
	var mu sync.Mutex

	total = bob + alice

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			alice -= 1
			/*mu.Unlock()
			mu.Lock()*/
			bob += 1
			mu.Unlock()

		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			bob -= 1
			/*mu.Unlock()
			mu.Lock()*/
			alice += 1
			mu.Unlock()

		}
	}()
	start := time.Now()

	for time.Since(start) < 1*time.Second {
		mu.Lock()
		if alice+bob != total {
			println("金额错误,{},{},{}", alice, bob, total)
		}
		mu.Unlock()
	}
	println("金额,{},{},{}", alice, bob, total)
}
