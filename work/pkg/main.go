package main

import "fmt"

func test(x int) (func(), func()) {
	return func() {
			fmt.Println(x)
			x += 10
		}, func() {
			fmt.Println(x)
		}
}

func main() {
	a, b := test(100)
	fmt.Println("a=", a, "b=", &b)
	a()
	b()
}
