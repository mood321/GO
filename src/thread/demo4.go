package main

import (
	"sync"
	"time"
)

var mu sync.Mutex
var done bool

/**
线程共享变量
*/
func main() {
	time.Sleep(1 * time.Second)
	println("start")

	go per1()
	time.Sleep(5 * time.Second)
	mu.Lock()
	done = true
	mu.Unlock()
	println("channel")
	time.Sleep(3 * time.Second)

}

func per1() {
	for {
		println("tick")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			return
		}
		mu.Unlock()
	}
}
