package main

import (
	"fmt"
	"sync"
	"time"
)
//output: 没有进行临界区访问控制
//9667
//没有进行临界区访问控制
//9790
func main()  {
	x := 0
	sum := 10000
	wg := sync.WaitGroup{}
	wg.Add(sum)
	for  i:=0;i<sum;i++{
		go func() {
			x = x + 1
			wg.Done()
		}()
	}
	if sum != x{
		fmt.Println("没有进行临界区访问控制")
	}
	time.Sleep(time.Second)
	fmt.Println(x)
	wg.Wait()
}
