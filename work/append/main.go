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
func test4() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println("test4(): ", s1)
}
func main() {
	test1()
	test2()
	test3()
	test4()
}
