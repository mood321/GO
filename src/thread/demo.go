package main

import "sync"

func main() {

	var s string
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		s = "hello"
		wg.Done()
	}()

	wg.Wait()
	println(s)
}
