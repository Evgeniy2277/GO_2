package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const n = 1000

func main() {
	var 	(
		counter int64

		wg = sync.WaitGroup{}
	)
	//инициализируем семафор n
	wg.Add(n)
	fmt.Println(wg)
	// запускаем в цикле n потоков
	for i := 0; i < n; i += 1 {
		go func() {
			atomic.AddInt64(&counter, 1)
			// уменьшаем семафор
			wg.Done()
		}()
	}
	// ждем в сеамфоре 0
	wg.Wait()
	fmt.Println(wg)

	fmt.Println(counter)

}
