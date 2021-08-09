package main

import (
	"errors"
	"fmt"
	"time"
)
var err = errors.New("error")

func test()  {

	defer func() {
		if v := recover(); v != nil {
			t := time.Now()
			err = fmt.Errorf(" %w Время возникновения ошибки %s\n", v, t.Format("15.04.05"))
		}
	}()

	var a int
	fmt.Println(1/a)
}

func main() {
	test()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("OK!")


}
