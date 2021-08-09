package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

var err = errors.New("")


func test()  {

	defer func() {
		if v := recover(); v != nil {
			t := time.Now()
			err = fmt.Errorf(" %w Время возникновения ошибки %s\n", v, t.Format("15.04.05"))
		}
	}()


	pathStr  := ""
	for i := 1 ; i < 1000001; i++ {
		pathStr  = "directory/test" + strconv.Itoa(i)
		os.Create(pathStr)
	}

}

func main() {
	test()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("OK!")
}
