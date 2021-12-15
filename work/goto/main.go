package main

import "fmt"

func test1() {
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为： %d\n", a)
		a++
	}
}

//嵌套for循环打印九九乘法表
func print9x() {
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
		}
		fmt.Println("")
	}
}

//for循环配合goto打印九九乘法表
func gotoTag() {
	for m := 1; m < 10; m++ {
		n := 1
	LOOP:
		if n <= m {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}
}
func main() {
	test1()
	gotoTag()
}
