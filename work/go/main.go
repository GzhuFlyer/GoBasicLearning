package main

import (
	"fmt"
	"time"
)
//three three three

func main()  {
	str := []string{"one","two","three"}
	for _,s := range str{
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ",s)
		}()
	}
	time.Sleep(3 * time.Second)
}