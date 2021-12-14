package main

import "fmt"

func main()  {
	//i := 1		//bug ??
	//j := 2
	//k := 3
	//i,j,k = j,k,i
	//fmt.Printf("i=%d,j=%d\n",i,j,k)
	i := 1
	j := 2
	i,j = j,i
	fmt.Printf("i=%d,j=%d\n",i,j)

}
