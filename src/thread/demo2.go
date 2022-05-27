package main

import "sync"

/**

 */
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(x int) {
			println(x)
			wg.Done()
		}(i)
		/*go prin(wg, i)
		wg.Done()*/
	}
	wg.Wait()
}

func prin(wg sync.WaitGroup, i int) {
	println(i)

}
