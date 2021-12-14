package main

import "fmt"
//output: 3,2,1,panic:4

func main()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("4")
}
