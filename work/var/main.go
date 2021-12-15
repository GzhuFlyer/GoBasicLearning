package main

import "fmt"

const name = 100

var (
	//size1   := 1024 //error unexpected :=, expecting =
	size    = 1024
	maxSize = size * 2
)

func main() {
	fmt.Println(size, maxSize)
	//fmt.Println(&name, name) //err cannot take the address of name
	fmt.Println(&size, size)
}
