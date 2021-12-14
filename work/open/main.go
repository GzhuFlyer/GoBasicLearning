package main

import (
	"fmt"
	"os"
)

func main()  {
	file,err := os.Open("main1.go")
	defer file.Close()
	if err != nil{
		fmt.Println("open failed:",err)
		//panic("open fail")
	}else {
		fmt.Println("OK")
	}
	fmt.Println("main exit")
}
