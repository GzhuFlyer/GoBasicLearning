package main

import "C"

import (
	"fmt"
)

//export Foo
func Foo(a, b int) int {
	return a + b;
}

//export Bar
func Bar() {
	fmt.Println("I am Go language!")
}

func main() {}
