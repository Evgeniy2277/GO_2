package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	const n = 1000
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
