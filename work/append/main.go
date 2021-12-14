package main

import "fmt"

func test1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
func test2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}
func test3() {
	s := new([]int)
	//var y []int
	//s = &y
	*s = append(*s, 1, 2, 3)
	fmt.Println(s)
}
func main() {
	test1()
	test2()
	test3()
}
