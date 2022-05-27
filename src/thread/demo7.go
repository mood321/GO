package main

import (
	"math/rand"
	"sync"
	"time"
)

/**
简易版raft
*/

func main() {

	rand.Seed(time.Now().UnixNano())
	count := 0
	fin := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	for i := 0; i < 10; i++ {

		go func() {

			vote := requestVote()
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			fin++
			cond.Broadcast()
		}()
	}
	mu.Lock()
	for count < 5 && fin != 10 {
		//mu.Lock()
		println("wait ...")
		/*time.Sleep(500 * time.Millisecond)
		mu.Unlock()*/
		cond.Wait()
	}
	if count >= 5 {
		println("大于5")
	} else {
		println("loost")
	}
	mu.Unlock()

}

func requestVote() bool {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	println("requestVote ...")

	return rand.Int()%2 == 0
}
