package main

import (
	"fmt"
	calc "pkg_test/calculator"
	"pkg_test/initTest"
)

func init() {
	fmt.Println("main -- init ")
}
func main() {
	var (
		num1 int = 12
		num2 int = 4
	)

	result, err := calc.Calc(num1, num2, "-")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("运算结果为：%v\n", result)

	fmt.Println("main--main")
	initTest.Test()
}
