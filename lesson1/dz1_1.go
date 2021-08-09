package main

import (
	"fmt"
)

func test() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	var a int
	fmt.Println(1/a)
}

func main() {

	test()
	fmt.Println("OK!")
}
