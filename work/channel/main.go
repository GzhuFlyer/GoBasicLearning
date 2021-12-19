package main

import "fmt"

func main() {
	fmt.Println("begin doing something!")
	c := make(chan bool)
	go func() {
		fmt.Println("sub goroutine")
		//close(c)
		c <- true
	}()
	//<-c
	fmt.Println("main routine exit", <-c)
}
