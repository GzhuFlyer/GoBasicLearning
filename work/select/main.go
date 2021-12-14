package main

import (
	"fmt"
	"runtime"
)

func main()  {
	runtime.GOMAXPROCS(2)
	intChan := make(chan int,1)
	stringChan := make(chan string,1)
	intChan <- 1
	select {
	case value:=<-intChan:
		fmt.Println(value)
	case value:=<-stringChan:
		fmt.Println(value)
	}

}
