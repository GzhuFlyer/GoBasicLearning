package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b float64
	c int16
}
var y struct {
	a bool
	c int16
	b float64
}

func main() {
	fmt.Println("sizeof(x)=", unsafe.Sizeof(x))
	fmt.Println("unsafe.Alignof(x)=", unsafe.Alignof(x))
	fmt.Println("sizeof(x)=", unsafe.Sizeof(x))
	fmt.Println("\n-------------------\n")
	fmt.Println("sizeof(y)=", unsafe.Sizeof(y))
	fmt.Println("unsafe.Alignof(y)=", unsafe.Alignof(y))
	fmt.Println("sizeof(y)=", unsafe.Sizeof(y))

}
