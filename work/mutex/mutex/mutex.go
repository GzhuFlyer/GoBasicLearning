package main

import (
	"fmt"
	"sync"
)
//output: 没有进行临界区访问控制
//9667
//没有进行临界区访问控制
//9790
var x1 = 0
var mu1 sync.Mutex
var wg1 = sync.WaitGroup{}
func add()  {
	mu1.Lock()
	x1 = x1 + 1
	mu1.Unlock()
	wg1.Done()
}
func main()  {
	sum := 5
	wg1.Add(sum)
	for  i:=0;i<sum;i++{
		go add()
		//fmt.Println(i)
	}
	//time.Sleep(time.Second)
	fmt.Println("\n---------",x1,"---------")
	wg1.Wait()
}
