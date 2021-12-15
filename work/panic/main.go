package main

import (
	"fmt"
	"time"
)

func arrayTest() {
	var s []string
	fmt.Println(s)
	fmt.Println(s[0])
	s = append(s, "hello world")
	fmt.Println(s[0])
}

func main() {
	go arrayTest()
	time.Sleep(time.Second)
	fmt.Println("hello")
	time.Sleep(time.Second)
}
