package main

import "time"

/**
模拟定时任务
*/
func main() {
	time.Sleep(1 * time.Second)
	println("start")

	go per()
	time.Sleep(5 * time.Second)
}

func per() {
	for {
		println("tick")
		time.Sleep(1 * time.Second)
	}
}
