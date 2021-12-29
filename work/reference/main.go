package main

import "fmt"

func swapReference(x, y *int) {
	fmt.Printf("address &x=%p,&y=%p\n", x, y)
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
func swapValue(x, y int) {
	x, y = y, x
	fmt.Printf("x=%d,y=%d\n", x, y)
}

func main() {
	var a int = 100
	var b int = 200
	fmt.Printf("before change,a=%d,b=%d\n", a, b)
	fmt.Printf("address &x=%p,&y=%p\n", &a, &b)

	swapReference(&a, &b)
	fmt.Printf("after swapReference,a=%d,b=%d\n", a, b)
	a = 100
	b = 200
	swapValue(a, b)
	fmt.Printf("after swapValue,a=%d,b=%d\n", a, b)
}
