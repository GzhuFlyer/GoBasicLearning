package main

import "fmt"

func main() {
	//var i *int

	list := new([]int)
	*list = append(*list, 1, 2, 3, 4, 5)
	fmt.Println(*list)

	var listA []int
	listA = append(listA, 1, 2, 3, 4)
	fmt.Println(listA)

	listB := make([]int, 0)
	listB = append(listB, 1, 2, 3)
	fmt.Println(listB)
}
