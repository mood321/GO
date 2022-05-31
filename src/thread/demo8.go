package main

import "time"

/**
chnnal 使用

*/

func main() {
	c := make(chan bool)

	go func() {
		println("send took func")
		time.Sleep(1 * time.Second)
		<-c
	}()
	now := time.Now()
	c <- true

	println("send took %v", time.Since(now))
}
