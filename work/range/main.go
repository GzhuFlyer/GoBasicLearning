package main

import "fmt"
//==output 0 1 2
func main()  {
	x := []string{"a","b","c"}
	for v:= range x{
		fmt.Println(v)
	}
}
