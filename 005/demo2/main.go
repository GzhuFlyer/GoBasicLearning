package main

import (
	"fmt"
	"strings"
)

//lack off deffer panic test
func main() {
	test1()
	panic("test")
	test2()
	for {
	}
}

func test1() {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "HAL-9000")) //匿名函数
}

func add1(r rune) rune {
	return r + 1
}

func test2() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(temp())
	fmt.Println(temp())
}
func squares() func() int {
	var x int //这里函数变量拥有状态
	return func() int {
		x++
		return x * x
	}
}
func temp() int {
	var x int
	x++
	return x
}
