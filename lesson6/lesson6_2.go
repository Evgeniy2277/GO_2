package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	go fmt.Println("I'm working!")
	for i := 0; i < 1e9; i += 1 {
		if i%1e6 == 0 {
			runtime.Gosched()
		}
	}
	go fmt.Println("I'm working2!")
	for i := 0; i < 1e9; i += 1 {
		if i%1e6 == 0 {
			runtime.Gosched()
		}
	}

}
