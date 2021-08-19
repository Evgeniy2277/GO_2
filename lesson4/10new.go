package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		workers = make(chan struct{}, 10)
		count int64
		mutex sync.Mutex
	)
	for i := 1; i <= 1010; i++ { //долго играл, но до конца так и не понял за необходимое время.
		workers <- struct{}{}
		go func(){
			defer func() {
				mutex.Lock()
				count++
				<-workers
				mutex.Unlock()
			}()
			time.Sleep(200  * time.Microsecond)

		}()
	}

	fmt.Println("count = ",count)


}
