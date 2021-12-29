package main

//使用#cgo定义库路径


/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L ./helloC -lhello
#include "./helloC/hello.h"
*/
import "C"

func main() {
	C.hello()
}