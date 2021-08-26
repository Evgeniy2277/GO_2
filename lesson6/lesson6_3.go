package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		count = 6
		wg = sync.WaitGroup{}

	)
	wg.Add(count)
	m := make(map[string]string)
	go func() {
		defer wg.Done()
		m["1"] = "a"
	}()
	go func() {
		defer wg.Done()
		m["2"] = "b"
	}()
	go func() {
		defer wg.Done()
		m["3"] = "c"
	}()
	go func() {
		defer wg.Done()
		m["4"] = "d"
	}()
	go func() {
		defer wg.Done()
		m["5"] = "e"
	}()
	go func() {
		defer wg.Done()
		m["6"] = "f"
	}()

	wg.Wait()


	for k, v := range m {
		fmt.Println(k, v)
	}

}
