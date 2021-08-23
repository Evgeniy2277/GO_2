package main

import (
	"fmt"
	"sync"
)

const n = 1000

func main() {
	var 	(
		counter int64
		mutex sync.Mutex
		wg = sync.WaitGroup{}
	)
	wg.Add(n)

	for i := 0; i < n; i += 1 {
		go func() {
			defer func() {
				mutex.Unlock()
			}()
			mutex.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)

}
