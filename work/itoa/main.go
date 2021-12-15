package main

import "fmt"

//output: 0 1 zz zz 4
const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func main() {
	fmt.Println(x, y, z, k, p)
}
